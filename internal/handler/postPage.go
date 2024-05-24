package handler

import (
	"fmt"
	"forum/internal/models"
	"net/http"
	"strconv"
	"time"
)

func (h *Handler) PostPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/post/" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return

	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if id == 0 || err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
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

	post, err := h.Service.GetOnePost(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}
	categories, err := h.Service.PostSer.GetCategories()
	comments, err := h.Service.GetAllComment(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return

	}
	switch r.Method {
	case http.MethodPost:
		comment := r.FormValue("comment")
		date := time.Now().Format("January 2, 2006 15:04:05")
		if err := h.Service.Comment.CreateComment(comment, user.Name, user.ID, id, date); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, r.URL.Path+fmt.Sprintf("/?id=%d", id), http.StatusSeeOther)
	case http.MethodGet:
		info := struct {
			Post     models.Post
			User     models.User
			Comments []models.Comment
			Category []models.Category
		}{
			Post:     post,
			User:     user,
			Comments: comments,
			Category: categories,
		}
		if err := h.Tmp.ExecuteTemplate(w, "postPage.html", info); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return

	}

}
