package main

import (
	"net/http"
	"os"
	"log"
)

func main() {
	bindPort := os.Getenv("PORT")

	if bindPort == "" {
		bindPort = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	bindAddr := "localhost:"+bindPort
	log.Println("Application is running on", bindAddr)
	err := http.ListenAndServe(bindAddr, nil)
	if err != nil {
		panic(err)
	}
}
