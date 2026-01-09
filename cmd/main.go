package main

import (
	"hrms.local/infra/api"
	"hrms.local/infra/api/config"
)

func main() {
	cfg := config.LoadConfig()
	server := api.NewServer(cfg)
	server.StartServer()
}
