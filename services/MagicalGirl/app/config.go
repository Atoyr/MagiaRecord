package app

import (
	"fmt"
	"sync"
)

type Config struct {
	Server   string
	Port     uint16
	User     string
	Password string
	Database string
}

var (
	once sync.Once

	c Config
)

func New() Config {
	once.Do(func() {
		c = Config{}
	})

	return c
}

func (c *Config) ConnString() string {
	return fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", c.Server, c.User, c.Password, c.Port, c.Database)
}
