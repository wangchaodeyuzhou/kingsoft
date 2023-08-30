package grpcnet

import "sync"

type SameTypeClientMgr struct {
	lock          sync.RWMutex
	serverName    string
	remoteClients map[string]*GrpcClient
}

func NewSampTypeClientMgr(name string) *SameTypeClientMgr {
	return &SameTypeClientMgr{
		serverName:    name,
		remoteClients: make(map[string]*GrpcClient),
	}
}

func (s *SameTypeClientMgr) IsSameRPCServerExist(rpcAddr string) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()

	for _, v := range s.remoteClients {
		if v.clientRPCAddr == rpcAddr {
			return true
		}
	}

	return false
}

func (s *SameTypeClientMgr) AddRPCServer(client *GrpcClient) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.remoteClients[client.serverID] = client
}

func (s *SameTypeClientMgr) DeleteRPCServerById(serverId string) {
	s.lock.Lock()
	defer s.lock.Unlock()

}

func (s *SameTypeClientMgr) deleteServerWithoutLock(serverId string) {
	client := s.remoteClients[serverId]
	if client != nil {
		delete(s.remoteClients, serverId)
		client.GetClientConn().Close()
	}
}
