package discovery

import (
	"git.shiyou.kingsoft.com/server/consul"
	"log/slog"
)

const (
	GMetadataID      = "id"
	GServiceNameFlag = "-"
)

type IDiscovery interface {
	ConnectRemoteServer(string, string, string, map[string]string, []string) error
	DelServer(string, string)
}

type Consul struct {
	Address   string
	Tags      []string
	FilterTag string
}

type ConsulDiscovery struct {
	cfg          *Consul
	selfServerID string // 自定义 serverId
	serviceNames []string
	disManager   IDiscovery
	watcher      *consul.CatalogWatcher // kgs consul need copy it
}

func NewConsulDiscovery(cfg *Consul, filter []string, id string, d IDiscovery) error {
	slog.Info("new service consul discovery")

	names := make([]string, 0, len(filter))
	copy(names, filter)
	slog.Debug("new all service name", "name", names)

	ret := &ConsulDiscovery{
		cfg:          cfg,
		selfServerID: id,
		serviceNames: filter,
		disManager:   d,
	}
	consulConfig := consul.Config{
		Address: cfg.Address,
	}

	watcher, err := consul.NewWatcherAndRun(consulConfig,
		consul.WithCatalogNameFilter(filter...),
		consul.WithCatalogChangedFunc(func(s string, ss []*consul.ServiceStatus) {
			ret.onChange(ss)
		}),
		consul.WithAllowStale(true))
	if err != nil {
		slog.Error("start consul watch failed", "err", err)
	}
	ret.watcher = watcher

	return nil
}

// onChange 检测服务出现变化
func (d *ConsulDiscovery) onChange(status []*consul.ServiceStatus) {
	slog.Info("watch discovery service onchange")

	for _, v := range status {
		slog.Info("service has changed", "name", v.Service.Name, "status", v.Service, "alive", v.Alive)

		realId := getServerID(v.Service.Meta)
		if v.Alive && v.Service != nil && realId != d.selfServerID {
			if err := d.disManager.ConnectRemoteServer(v.Service.Name, v.Service.Host, realId, v.Service.Meta, v.Service.Tags); err != nil {
				slog.Error("connect remote server", "err", err, "type", v.Service.Name)
				continue
			}
		}

		if !v.Alive && v.Service != nil {
			d.disManager.DelServer(v.Service.Name, realId)
			slog.Debug("del server", "name", v.Service.Name, "id", realId)
		}

	}
}

func getServerID(meta map[string]string) string {
	if id, ok := meta[GMetadataID]; ok {
		return id
	}

	slog.Warn("getServerID empty")

	return ""
}
