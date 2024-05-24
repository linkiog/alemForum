package handler

import (
	"forum/internal/models"
	"net/http"
)

func (h *Handler) myPosts(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/myPosts" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
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
	myPost, err := h.Service.PostSer.GetMyPosts(user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	categories, err := h.Service.PostSer.Category()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	info := struct {
		AllPosts []models.Post
		User     models.User
		Category []models.Category
	}{
		AllPosts: myPost,
		User:     user,
		Category: categories,
	}

	if err := h.Tmp.ExecuteTemplate(w, "homePage.html", info); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}