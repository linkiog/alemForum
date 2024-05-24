package storage

import (
	"database/sql"
	"fmt"
	"forum/internal/models"
)

func (ld *ReactionStore) ReactionCreateComment(reaction models.Reaction) error {
	query := `INSERT INTO reactionComment(userId,commentId,reaction) VALUES($1,$2,$3);`
	_, err := ld.db.Exec(query, reaction.UserId, reaction.CommentId, reaction.Islike)
	if err != nil {
		return fmt.Errorf("Reaction Create Comment %w" + err.Error())
	}
	return nil

}

func (ld *ReactionStore) ReactionCommentExists(reaction models.Reaction) (models.Reaction, error, bool) {
	query := "SELECT reaction FROM reactionComment WHERE commentId=$1 AND userId=$2;"
	var emotion models.Reaction
	err := ld.db.QueryRow(query, reaction.CommentId, reaction.UserId).Scan(&emotion.Islike)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Reaction{}, nil, false

		}
		return models.Reaction{}, fmt.Errorf("ReactionCommenttExists", err), false
	}
	return emotion, nil, true

}

func (ld *ReactionStore) ReactionDeleteComment(reaction models.Reaction) error {
	query := "DELETE FROM reactionComment WHERE userId=$1 AND commentId=$2;"
	_, err := ld.db.Exec(query, reaction.UserId, reaction.CommentId)
	if err != nil {
		return fmt.Errorf("storage func: Reaction Delete Comment: %w", err)
	}
	return nil
}

func (ld *ReactionStore) UpdateReactionComment(reaction models.Reaction) error {
	query := `UPDATE reactionComment SET reaction = $1 WHERE commentId = $2 AND userId = $3;`
	_, err := ld.db.Exec(query, reaction.Islike, reaction.CommentId, reaction.UserId)
	if err != nil {
		return fmt.Errorf("storage func: Update Reaction Comment: %w", err)
	}

	return nil
}

func (ld *ReactionStore) CountLikeAndDislikeComment(reaction models.Reaction) (models.Reaction, error) {
	query := `SELECT
	(SELECT COUNT(*) FROM reactionComment WHERE commentId=$1 AND reaction=1) AS likes,
	(SELECT COUNT(*) FROM reactionComment WHERE commentId=$1 AND reaction=-1) AS dislike;
	`
	row := ld.db.QueryRow(query, reaction.CommentId)
	if err := row.Scan(&reaction.CountLike, &reaction.CountDislike); err != nil {
		return reaction, fmt.Errorf("problem with CountLikeAndDislikeComment: /%w", err)
	}

	return reaction, nil

}
func (ld *ReactionStore) UpdateAllInformationComment(reaction models.Reaction) error {
	react, err := ld.CountLikeAndDislikeComment(reaction)
	if err != nil {
		return fmt.Errorf("problem UpdateAllInformationPost: %w", err)
	}
	query := `UPDATE comment SET like =$1, dislike=$2 WHERE idComment=$3;`
	if _, err := ld.db.Exec(query, react.CountLike, react.CountDislike, reaction.CommentId); err != nil {
		return fmt.Errorf("problem with UpdateAllInformationComment EXES: %w", err)
	}
	return nil

}
