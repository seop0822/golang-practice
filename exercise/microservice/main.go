package main

import (
	"exercise/microservice/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.Goodbye{l}

	sm := http.NewServeMux()
	sm.Handle("/",hh)
	sm.Handle("/goodbye",gh)

	http.ListenAndServe(":9000", nil)
}
