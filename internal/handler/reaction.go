package handler

import (
	"fmt"
	"forum/internal/models"
	"net/http"
	"strconv"
)

func (h *Handler) reactionPost(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/reaction/post/" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return

	}
	if r.Method != http.MethodPost {
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
	postId, err := strconv.Atoi(r.URL.Query().Get("id"))
	if postId == 0 || err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	reaction := r.FormValue("reaction")
	if reaction == "like" {
		if err := h.Service.Reaction.CreateOrUpdateLikePost(models.Reaction{
			UserId: user.ID,
			PostId: postId,
			Islike: 1,
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	} else if reaction == "dislike" {
		if err := h.Service.Reaction.CreateOrUpdateDislikePost(models.Reaction{
			UserId: user.ID,
			PostId: postId,
			Islike: -1,
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	} else {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	link := fmt.Sprintf("/post/?id=%d", postId)
	http.Redirect(w, r, link, http.StatusSeeOther)

}
func (h *Handler) reactionComment(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/reaction/comment/" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if r.Method != http.MethodPost {
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
	commentId, err := strconv.Atoi(r.URL.Query().Get("id"))
	postId, err := strconv.Atoi(r.URL.Query().Get("postId"))
	if err != nil || commentId == 0 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	comment, err := h.Service.Comment.GetOneCommentByIdComment(commentId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	reaction := r.FormValue("reactionComment")
	if reaction == "like" {
		if err := h.Service.Reaction.CreateOrUpdateLikeComment(models.Reaction{
			UserId:    user.ID,
			CommentId: comment.IdComment,
			Islike:    1,
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	} else if reaction == "dislike" {
		if err := h.Service.Reaction.CreateOrUpdateDislikeComment(models.Reaction{
			UserId:    user.ID,
			CommentId: commentId,
			Islike:    -1,
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	} else {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	link := fmt.Sprintf("/post/?id=%d", postId)
	http.Redirect(w, r, link, http.StatusSeeOther)

}
