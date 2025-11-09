package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/denerkrebs/WorkflowEngine/config"
	"github.com/denerkrebs/WorkflowEngine/internal/infrastructure/container"
	"github.com/denerkrebs/WorkflowEngine/internal/infrastructure/http/router"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	c, err := container.New(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize container: %v", err)
	}

	r := router.NewRouter(c)

	port := cfg.GetServerPort()

	fmt.Println("Starting")
	fmt.Printf("Server running at http://localhost%s\n", port)

	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal("Error while starting server:", err)
	}
}
