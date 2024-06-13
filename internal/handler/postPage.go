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
		h.ErrorPage(w, http.StatusNotFound)
		return

	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if id == 0 || err != nil {
		h.ErrorPage(w, http.StatusBadRequest)
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

	post, err := h.Service.GetOnePost(id)
	if err != nil {
		fmt.Println(err.Error())
		h.ErrorPage(w, http.StatusInternalServerError)
		return

	}
	categories, err := h.Service.PostSer.GetCategories()
	if err != nil {
		fmt.Println(err.Error())
		h.ErrorPage(w, http.StatusInternalServerError)
	}
	comments, err := h.Service.GetAllComment(id)
	if err != nil {
		fmt.Println(err.Error())
		h.ErrorPage(w, http.StatusInternalServerError)
		return

	}
	switch r.Method {
	case http.MethodPost:
		if !user.IsAuth {
			h.ErrorPage(w, http.StatusUnauthorized)
			return
		}
		comment := r.FormValue("comment")
		if comment == "" || len(comment) > 200 {
			h.ErrorPage(w, http.StatusBadRequest)
			return
		}
		date := time.Now().Format("January 2, 2006 15:04:05")
		if err := h.Service.Comment.CreateComment(comment, user.Name, user.ID, id, date); err != nil {
			fmt.Println(err.Error())
			h.ErrorPage(w, http.StatusInternalServerError)
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
			fmt.Println(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	default:
		h.ErrorPage(w, http.StatusMethodNotAllowed)
		return

	}

}
