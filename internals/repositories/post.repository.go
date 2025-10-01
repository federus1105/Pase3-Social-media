package repositories

import (
	"context"
	"log"

	"github.com/federus1105/socialmedia/internals/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostRepository struct {
	db *pgxpool.Pool
}

func NewPostRepository(db *pgxpool.Pool) *PostRepository {
	return &PostRepository{db: db}
}

func (pr *PostRepository) CreatePost(rctx context.Context, body models.PostBody) (models.PostBody, error) {
	sql := `INSERT INTO postingan (user_id, image, content) 
        VALUES ($1, $2, $3)
        RETURNING id, user_id, image, content`
	values := []any{body.User, body.ImageStr, body.Caption}
	var newPost models.PostBody
	if err := pr.db.QueryRow(rctx, sql, values...).Scan(&newPost.Id, &newPost.User, &newPost.ImageStr, &newPost.Caption); err != nil {
		log.Println("Failed to insert Postingan:", err)
		return models.PostBody{}, err
	}
	return newPost, nil
}