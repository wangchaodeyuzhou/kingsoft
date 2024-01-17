package cluserRedis

import (
	"context"
	"log/slog"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"github.com/pkg/errors"
)

type Client struct {
	iClient

	Scripter func() redis.Scripter

	// Distributed locks.
	RedSync *redsync.Redsync
}

type iClient interface {
	redis.Cmdable
	redis.UniversalClient
}

func New(ctx context.Context, cfg *Redis) (*Client, error) {
	ret := &Client{}

	if cfg.Cluster {
		slog.Info("new redis cluster", "addrs", cfg.Addrs)

		client := redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:        cfg.Addrs,
			Username:     cfg.Username,
			Password:     cfg.Password,
			ReadTimeout:  cfg.ReadTimeout,
			WriteTimeout: cfg.WriteTimeout,
			PoolSize:     int(cfg.PoolSize),
			MinIdleConns: int(cfg.MinIdleConns),
			PoolTimeout:  cfg.PoolTimeout,
		})

		ret.iClient = client
		ret.Scripter = func() redis.Scripter {
			return client
		}
	} else {
		slog.Info("new redis single", "addrs", cfg.Addrs)

		client := redis.NewClient(&redis.Options{
			Addr:         cfg.Addrs[0],
			Username:     cfg.Username,
			Password:     cfg.Password,
			ReadTimeout:  cfg.ReadTimeout,
			WriteTimeout: cfg.WriteTimeout,
			PoolSize:     int(cfg.PoolSize),
			MinIdleConns: int(cfg.MinIdleConns),
			PoolTimeout:  cfg.PoolTimeout,
		})

		ret.iClient = client
		ret.Scripter = func() redis.Scripter {
			return client
		}

		// Select db.
		ret.iClient.Do(ctx, "SELECT", cfg.DB)
	}

	// New redsync.
	pool := goredis.NewPool(ret.iClient)
	ret.RedSync = redsync.New(pool)

	return pingWithRetry(ctx, cfg, ret)
}

func pingWithRetry(ctx context.Context, cfg *Redis, client *Client) (*Client, error) {
	var retries uint32
	if cfg.Retries > 1 {
		retries = cfg.Retries
	} else {
		retries = 1
	}

	var triedSum uint32

	for {
		result, err := client.iClient.Ping(ctx).Result()
		if err != nil {
			slog.Warn("failed to ping redis", "err", err)

			if triedSum < retries {
				triedSum++
				slog.Info("retry ping", "tried_sum", triedSum, "total_retries", retries)
				time.Sleep(cfg.RetryInterval)
				continue
			}

			return nil, errors.Wrap(err, "redis not available")
		}

		slog.Info("ping redis success", "result", result)
		return client, nil
	}
}
