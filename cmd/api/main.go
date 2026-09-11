package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/azulgautam79/olx-api/internal/config"
	"github.com/azulgautam79/olx-api/internal/db"
	"github.com/azulgautam79/olx-api/internal/handlers"
)

func main() {

	cfg := config.MustLoad()

	_, err := db.Connect(cfg.DBUrl)
	if err != nil {
		log.Fatalf("main.db.connect: %v", err)
	}

	fmt.Println("database connected")
	fmt.Println("starting olx server...")

	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", handlers.Healthz)

	srv := http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      mux,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
	}

	log.Printf("Server is listening on %s", srv.Addr)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
