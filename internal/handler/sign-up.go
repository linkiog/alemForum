package handler

import (
	"fmt"
	"forum/internal/models"
	"net/http"
)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/signUp" {
		h.ErrorPage(w, http.StatusNotFound)
		return
	}
	switch r.Method {
	case http.MethodPost:
		email := r.FormValue("email")
		userName := r.FormValue("name")
		passw := r.FormValue("passw")
		rPassw := r.FormValue("rPassw")
		err := h.Service.Auth.CreateUser(models.User{Email: email, Name: userName, Password: passw, RPassw: rPassw})
		if err != nil {
			info := models.User{
				Email:    email,
				Name:     userName,
				Password: passw,
				RPassw:   rPassw,
				Error:    err,
			}
			fmt.Println(err)

			if err := h.Tmp.ExecuteTemplate(w, "sign-up.html", info); err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return

			}
			return
		}
		info := models.User{
			Status: "You have successfully registered",
			Email:  email,
		}
		if err := h.Tmp.ExecuteTemplate(w, "sign-in.html", info); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

	case http.MethodGet:
		if err := h.Tmp.ExecuteTemplate(w, "sign-up.html", nil); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	default:
		h.ErrorPage(w, http.StatusMethodNotAllowed)
		return

	}
}
