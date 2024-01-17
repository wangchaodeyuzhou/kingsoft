package cluserRedis

import "time"

type Redis struct {
	Cluster       bool
	Addrs         []string
	Username      string
	Password      string
	DB            uint32
	ReadTimeout   time.Duration
	WriteTimeout  time.Duration
	PoolTimeout   time.Duration
	MinIdleConns  uint32
	PoolSize      uint32
	KeyPrefix     string
	Retries       uint32
	RetryInterval time.Duration
}

func DefaultRedis() *Redis {
	return &Redis{
		Cluster:       false,
		ReadTimeout:   time.Second * 3,
		WriteTimeout:  time.Second * 3,
		PoolTimeout:   time.Second * 5,
		MinIdleConns:  5,
		PoolSize:      100,
		RetryInterval: time.Second * 3,
		Retries:       3,
	}
}
