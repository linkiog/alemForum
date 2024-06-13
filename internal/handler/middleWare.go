package handler

import (
	"context"
	"forum/internal/models"
	"net/http"
	"time"
)

func (h *Handler) middleWare(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		t, err := r.Cookie("token")
		if err != nil {
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "user", models.User{})))
			return

		}
		user, err = h.Service.GetUserByToken(t.Value)
		if err != nil {
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "user", models.User{})))
			return
		}
		if user.Token_duration.Before(time.Now()) {
			if err := h.Service.DeleteToken(t.Value); err != nil {
				h.ErrorPage(w, http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/sigIn", http.StatusSeeOther)
			return

		}
		user.IsAuth = true
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "user", user)))
	}
}
