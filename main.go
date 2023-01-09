package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/arshsuri96/test/handler"
)

func main() {
	l := log.New(os.Stdout, "producit-api", log.LstdFlags)

	hh := handler.NewHello(l)
	gb := handler.Bye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/bye", gb)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	s.ListenAndServe()

}
