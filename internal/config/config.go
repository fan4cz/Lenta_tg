package config

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	HTTPServer
}

type HTTPServer struct {
	Address string `env:"address" env-default:"localhost:8085"`
}


func MustLoad() (*Config, error){
	cfg := &Config{}

    err := cleanenv.ReadEnv(cfg)
    if err != nil {
        return nil, err
    }

    return cfg, nil
}