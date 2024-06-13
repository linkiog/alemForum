package handler

// import (
// 	"forum/internal/models"
// 	"net/http"
// 	"strconv"
// 	"time"
// )

// func (h *Handler) editPost(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/post/edit/" {
// 		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
// 		return
// 	}
// 	if r.Method != http.MethodPost {
// 		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
// 		return
// 	}
// 	userValue := r.Context().Value("user")
// 	if userValue == nil {
// 		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
// 		return
// 	}

// 	user, ok := userValue.(models.User)
// 	if !ok {
// 		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
// 		return
// 	}
// 	if !user.IsAuth {
// 		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
// 		return

// 	}
// 	id, err := strconv.Atoi(r.URL.Query().Get("id"))
// 	if err != nil || id == 0 {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	title := r.FormValue("title")
// 	content := r.FormValue("content")
// 	category := r.Form["categories"]
// 	if err := h.Service.PostSer.EditPost(models.Post{
// 		IdAuth:     user.ID,
// 		IdPost:     id,
// 		Title:      title,
// 		Content:    content,
// 		Category:   category,
// 		CreateDate: time.Now(),
// 	}); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	http.Redirect(w, r, "/", http.StatusSeeOther)

// }
