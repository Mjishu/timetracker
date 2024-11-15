package main

import (
	"log"
	"net/http"
)

func main() {
	server := &TaskServer{}
	log.Fatal(http.ListenAndServe(":5000", server))
}
