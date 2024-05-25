package handler

import (
	"fmt"
	"forum/internal/models"
	"net/http"
	"time"
)

func (h *Handler) postCreate(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/post/create" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	userValue := r.Context().Value("user")
	if userValue == nil {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	user, ok := userValue.(models.User)
	if !ok {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	if !user.IsAuth {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	categories, err := h.Service.PostSer.GetCategories()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}
	switch r.Method {
	case http.MethodPost:
		title := r.FormValue("title")
		content := r.FormValue("content")
		categories := r.Form["categories"]
		if err := h.Service.CreatePost(models.Post{
			IdAuth:     user.ID,
			Author:     user.Name,
			Title:      title,
			Content:    content,
			Category:   categories,
			CreateDate: time.Now(),
		}); err != nil {
			fmt.Println(err.Error())
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)

	case http.MethodGet:
		if err := h.Tmp.ExecuteTemplate(w, "postCreate.html", categories); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return

		}
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

}
