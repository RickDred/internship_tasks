package main

import (
	"fmt"
	"net/http"

	"github.com/RickDred/internship_tasks/tree/third_task/config"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered error:", r)
		}
	}()

	cfg, err := config.Decode("config.yml")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("everything is okay"))
	})

	addr := fmt.Sprintf("%v:%v", cfg.Server.Host, cfg.Server.Port)
	fmt.Printf("Starting on http://%v\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}
