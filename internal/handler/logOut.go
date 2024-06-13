package handler

import (
	"fmt"
	"net/http"
)

func (h *Handler) logOut(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/logOut" {
		h.ErrorPage(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		h.ErrorPage(w, http.StatusNotFound)
		return
	}
	c, err := r.Cookie("token")
	if err != nil {
		fmt.Println(err.Error())
		h.ErrorPage(w, http.StatusInternalServerError)
		return
	}
	if err := h.Service.DeleteToken(c.Value); err != nil {
		fmt.Println(err.Error())
		h.ErrorPage(w, http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:   "token",
		MaxAge: -1,
	})
	http.Redirect(w, r, "/", http.StatusSeeOther)

}
