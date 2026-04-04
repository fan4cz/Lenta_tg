package container

import (
	handler "go.mod/internal/Handler"
	"go.mod/internal/config"
)

type Container struct {
	Config *config.Config

	ExampleHandler *handler.ExampleHandler 
}

func New(cfg *config.Config) (*Container, error) {
	c := &Container{}

	c.Config = cfg

	c.ExampleHandler = handler.NewExampleHandler()

	return &Container{}, nil
}