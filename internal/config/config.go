package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
)

type Config struct {
	IsDebug             *bool `yaml:"is_debug"`
	Listen  struct {
		Type            string `yaml:"type" env-default:"port"`
		BindIP          string `yaml:"bind_ip" env-default:"localhost"`
		Port            string `yaml:"port" env-default:"8080"`
	}
	MySql struct {
		Host            string `yaml:"host" env-required:"true"`
		Port            string `yaml:"port" env-required:"true"`
		Username        string `yaml:"username"`
		Password        string `yaml:"password"`
		Database        string `yaml:"database" env-required:"true"`
	} `yaml:"mysql" env-required:"true"`
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
		}
	})
	return instance
}