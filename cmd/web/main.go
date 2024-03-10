package main

import (
	"github.com/RickDred/internship_tasks/tree/sixth_task/config"
	"github.com/RickDred/internship_tasks/tree/sixth_task/internal/app"
)

func main() {
	cfg, err := config.Decode("config.yml")
	if err != nil {
		panic(err)
	}

	app.StartServer(cfg)
}
