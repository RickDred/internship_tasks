package main

import (
	"github.com/RickDred/internship_tasks/blob/fourth_task/internal/app"
	"github.com/RickDred/internship_tasks/blob/fourth_task/internal/config"
)

func main() {
	cfg, err := config.Decode("config.yml")
	if err != nil {
		panic(err)
	}

	app.StartServer(cfg)
}
