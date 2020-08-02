package main

import (
	"exercise/restful/myapp"
	"net/http"
)


func main() {
	http.ListenAndServe(":3000", myapp.NewHandler())
}