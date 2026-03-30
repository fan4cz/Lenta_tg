package main

import (
	"log"

	"go.mod/internal/config"
	container "go.mod/internal/infrastructure/Container"
)

func main() {
	cfg, err := config.MustLoad()
	if err != nil {
		log.Fatalf("failed to load %s", err)
	}

	container, err := container.New(cfg)
	if err != nil {
		log.Fatalf("di containre not run %s", err)
	}

	
}