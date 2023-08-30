package grpcnet

import (
	"context"
	"git.kingsoft.go/base/grpcnet/state"
	promgrpc "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/pkg/errors"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type GrpcClient struct {
	conn          *grpc.ClientConn // connection
	gClient       any              // GRPC connection
	clientRPCAddr string           // rpc client listen address
	serverName    string           // server name
	serverID      string           // server id
	healthState   int32
	tags          []string
	meta          map[string]string
}

func (g *GrpcClient) GetClientConn() *grpc.ClientConn {
	return g.conn
}

func (g *GrpcClient) GetClientRPCAddr() string {
	return g.clientRPCAddr
}

func (g *GrpcClient) GetServerId() string {
	return g.serverID
}

func (g *GrpcClient) GetServerName() string {
	return g.serverName
}

func (g *GrpcClient) GetGrpcClient() any {
	return g.gClient
}

func (g *GrpcClient) ContainsTag(tag string) bool {
	for _, s := range g.tags {
		if s == tag {
			return true
		}
	}
	return false
}

func Dial(addr string, callback state.ClosedCallBack) (*grpc.ClientConn, error) {
	grpcOpts := []grpc.DialOption{
		grpc.WithChainUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		grpc.WithChainStreamInterceptor(promgrpc.StreamClientInterceptor),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithReturnConnectionError(),
		grpc.WithBlock(),
		grpc.WithStatsHandler(&state.ClientHandler{CallBack: callback}),
	}

	promgrpc.EnableHandlingTimeHistogram()

	ctx, cancelCtx := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancelCtx()

	slog.Info("dial context", "addr", addr)

	return grpc.DialContext(ctx, addr, grpcOpts...)
}

// NewRPCClient creates a new RPC client
func NewRPCClient(addr string, serverName string, sid string, meta map[string]string, tags []string) (*GrpcClient, error) {
	conn, err := Dial(addr, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect")
	}

	var client any
	switch serverName {
	case "DDD":
		// todo: implement new detail client
		client = nil
	}

	return &GrpcClient{
		conn:          conn,
		gClient:       client,
		clientRPCAddr: addr,
		serverID:      sid,
		serverName:    serverName,
		meta:          meta,
		tags:          tags,
	}, nil
}
