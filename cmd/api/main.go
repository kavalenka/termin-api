package main

import (
	"log"
	"net/http"
	"termin-api/api/router"
	"time"
)

func main() {
	r := router.New()

	s := &http.Server{
		Addr:         ":3000",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Fatal(s.ListenAndServe())
}
