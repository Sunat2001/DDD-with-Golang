package main

import (
	"CRUD/pkg/config"
	"CRUD/pkg/database"
	"CRUD/pkg/routes"
	"context"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()

	routes.Api(r)

	database.InitDB()

	initServer(r, cfg)
}

func initServer(r *mux.Router, cfg *config.Config) {
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1" + cfg.HttpAddr,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		log.Println("Server started on port " + cfg.HttpAddr)
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)
	<-signals

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	log.Println("Shutting down server...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
