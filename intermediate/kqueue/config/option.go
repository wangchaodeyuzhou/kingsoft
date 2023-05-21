package config

import "git.shiyou.kingsoft.com/go/log"

type Option func(*Config)

// When the config changes, this method f will be called.
func WithReloadFunc(f func(any)) Option {
	return func(c *Config) {
		watch(c, f)
	}
}

// It is used to verify the correctness during config loading.
func WithVerifyFunc(f func(any) error) Option {
	return func(c *Config) {
		c.verifyFunc = f
	}
}

// It is required that the data struct and the fields in the config file are strictly matched.
func WithExactDataStruct() Option {
	return func(c *Config) {
		c.exactDataStruct = true
	}
}

func WithCheckMissingKey() Option {
	return func(c *Config) {
		c.checkMissingKey = true
	}
}

func WithLogLevel(level string) Option {
	return func(c *Config) {
		log.SetLevel(level)
	}
}

func WithDebug() Option {
	return func(c *Config) {
		c.debug = true
	}
}
