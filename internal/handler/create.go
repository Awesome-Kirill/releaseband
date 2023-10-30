package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"releaseband/internal/domain"

	"github.com/gorilla/mux"
)

type OkResponse struct{}

// CreateReels godoc
// @Summary     Create reels
// @Description Create reels
// @Param id   path string true "Game ID"
// @Param		some_id	body		domain.Reels		true	"Some ID"
// @Produce     json
// @Success 200 {object} OkResponse
// @Router      /game/{id}/reels [post]
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
		slog.Error("error in CreateReels", "error", err)
		JSONResponse(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	defer r.Body.Close()

	err = h.service.CreateReels(id, &reels)
	if err != nil {
		slog.Error("error in CreateReels", "error", err)
		JSONResponse(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	JSONResponse(w, http.StatusOK, OkResponse{})
}

// CreatePayouts godoc
// @Summary     Create payouts
// @Description Create payouts
// @Param id   path string true "Game ID"
// @Param		some_id	body		domain.Payouts		true	"Some ID"
// @Produce     json
// @Success 200 {object} OkResponse
// @Router      /game/{id}/payouts [post]
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
		slog.Error("error in CreatePayouts", "error", err)
		JSONResponse(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	defer r.Body.Close()

	err = h.service.CreatePayouts(id, pay)
	if err != nil {
		slog.Error("error in CreatePayouts", "error", err)
		JSONResponse(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	JSONResponse(w, http.StatusOK, OkResponse{})
}

// CreateLines godoc
// @Summary     Create lines
// @Description Create lines
// @Param id   path string true "Game ID"
// @Param		some_id	body		domain.Lines		true	"Some ID"
// @Produce     json
// @Success 200 {object} OkResponse
// @Router      /game/{id}/lines [post]
func (h *Handler) CreateLines(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		slog.Error("error in CreateLines", "error", "empty id game")
		JSONResponse(w, http.StatusBadRequest, ErrorResponse{Error: "empty id game"})
		return
	}

	var lines domain.Lines
	err := json.NewDecoder(r.Body).Decode(&lines)
	if err != nil {
		slog.Error("error in CreateLines", "error", err)
		JSONResponse(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	defer r.Body.Close()

	err = h.service.CreateLines(id, lines)
	if err != nil {
		slog.Error("error in CreateLines", "error", err)
		JSONResponse(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	JSONResponse(w, http.StatusOK, OkResponse{})
}
