package main

import (
	"log"
	"net/http"

	"server/auth"
)

func main() {
	http.HandleFunc("/", auth.Auth)
	log.Println("listening (port: 8080)")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
