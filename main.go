package main

import (
	"log"
	"net/http"

	"go_server_sample/handler"
)

func main() {
	http.HandleFunc("/", handler.Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
