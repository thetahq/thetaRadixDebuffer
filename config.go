package main

import "fmt"

type Config struct {
	Port          int
	DebufferPort  int
	ServerAddress string
}

func loadConfig() Config {
	cfg := Config{
		Port: 80, // TODO: Replace with config loaded from ../
	}
	cfg.preCompute()
	return cfg
}

func (c *Config) preCompute() {
	c.ServerAddress = fmt.Sprintf("http://localhost:%d", c.Port)
}
