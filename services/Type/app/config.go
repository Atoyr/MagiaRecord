package app

import (
	"fmt"
	"sync"
)

type Config struct {
	Name             string
	HttpPort         uint
	DatabaseServer   string
	DatabasePort     uint
	DatabaseUser     string
	DatabasePassword string
	UsingDatabase    string
	LogLevel         int
	AcceccLog        bool
}

var (
	once sync.Once

	c Config
)

func NewConfig() *Config {
	once.Do(func() {
		c = Config{}
	})

	return &c
}

func (c *Config) ConnString() string {
	return fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", c.DatabaseServer, c.DatabaseUser, c.DatabasePassword, c.DatabasePort, c.UsingDatabase)
}
