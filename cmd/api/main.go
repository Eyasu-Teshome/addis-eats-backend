package main

import (
	"log"
	"net/http"
	"os"
	"time"
	"github.com/eyasu-teshome/addis-eats-api/internal/menu"
)

type application struct {
	logger *log.Logger
	dishes menu.DishReader
}



func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags)
	store := menu.NewMemoryStore()
	app := &application{
		logger: logger,
		dishes: store, // Satisfies DishReader interface implicitly
	}

	addr := os.Getenv("API_ADDR")
	if addr == "" {
		addr = ":8080"
	}

	srv := &http.Server{
		Addr:         addr,
		Handler:      app.routes(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  time.Minute,
	}

	logger.Printf("starting server on %s", addr)
	logger.Fatal(srv.ListenAndServe())
}