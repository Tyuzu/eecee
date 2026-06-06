package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write([]byte("Hello, HTTPS!"))
	})

	router.GET("/health", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write([]byte("Healty spelling!"))
	})

	router.GET("/status", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Write([]byte("200 ok!"))
	})

	server := &http.Server{
		Addr:    ":4000",
		Handler: router,
	}

	log.Println("HTTPS server listening on https://localhost:4000")

	err := server.ListenAndServeTLS("server.crt", "server.key")
	if err != nil {
		log.Fatal(err)
	}
}
