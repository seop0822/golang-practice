package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("receivced request")
		fmt.Fprintf(w,"Hello DOcker!!")
	})

	log.Println("start server")
	server := & http.Server{
		Addr: ":8080",
	}

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}