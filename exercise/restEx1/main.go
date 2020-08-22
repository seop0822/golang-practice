package main

import (
	"exercise/restEx1/myapp"
	"net/http"
)


func main() {
	http.ListenAndServe(":3000", myapp.NewHandler())
}