package handler

import (
	"encoding/json"
	"net/http"
	"releaseband/internal/domain"

	"github.com/gorilla/mux"
)

func (h *Handler) CreateReels(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		JSONResponse(w, http.StatusBadRequest, ErrorResponse{Error: "empty id game"})
		return
	}

	var reels domain.Reels
	err := json.NewDecoder(r.Body).Decode(&reels)
	if err != nil {
		JSONResponse(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	defer r.Body.Close()

	err = h.service.CreateReels(id, reels)
	if err != nil {
		JSONResponse(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	JSONResponse(w, http.StatusOK, struct{}{})
}

func (h *Handler) CreatePayouts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		JSONResponse(w, http.StatusBadRequest, ErrorResponse{Error: "empty id game"})
		return
	}

	var pay domain.Payouts
	err := json.NewDecoder(r.Body).Decode(&pay)
	if err != nil {
		JSONResponse(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	defer r.Body.Close()

	err = h.service.CreatePayouts(id, pay)
	if err != nil {
		JSONResponse(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	JSONResponse(w, http.StatusOK, struct{}{})
}

func (h *Handler) CreateLines(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		JSONResponse(w, http.StatusBadRequest, ErrorResponse{Error: "empty id game"})
		return
	}

	var lines domain.Lines
	err := json.NewDecoder(r.Body).Decode(&lines)
	if err != nil {
		JSONResponse(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	defer r.Body.Close()

	err = h.service.CreateLines(id, lines)
	if err != nil {
		JSONResponse(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	JSONResponse(w, http.StatusOK, struct{}{})
}
