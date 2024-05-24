package handler

import (
	"fmt"
	"net/http"
)

func (h *Handler) logOut(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/logOut" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	c, err := r.Cookie("token")
	if err != nil {
		http.Error(w, "func logOut cookie:"+http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if err := h.Service.DeleteToken(c.Value); err != nil {
		http.Error(w, fmt.Sprintf("func logOut deleteToken%w"+err.Error()), http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:   "token",
		MaxAge: -1,
	})
	http.Redirect(w, r, "/", http.StatusSeeOther)

}
