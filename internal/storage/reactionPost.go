package storage

import (
	"database/sql"
	"fmt"
	"forum/internal/models"
)

type Reaction interface {
	ReactionCreatePost(reaction models.Reaction) error
	ReactionPostExists(reaction models.Reaction) (models.Reaction, error, bool)
	ReactionDeletePost(reaction models.Reaction) error
	UpdateReactionPost(reaction models.Reaction) error
	UpdateAllInformationPost(reaction models.Reaction) error

	ReactionCreateComment(reaction models.Reaction) error
	ReactionCommentExists(reaction models.Reaction) (models.Reaction, error, bool)
	ReactionDeleteComment(reaction models.Reaction) error
	UpdateReactionComment(reaction models.Reaction) error
	UpdateAllInformationComment(reaction models.Reaction) error
}

type ReactionStore struct {
	db *sql.DB
}

func InitReaction(db *sql.DB) Reaction {
	return &ReactionStore{
		db: db,
	}

}

func (ld *ReactionStore) ReactionCreatePost(reaction models.Reaction) error {
	query := `INSERT INTO reaction(userId,postId,reaction) VALUES($1,$2,$3);`
	_, err := ld.db.Exec(query, reaction.UserId, reaction.PostId, reaction.Islike)
	if err != nil {
		return fmt.Errorf("Reaction models %w" + err.Error())
	}
	return nil

}

func (ld *ReactionStore) ReactionPostExists(reaction models.Reaction) (models.Reaction, error, bool) {
	query := "SELECT reaction FROM reaction WHERE postId=$1 AND userId=$2;"
	var emotion models.Reaction
	err := ld.db.QueryRow(query, reaction.PostId, reaction.UserId).Scan(&emotion.Islike)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Reaction{}, nil, false

		}
		return models.Reaction{}, fmt.Errorf("ReactionPOstExists", err), false
	}
	return emotion, nil, true

}

func (ld *ReactionStore) ReactionDeletePost(reaction models.Reaction) error {
	query := "DELETE FROM reaction WHERE userId=$1 AND postId=$2;"
	_, err := ld.db.Exec(query, reaction.UserId, reaction.PostId)
	if err != nil {
		return fmt.Errorf("storage func: Reaction Delete Post: %w", err)
	}
	return nil
}

func (ld *ReactionStore) UpdateReactionPost(reaction models.Reaction) error {
	query := `UPDATE reaction SET reaction = $1 WHERE postId = $2 AND userId = $3;`
	_, err := ld.db.Exec(query, reaction.Islike, reaction.PostId, reaction.UserId)
	if err != nil {
		return fmt.Errorf("storage func: Update Reaction Post: %w", err)
	}

	return nil
}

func (ld *ReactionStore) CountLikeAndDislike(reaction models.Reaction) (models.Reaction, error) {
	query := `SELECT
	(SELECT COUNT(*) FROM reaction WHERE postId=$1 AND reaction=1) AS likes,
	(SELECT COUNT(*) FROM reaction WHERE postId=$1 AND reaction=-1) AS dislike;
	`
	row := ld.db.QueryRow(query, reaction.PostId)
	if err := row.Scan(&reaction.CountLike, &reaction.CountDislike); err != nil {
		return reaction, fmt.Errorf("problem with CountLikeAndDislike: /%w", err)
	}

	return reaction, nil

}
func (ld *ReactionStore) UpdateAllInformationPost(reaction models.Reaction) error {
	react, err := ld.CountLikeAndDislike(reaction)
	if err != nil {
		return fmt.Errorf("problem UpdateAllInformationPost: %w", err)
	}
	query := `UPDATE post SET like =$1, dislike=$2 WHERE idPost=$3;`
	if _, err := ld.db.Exec(query, react.CountLike, react.CountDislike, reaction.PostId); err != nil {
		return fmt.Errorf("problem with UpdateAllInformationPost EXES: %w", err)
	}
	return nil

}
