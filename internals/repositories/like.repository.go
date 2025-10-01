package repositories

import (
	"context"
	"log"

	"github.com/federus1105/socialmedia/internals/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type LikeRepository struct {
	db *pgxpool.Pool
}

func NewLikeRepository(db *pgxpool.Pool) *LikeRepository {
	return &LikeRepository{db: db}
}

// Tambah Like
func (lr *LikeRepository) CreateLike(ctx context.Context, body models.Like) (models.Like, error) {
	sql := `INSERT INTO likes (user_id, post_id)
            VALUES ($1, $2)
            RETURNING id, user_id, post_id, created_at`

	values := []any{body.UserID, body.PostID}

	var newLike models.Like
	if err := lr.db.QueryRow(ctx, sql, values...).
		Scan(&newLike.ID, &newLike.UserID, &newLike.PostID, &newLike.CreatedAt); err != nil {
		log.Println("Failed to insert Like:", err)
		return models.Like{}, err
	}

	return newLike, nil
}
