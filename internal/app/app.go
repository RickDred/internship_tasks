package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RickDred/internship_tasks/tree/sixth_task/config"
	"github.com/RickDred/internship_tasks/tree/sixth_task/pkg/router"
)

func StartServer(cfg *config.Config) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered error: %v\n", r)
		}
	}()

	r := router.NewRouter()

	SetRoutes(r, cfg)

	addr := fmt.Sprintf("%v:%v", cfg.Server.Host, cfg.Server.Port)
	fmt.Printf("Starting on http://%v\n", addr)

	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal(err)
	}
}
