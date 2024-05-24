package handler

import (
	"forum/internal/models"
	"net/http"
)

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/signIn" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	switch r.Method {
	case http.MethodPost:
		email := r.FormValue("email")
		password := r.FormValue("password")
		token, expired, err := h.Service.CheckUserFormDb(models.User{
			Email:    email,
			Password: password,
		})
		if err != nil {
			info := models.User{
				Error:    err,
				Email:    email,
				Password: password,
			}
			if err := h.Tmp.ExecuteTemplate(w, "sign-in.html", info); err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return

			}
			return
		}
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   token,
			Path:    "/",
			Expires: expired,
		})
		http.Redirect(w, r, "/", http.StatusSeeOther)
	case http.MethodGet:
		if err := h.Tmp.ExecuteTemplate(w, "sign-in.html", nil); err != nil {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return

	}

}
