package service

import "forum/internal/models"

func (ld *ReactionService) CreateOrUpdateLikeComment(postReaction models.Reaction) error {
	reaction, err, exists := ld.storage.ReactionCommentExists(postReaction)
	if err != nil {
		return err
	}
	if !exists {
		if err := ld.storage.ReactionCreateComment(postReaction); err != nil {
			return err
		}
	} else {
		if reaction.Islike == 1 {
			if err := ld.storage.ReactionDeleteComment(postReaction); err != nil {
				return err
			}
			postReaction.Islike = 0
			if err := ld.storage.UpdateReactionComment(postReaction); err != nil {
				return err
			}

		} else if reaction.Islike == -1 {
			reaction.Islike = 1
			if err := ld.storage.UpdateReactionComment(postReaction); err != nil {
				return err
			}

		}

	}
	if err := ld.storage.UpdateAllInformationComment(postReaction); err != nil {
		return err
	}

	return nil

}
func (ld *ReactionService) CreateOrUpdateDislikeComment(postReaction models.Reaction) error {
	reaction, err, exists := ld.storage.ReactionCommentExists(postReaction)
	if err != nil {
		return err
	}
	if !exists {
		if err := ld.storage.ReactionCreateComment(postReaction); err != nil {
			return err
		}
	} else {
		if reaction.Islike == -1 {
			if err := ld.storage.ReactionDeleteComment(postReaction); err != nil {
				return err
			}
			postReaction.Islike = 0
			if err := ld.storage.UpdateReactionComment(postReaction); err != nil {
				return err
			}

		} else if reaction.Islike == 1 {
			postReaction.Islike = -1
			if err := ld.storage.UpdateReactionComment(postReaction); err != nil {
				return err
			}

		}

	}
	if err := ld.storage.UpdateAllInformationComment(postReaction); err != nil {
		return err
	}

	return nil

}
