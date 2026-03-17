package main

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
)

//go:embed static/index.html
var indexHTML []byte

func startServer(addr string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Write(indexHTML)
	})
	fmt.Printf("Serving on http://localhost%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
