package container

import "go.mod/internal/config"

type Container struct {
	Config *config.Config
}

func New(cfg *config.Config) (*Container, error) {
	c := &Container{}

	c.Config = cfg
}