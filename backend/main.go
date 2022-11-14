package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprintf(w, "Hi there, I love %s!", name)
}

func main() {
	http.HandleFunc("/sayhello", handler)
	log.Fatal(http.ListenAndServe(":8090", nil))
}
