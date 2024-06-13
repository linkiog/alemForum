package handler

import (
	"fmt"
	"forum/internal/models"
	"net/http"
)

func (h *Handler) ErrorPage(w http.ResponseWriter, status int) {
	errData := models.Error{Status: status, StatusText: http.StatusText(status)}
	w.WriteHeader(status)
	if err := h.Tmp.ExecuteTemplate(w, "error.html", errData); err != nil {
		fmt.Printf("error handler: execute: %s\n", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
