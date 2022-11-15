package main

import (
	"fmt"
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
	mux.HandleFunc("/v1/message", sendHandler)
	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		return
	}
}

func sendHandler(w http.ResponseWriter, r *http.Request) {
	var method = r.Method
	var err error
	var res string
	switch method {
	case http.MethodGet:
		res, err = "Get messages", nil
	case http.MethodPost:
		res, err = "Produce message", nil
	default:
		res, err = fmt.Sprintf("method %s not allowed", method), nil
	}
	if err != nil {
		w.Write([]byte("Непредвиденная ошибка!"))
		return
	}

	w.Write([]byte(res))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World"))
}
