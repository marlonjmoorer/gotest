package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	r.BasicAuth()
	fmt.Fprint(w, "Hi there")
}

func main() {

	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "8000"
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
