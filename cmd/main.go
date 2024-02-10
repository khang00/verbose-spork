package main

import (
	"fmt"
	"github.com/khang00/verbose-spork/internal/handler"
	"net/http"
)

func main() {
	setupHandler()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Server listening on port 8080")
	}
}

func setupHandler() {
	http.HandleFunc("/health", handler.HealthHandler)
}
