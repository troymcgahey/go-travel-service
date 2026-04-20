package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/troymcgahey/go-travel-service/internal/service"
)

type TodoHandler struct {
	service *service.TodoService
}

func NewTodoHandler(s *service.TodoService) *TodoHandler {
	return &TodoHandler{service: s}
}

func (h *TodoHandler) ListTools(w http.ResponseWriter, r *http.Request) {
	todos := h.service.List()
	writeJSON(w, http.StatusOK, todos)
}

func (h *TodoHandler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Customer    string
		Destination string
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
		return
	}

	//if req.Text == "" {
	//	writeJSON(w, http.StatusBadRequest, map[string]string{
	//		"erorr": "text is required",
	//	})
	//	return
	//}

	//todo := h.service.Create(req.Text)
	writeJSON(w, http.StatusCreated, "Booking Created")
}
