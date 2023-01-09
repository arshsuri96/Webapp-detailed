package handler

import (
	"log"
	"net/http"
)

type Bye struct {
	l *log.Logger
}

func goodbye(l *log.Logger) *Bye {
	return &Bye{l}
}

func (g *Bye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Byee"))
}
