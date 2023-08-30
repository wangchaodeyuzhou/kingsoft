package discovery

import (
	"github.com/hashicorp/consul/api"
)

func Register(addr string, s *api.AgentServiceRegistration) error {
	config := api.DefaultConfig()
	config.Address = addr
	client, err := api.NewClient(config)
	if err != nil {
		return err
	}

	if err = client.Agent().ServiceRegister(s); err != nil {
		return err
	}
	return nil
}
