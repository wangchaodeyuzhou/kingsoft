package conf

import (
	"errors"
	"flag"
	"git.kingsoft.go/intermediate/kqueue/config"
	"golang.org/x/exp/slog"
)

func LoadConfig() (*Bootstrap, error) {
	path := flag.String("configs", "./kqueue/conf/configs.yaml", "the configs file path")
	flag.Parse()
	c, err := config.New(*path, func() any {
		return DefaultConfig()
	})
	if err != nil {
		return nil, err
	}

	if _, ok := c.GetData().(*Bootstrap); !ok {
		return nil, errors.New("cfg data must be bootstrap")
	}

	bs, ok := c.GetData().(*Bootstrap)
	if !ok {
		slog.Error("cfg data must be BootStrap")
		return nil, errors.New("cfg data must be BootStrap")
	}

	bootstrap = bs
	// load to services
	for _, worker := range bs.Services.Workers {
		for _, node := range worker.Node {
			services[worker.WorkerId] = append(services[worker.WorkerId], node)
		}
	}

	return bs, nil
}

func GetConfig() *Bootstrap {
	return bootstrap
}

func GetConfigServices() map[string][]WorkerNode {
	return services
}
