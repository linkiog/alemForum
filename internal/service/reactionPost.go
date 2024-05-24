package service

import (
	"forum/internal/models"
	"forum/internal/storage"
)

type Reaction interface {
	CreateOrUpdateLikePost(postReaction models.Reaction) error
	CreateOrUpdateDislikePost(postReaction models.Reaction) error

	CreateOrUpdateLikeComment(postReaction models.Reaction) error
	CreateOrUpdateDislikeComment(postReaction models.Reaction) error
}

type ReactionService struct {
	storage storage.Reaction
}

func InitReactionService(strorage storage.Reaction) Reaction {
	return &ReactionService{
		storage: strorage,
	}

}
func (ld *ReactionService) CreateOrUpdateLikePost(postReaction models.Reaction) error {
	reaction, err, exists := ld.storage.ReactionPostExists(postReaction)
	if err != nil {
		return err
	}
	if !exists {
		if err := ld.storage.ReactionCreatePost(postReaction); err != nil {
			return err
		}
	} else {
		if reaction.Islike == 1 {
			if err := ld.storage.ReactionDeletePost(postReaction); err != nil {
				return err
			}
			postReaction.Islike = 0
			if err := ld.storage.UpdateReactionPost(postReaction); err != nil {
				return err
			}

		} else if reaction.Islike == -1 {
			reaction.Islike = 1
			if err := ld.storage.UpdateReactionPost(postReaction); err != nil {
				return err
			}

		}

	}
	if err := ld.storage.UpdateAllInformationPost(postReaction); err != nil {
		return err
	}

	return nil

}
func (ld *ReactionService) CreateOrUpdateDislikePost(postReaction models.Reaction) error {
	reaction, err, exists := ld.storage.ReactionPostExists(postReaction)
	if err != nil {
		return err
	}
	if !exists {
		if err := ld.storage.ReactionCreatePost(postReaction); err != nil {
			return err
		}
	} else {
		if reaction.Islike == -1 {
			if err := ld.storage.ReactionDeletePost(postReaction); err != nil {
				return err
			}
			postReaction.Islike = 0
			if err := ld.storage.UpdateReactionPost(postReaction); err != nil {
				return err
			}

		} else if reaction.Islike == 1 {
			postReaction.Islike = -1
			if err := ld.storage.UpdateReactionPost(postReaction); err != nil {
				return err
			}

		}

	}
	if err := ld.storage.UpdateAllInformationPost(postReaction); err != nil {
		return err
	}

	return nil

}
