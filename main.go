package main

import (
	"fmt"
	"net/http"

	"github.com/RickDred/internship_tasks/tree/third_task/config"
)

func main() {
	// recover from panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered error:", r)
		}
	}()

	// read config file
	cfg, err := config.Decode("config.yml")
	if err != nil {
		panic(err)
	}

	// set up handler
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("everything is okay"))
	})

	// set up addr
	addr := fmt.Sprintf("%v:%v", cfg.Server.Host, cfg.Server.Port)
	fmt.Printf("Starting on http://%v\n", addr)

	// listen port
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}
