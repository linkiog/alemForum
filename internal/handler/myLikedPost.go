package handler

import (
	"fmt"
	"forum/internal/models"
	"net/http"
)

func (h *Handler) GetMyLikedPost(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/likedPosts" {
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
	myLikedPost, err := h.Service.PostSer.GetMyLikedPost(user.ID)
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	categories, err := h.Service.PostSer.Category()
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	info := struct {
		AllPosts []models.Post
		User     models.User
		Category []models.Category
	}{
		AllPosts: myLikedPost,
		User:     user,
		Category: categories,
	}

	if err := h.Tmp.ExecuteTemplate(w, "homePage.html", info); err != nil {
		fmt.Println(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

}
