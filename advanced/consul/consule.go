package consul

import (
	"fmt"
	consulApi "github.com/hashicorp/consul/api"
	"log/slog"
)

const (
	consulAddr = "localhost:8500"
	serverId   = "wrc"
)

func RegistryConsul() {

	config := consulApi.DefaultConfig()
	config.Address = consulAddr
	client, err := consulApi.NewClient(config)
	if err != nil {
		slog.Error("Failed to create consul client", "err", err)
		return
	}

	registertion := new(consulApi.AgentServiceRegistration)
	registertion.ID = serverId
	registertion.Name = "test-consul"
	registertion.Port = 12000
	registertion.Tags = []string{"go-test-wrc", "new-test-wrc"}
	registertion.Address = "localhost"

	// 增加健康检查回调函数
	check := new(consulApi.AgentServiceCheck)
	check.HTTP = fmt.Sprintf("http://%s:%d", registertion.Address, registertion.Port)
	check.Timeout = "5s"
	check.Interval = "5s"
	check.DeregisterCriticalServiceAfter = "30s"
	registertion.Check = check

	if err = client.Agent().ServiceRegister(registertion); err != nil {
		slog.Error("ServiceRegistration ", "err", err)
		return
	}
}

func DeRegister() {
	config := consulApi.DefaultConfig()
	config.Address = consulAddr
	client, err := consulApi.NewClient(config)
	if err != nil {
		slog.Error("new client have err", "err", err)
		return
	}

	if err = client.Agent().ServiceDeregister(serverId); err != nil {
		fmt.Println("ServiceDeregister fail", "err", err)
		return
	}
	slog.Info("ServiceDeregister success", "serverId", serverId)
}

func consulFindServer() {
	config := consulApi.DefaultConfig()
	config.Address = consulAddr
	client, err := consulApi.NewClient(config)
	if err != nil {
		slog.Error("new client have err", "err", err)
		return
	}

	services, _ := client.Agent().Services()
	for _, service := range services {
		fmt.Println(fmt.Sprintf("http://%s:%d", service.Address, service.Port))
	}

	service, _, err := client.Agent().Service(serverId, nil)
	if err == nil {
		fmt.Println(fmt.Sprintf("http://%s:%d", service.Address, service.Port))
	}
	fmt.Println("Service find server done")
}

func consulCheckHeath() {
	config := consulApi.DefaultConfig()
	config.Address = consulAddr
	client, err := consulApi.NewClient(config)
	if err != nil {
		slog.Error("new client have err", "err", err)
		return
	}

	a, b, _ := client.Agent().AgentHealthServiceByID(serverId)
	fmt.Println("val1: ", a, "val2: ", b)
	fmt.Println("consul check heath done")
}

func consulKVTest() {
	config := consulApi.DefaultConfig()
	config.Address = consulAddr
	client, err := consulApi.NewClient(config)
	if err != nil {
		slog.Error("new client fail", "err", err)
		return
	}

	key := "hello"
	value := "world"
	client.KV().Put(&consulApi.KVPair{Key: key, Value: []byte(value), Flags: 0}, nil)

	data, meta, _ := client.KV().Get(key, nil)
	fmt.Println(data, meta)
}
