package main

import (
	"fmt"
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
	http.HandleFunc("/health", healthHandler)
}
