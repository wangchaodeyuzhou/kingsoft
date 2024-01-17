package cluserRedis

import (
	"context"
	"log/slog"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	gomyredis "github.com/go-redis/redis/v8"
)

var mock *miniredis.Miniredis

func TestNew(t *testing.T) {
	// Mock redis.
	m, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	t.Log("mock redis", m.Addr())
	mock = m

	cfg := DefaultRedis()
	cfg.Addrs = []string{m.Addr()}

	ctx := context.Background()
	redisClient, err := New(ctx, cfg)
	if err != nil {
		t.Error("failed to New", err)
	}

	result, err := redisClient.Ping(ctx).Result()
	if err != nil {
		t.Error("failed to Ping", "err", err)
	}

	t.Log("redis ping done", result)
}

var lua = gomyredis.NewScript(`
local uidKey = KEYS[1]

local redisGatewayID = ARGV[1]
local gatewayID = ARGV[2]

local redisOnline = ARGV[3]
local online = ARGV[4]

local redisOfflineMilli = ARGV[5]
local nowMilli = ARGV[6]

local redisOnlineMilli = ARGV[7]

local oldGatewayID = redis.call("HGET", uidKey, redisGatewayID)
if oldGatewayID == gatewayID
then
	redis.call("HSET", uidKey, redisOnline, online, redisOnlineMilli, 222, redisOfflineMilli, nowMilli)
	return 1
else
	return 0
end
`)

func TestLua(t *testing.T) {
	// Mock redis.
	m, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	t.Log("mock redis", m.Addr())

	cfg := DefaultRedis()
	cfg.Addrs = []string{m.Addr()}

	ctx := context.Background()
	redisClient, err := New(ctx, cfg)
	if err != nil {
		t.Error("failed to New", err)
	}
	testLua(context.Background(), redisClient)
}

func testLua(ctx context.Context, client *Client) {
	key := "KGS:USER:1000"
	gatewayID := "g-200"

	ret, err := lua.Run(ctx, client.Scripter(), []string{key},
		"srv:gateway", gatewayID,
		"on", false,
		"onms", time.Now().UnixMilli(),
		"offms").Int()
	if err != nil {
		slog.Warn("failed to run lua on redis", "err", err)
		return
	}

	slog.Info("run lua done", "result", ret)
}
