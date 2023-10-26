package main

import (
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"releaseband/internal/handler"
	"releaseband/internal/service"
	"releaseband/internal/storage"
	"time"
)

func main() {

	handler := handler.New(service.New(storage.New()))
	r := mux.NewRouter()
	// Add your routes as needed
	r.HandleFunc("/v1/game/{id}/calculate", handler.CalculateGame).Methods("GET")

	r.HandleFunc("/v1/game/{id}/reels", handler.CreateReels).Methods("POST")
	r.HandleFunc("/v1/game/{id}/payouts", handler.CreatePayouts).Methods("POST")
	r.HandleFunc("/v1/game/{id}/lines", handler.CreateRLines).Methods("POST")

	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	srv.Shutdown(ctx)

	log.Println("shutting down")
	os.Exit(0)
}
