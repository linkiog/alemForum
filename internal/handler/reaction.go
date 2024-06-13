package handler

import (
	"fmt"
	"forum/internal/models"
	"net/http"
	"strconv"
)

func (h *Handler) reactionPost(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/reaction/post/" {
		h.ErrorPage(w, http.StatusNotFound)
		return

	}
	if r.Method != http.MethodPost {
		h.ErrorPage(w, http.StatusMethodNotAllowed)
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
	postId, err := strconv.Atoi(r.URL.Query().Get("id"))
	if postId == 0 || err != nil {
		h.ErrorPage(w, http.StatusInternalServerError)
		return
	}

	reaction := r.FormValue("reaction")
	if reaction == "like" {
		if err := h.Service.Reaction.CreateOrUpdateLikePost(models.Reaction{
			UserId: user.ID,
			PostId: postId,
			Islike: 1,
		}); err != nil {
			fmt.Println(err.Error())
			h.ErrorPage(w, http.StatusInternalServerError)
			return
		}

	} else if reaction == "dislike" {
		if err := h.Service.Reaction.CreateOrUpdateDislikePost(models.Reaction{
			UserId: user.ID,
			PostId: postId,
			Islike: -1,
		}); err != nil {
			fmt.Println(err.Error())
			h.ErrorPage(w, http.StatusInternalServerError)
			return
		}

	} else {
		h.ErrorPage(w, http.StatusInternalServerError)
		return
	}
	link := fmt.Sprintf("/post/?id=%d", postId)
	http.Redirect(w, r, link, http.StatusSeeOther)

}
func (h *Handler) reactionComment(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/reaction/comment/" {
		h.ErrorPage(w, http.StatusNotFound)
		return
	}
	if r.Method != http.MethodPost {
		h.ErrorPage(w, http.StatusMethodNotAllowed)
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
	commentId, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		h.ErrorPage(w, http.StatusBadRequest)
		return
	}
	postId, err := strconv.Atoi(r.URL.Query().Get("postId"))
	if err != nil || commentId == 0 {
		h.ErrorPage(w, http.StatusNotFound)
		return
	}

	comment, err := h.Service.Comment.GetOneCommentByIdComment(commentId)
	if err != nil {
		fmt.Println(err.Error())
		h.ErrorPage(w, http.StatusInternalServerError)
		return
	}
	reaction := r.FormValue("reactionComment")
	if reaction == "like" {
		if err := h.Service.Reaction.CreateOrUpdateLikeComment(models.Reaction{
			UserId:    user.ID,
			CommentId: comment.IdComment,
			Islike:    1,
		}); err != nil {
			fmt.Println(err.Error())
			h.ErrorPage(w, http.StatusInternalServerError)
			return
		}

	} else if reaction == "dislike" {
		if err := h.Service.Reaction.CreateOrUpdateDislikeComment(models.Reaction{
			UserId:    user.ID,
			CommentId: commentId,
			Islike:    -1,
		}); err != nil {
			fmt.Println(err.Error())
			h.ErrorPage(w, http.StatusInternalServerError)
			return
		}

	} else {
		h.ErrorPage(w, http.StatusInternalServerError)
		return
	}
	link := fmt.Sprintf("/post/?id=%d", postId)
	http.Redirect(w, r, link, http.StatusSeeOther)

}
