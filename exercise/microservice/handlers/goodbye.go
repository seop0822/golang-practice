package handlers

import (
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func (g *Goodbye) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	panic("implement me")
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTp(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("BYeee"))
}