package repositories

import (
	"context"
	"log"

	"github.com/federus1105/socialmedia/internals/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type FollowRepository struct {
	db *pgxpool.Pool
}

func NewFollowRepository(db *pgxpool.Pool) *FollowRepository {
	return &FollowRepository{db: db}
}

func (r *FollowRepository) Follow(ctx context.Context, followerID, followingID int) (models.Follow, error) {
	sql := `
        INSERT INTO follows (follower_id, following_id)
        VALUES ($1, $2)
        RETURNING id, follower_id, following_id, created_at
    `

	var follow models.Follow
	err := r.db.QueryRow(ctx, sql, followerID, followingID).
		Scan(&follow.ID, &follow.FollowerID, &follow.FollowingID, &follow.CreatedAt)
	if err != nil {
		log.Println("Failed to insert Follow:", err)
		return models.Follow{}, err
	}

	return follow, nil
}
