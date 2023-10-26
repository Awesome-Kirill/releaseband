package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"releaseband/internal/handler"
	"releaseband/internal/service"
	"releaseband/internal/storage"
	"time"

	"github.com/gorilla/mux"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	reg := prometheus.NewRegistry()
	reg.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
	)

	handler := handler.New(service.New(storage.New()))
	r := mux.NewRouter()
	// Add your routes as needed
	r.HandleFunc("/v1/game/{id}/calculate", handler.CalculateGame).Methods("GET")

	r.HandleFunc("/v1/game/{id}/reels", handler.CreateReels).Methods("POST")
	r.HandleFunc("/v1/game/{id}/payouts", handler.CreatePayouts).Methods("POST")
	r.HandleFunc("/v1/game/{id}/lines", handler.CreateRLines).Methods("POST")

	r.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg}))
	srv := &http.Server{
		Addr:         "0.0.0.0:8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		logger.Info("server start")
		if err := srv.ListenAndServe(); err != nil {
			logger.Error("ListenAndServe error:", err)
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	<-c

	logger.Info("shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	srv.Shutdown(ctx)
}
