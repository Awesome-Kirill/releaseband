package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"releaseband/internal/service"
)

type Handler struct {
	service *service.Service
}

func New(service *service.Service) *Handler {
	return &Handler{service: service}
}
func JSONResponse(w http.ResponseWriter, code int, output interface{}) {
	response, _ := json.Marshal(output)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

type ErrorResponse struct {
	Error string
}

func (h *Handler) CalculateGame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		JSONResponse(w, http.StatusBadRequest, ErrorResponse{Error: "empty id game"})
		return
	}
	result, err := h.service.GetCalculateDate(id)
	if err != nil {
		JSONResponse(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	JSONResponse(w, http.StatusOK, result)
	return
}
