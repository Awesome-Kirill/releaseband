package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"releaseband/internal/domain"
	"releaseband/internal/service"
)

type Handler struct {
	service *service.Service
}

func New(service *service.Service) *Handler {
	return &Handler{service: service}
}
func (h *Handler) CalculateGame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprintf(w, "error get id:")
	}
	result, err := h.service.GetCalculateDate(id)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprintf(w, "error: %v\n", err)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "calculate: %v\n", result)
}

func (h *Handler) CreateReels(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprintf(w, "error get id:")
	}

	var reels domain.Reels
	err := json.NewDecoder(r.Body).Decode(&reels)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprintf(w, "eror: %v\n", err)

	}
	defer r.Body.Close()

	err = h.service.CreateOrUpdate(id, &domain.GameDate{Reels: &reels})
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprintf(w, "eror: %v\n", err)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "reels : %v, id: %v", reels, vars["id"])
}

func (h *Handler) CreatePayouts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprintf(w, "error get id:")
	}

	var pay domain.Payouts
	err := json.NewDecoder(r.Body).Decode(&pay)
	if err != nil {
		fmt.Fprintf(w, "eror: %v\n", err)
	}
	defer r.Body.Close()

	err = h.service.CreateOrUpdate(id, &domain.GameDate{Payouts: &pay})
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprintf(w, "eror: %v\n", err)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Payouts: %v, id: %v", pay, vars["id"])
}

func (h *Handler) CreateRLines(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprintf(w, "error get id:")
	}

	var win domain.WinLines
	err := json.NewDecoder(r.Body).Decode(&win)
	if err != nil {
		fmt.Fprintf(w, "eror: %v\n", err)
	}
	defer r.Body.Close()

	err = h.service.CreateOrUpdate(id, &domain.GameDate{WinLines: &win})
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprintf(w, "eror: %v\n", err)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "win : %v, id: %v", win, vars["id"])
}
