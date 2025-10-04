package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("GET /", serveIndex)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal((err))
	}
}

func serveIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
