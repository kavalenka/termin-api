package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"termin-api/config"
	"termin-api/internal/app/router"
	"time"
)

func main() {
	dbConfig := config.NewDB()
	dbConfigStr := fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
	db, _ := sql.Open("postgres", dbConfigStr)

	r := router.New(db)

	s := &http.Server{
		Addr:         ":3000",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Fatal(s.ListenAndServe())
}
