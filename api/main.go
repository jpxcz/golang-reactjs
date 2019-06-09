package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the API!")
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("API listening on port 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
