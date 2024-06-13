package handler

import (
	"fmt"
	"forum/internal/models"
	"net/http"
	"time"
)

func (h *Handler) postCreate(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/post/create" {
		h.ErrorPage(w, http.StatusNotFound)
		return
	}
	userValue := r.Context().Value("user")
	if userValue == nil {
		h.ErrorPage(w, http.StatusUnauthorized)
		return
	}

	user, ok := userValue.(models.User)
	if !ok {
		h.ErrorPage(w, http.StatusUnauthorized)
		return
	}
	if !user.IsAuth {
		h.ErrorPage(w, http.StatusUnauthorized)
		return
	}
	categories, err := h.Service.PostSer.GetCategories()
	if err != nil {
		fmt.Println(err.Error())
		h.ErrorPage(w, http.StatusInternalServerError)
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
			CreateDate: time.Now().Format("January 2, 2006 15:04:05"),
		}); err != nil {
			fmt.Println(err.Error())
			h.ErrorPage(w, http.StatusBadRequest)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)

	case http.MethodGet:
		if err := h.Tmp.ExecuteTemplate(w, "postCreate.html", categories); err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return

		}
	default:
		h.ErrorPage(w, http.StatusMethodNotAllowed)
		return
	}

}
