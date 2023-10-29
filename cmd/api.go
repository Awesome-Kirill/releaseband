package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"releaseband/config"
	"releaseband/internal/handler"
	"releaseband/internal/metrics"
	"releaseband/internal/service"
	"releaseband/internal/storage"
	"time"

	"github.com/gorilla/mux"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

/*
Метрические данные должны включать в себя как минимум следующие метрики:

Общее количество обработанных запросов к API-endpoints;
Количество ошибок обработки HTTP запросов к API-endpoints;
Данные по времени обработки HTTP запросов к API-endpoint;
*/
// Логирование
// найминг
// тесты
// линтер
// общая архитетура
// валидация
func main() {

	cfg := config.New()
	fmt.Println(cfg.HTTPAddr)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	metricsMiddelware := metrics.New()

	handler := handler.New(service.New(storage.New()))

	r := mux.NewRouter()
	r.HandleFunc("/v1/game/{id}/calculate", handler.CalculateGame).Methods("GET")

	r.HandleFunc("/v1/game/{id}/reels", handler.CreateReels).Methods("POST")
	r.HandleFunc("/v1/game/{id}/payouts", handler.CreatePayouts).Methods("POST")
	r.HandleFunc("/v1/game/{id}/lines", handler.CreateLines).Methods("POST")
	r.Use(metricsMiddelware.After)
	r.Handle("/metrics", promhttp.Handler())

	srv := &http.Server{
		Addr:         "0.0.0.0:8080", //cfg.HTTPAddr
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
