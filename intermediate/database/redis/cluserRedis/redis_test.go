package cluserRedis

import (
	"context"
	"testing"

	"github.com/alicebob/miniredis/v2"
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
