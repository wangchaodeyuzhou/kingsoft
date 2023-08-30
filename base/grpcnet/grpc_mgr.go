package grpcnet

import (
	"google.golang.org/grpc"
	"time"
)

type GRPC struct {
	Network string
	Addr    string
	Timeout time.Duration
}

type GrpcManager struct {
	rpcAddr        string
	localRpcServer *grpc.Server
	rpcClients     map[string]*SameTypeClientMgr
	watchServices  []string
}

func NewGrpcManager(c *GRPC, ss []string) *GrpcManager {
	b := &GrpcManager{
		rpcAddr:       c.Addr,
		rpcClients:    make(map[string]*SameTypeClientMgr, 20),
		watchServices: ss,
	}

	b.localRpcServer = NewGrpcServer(c)
	return b
}

func (g *GrpcManager) ConnectRemoteServer(serverName string, addr string, serverID string, meta map[string]string, tags []string) error {

	remoteClients, ok := g.rpcClients[serverName]
	if !ok {
		// 不存在的化出现什么? 去 watchServices 里面寻找 没有 serverName
	}

	rpcClient, err := NewRPCClient(addr, serverName, serverID, meta, tags)
	if err != nil {
		return err
	}

	remoteClients.AddRPCServer(rpcClient)
	return nil
}

func (g *GrpcManager) DelServer(serverName string, serverID string) {
	remoteMap, ok := g.rpcClients[serverName]
	if ok {
		remoteMap.DeleteRPCServerById(serverID)
	}
}
