package repositories

import (
	"context"

	"github.com/federus1105/socialmedia/internals/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CommentRepository struct {
	db *pgxpool.Pool
}

func NewCommentRepository(db *pgxpool.Pool) *CommentRepository {
	return &CommentRepository{db: db}
}

func (cr *CommentRepository) CreateComment(ctx context.Context, comment models.Comment) (models.Comment, error) {
	sql := `
        INSERT INTO comment (postingan_id, user_id, teks)
        VALUES ($1, $2, $3)
        RETURNING id, postingan_id, user_id, teks, created_at
    `
	var newComment models.Comment
	err := cr.db.QueryRow(ctx, sql, comment.PostinganId, comment.UserID, comment.Teks).
		Scan(&newComment.Id, &newComment.PostinganId, &newComment.UserID, &newComment.Teks, &newComment.CreatedAt)
	if err != nil {
		return models.Comment{}, err
	}
	return newComment, nil
}