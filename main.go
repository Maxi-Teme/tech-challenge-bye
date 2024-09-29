package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type ByeHandler struct {
	env string
}

func (byeHandler ByeHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("Got request %v\n", req)
	fmt.Fprintf(w, "bye from environment: \"%v\"\n", byeHandler.env)
}

func health(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "OK")
}

func main() {
	var PORT = os.Getenv("PORT")
	if len(PORT) < 1 {
		log.Fatal("ERROR: please provide environment variable 'PORT'")
	}
	var ENV = os.Getenv("ENV")
	if len(ENV) < 1 {
		log.Fatal("ERROR: please provide environment varialbe 'ENV'")
	}

	byeHandler := ByeHandler{env: ENV}

	http.Handle("/bye", byeHandler)
	http.HandleFunc("/health", health)

	var addr = fmt.Sprintf(":%v", PORT)

	fmt.Printf("HTTP server listening on %v\n", addr)
	http.ListenAndServe(addr, nil)
}
