package main

import (
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		return
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World"))
}
