package app

import (
	"github.com/alisyahbana/tax-calculator/pkg/common/config"
	"github.com/alisyahbana/tax-calculator/pkg/common/env"
)

type Config struct {
	Version string `json:"version"`
	Port    int    `json:"port"`
}

var cfg *Config

func init() {
	config.LoadConfiguration(&cfg, "app", env.GetEnv())
}

func GetConfig() *Config {
	return cfg
}
