package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"releaseband/internal/service"

	"github.com/gorilla/mux"
)

type Handler struct {
	service *service.Service
}

func New(serv *service.Service) *Handler {
	return &Handler{service: serv}
}
func JSONResponse(w http.ResponseWriter, code int, output interface{}) {
	response, err := json.Marshal(output)
	if err != nil {
		slog.Error("response marshal err", "error", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err = w.Write(response)
	if err != nil {
		slog.Error("write response err", "error", err)
	}
}

type ErrorResponse struct {
	Error string
}

// CalculateGame godoc
// @Summary     Return game result
// @Description Return game result
// @Param id   path string true "Game ID"
// @Produce     json
// @Success     200               {object} domain.Result
// @Router      /game/{id}/calculate [get]
func (h *Handler) CalculateGame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		JSONResponse(w, http.StatusBadRequest, ErrorResponse{Error: "empty id game"})
		return
	}
	result, err := h.service.GetCalculateDate(id)
	if err != nil {
		slog.Error("error in CalculateGame", "error", err)
		JSONResponse(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	JSONResponse(w, http.StatusOK, result)
}
