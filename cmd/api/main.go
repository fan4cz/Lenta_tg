package main

import (
	"log"

	"go.mod/internal/config"
)

func main() {
	cfg, err := config.MustLoad()
	if err != nil {
		log.Fatalf("failed to load %s", err)
	}
	
}