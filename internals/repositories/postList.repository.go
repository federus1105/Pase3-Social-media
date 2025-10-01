package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/federus1105/socialmedia/internals/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type PostListRepository struct {
	db  *pgxpool.Pool
	rdb *redis.Client
}

func NewPostListRepository(db *pgxpool.Pool, rdb *redis.Client) *PostListRepository {
	return &PostListRepository{db: db, rdb: rdb}
}

func (pr *PostListRepository) GetUserPostsIfFollow(ctx context.Context, followerID, followingID int) ([]models.Post, error) {
	start := time.Now()
	redisKey := "firdaus:Get post user following"
	cmd := pr.rdb.Get(ctx, redisKey)
	if cmd.Err() != nil {
		if cmd.Err() == redis.Nil {
			log.Printf("Key %s does not exist\n", redisKey)
		} else {
			log.Println("Redis Error. \nCause: ", cmd.Err().Error())
		}
	} else {
		// cache hit
		var cachedSchedules []models.Post
		cmdByte, err := cmd.Bytes()
		if err != nil {
			log.Println("Internal server error.\nCause: ", err.Error())
		} else {
			if err := json.Unmarshal(cmdByte, &cachedSchedules); err != nil {
				log.Println("Internal Server Error. \nCause: ", err.Error())
			}
		}
		if len(cachedSchedules) > 0 {
			log.Printf("Key %s found in cache ✅", redisKey)
			log.Printf("Served in %s using Redis", time.Since(start))
			return cachedSchedules, nil
		}
	}
	// Cek apakah follower sudah follow following
	checkFollowSQL := `SELECT 1 FROM follows WHERE follower_id=$1 AND following_id=$2`
	var exists int
	err := pr.db.QueryRow(ctx, checkFollowSQL, followerID, followingID).Scan(&exists)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("anda belum mengikuti user ini")
		}
		return nil, err
	}

	// Ambil postingan user yang diikuti
	sql := `SELECT id, user_id, image, content, created_at
			FROM postingan
			WHERE user_id=$1
			ORDER BY created_at DESC`

	rows, err := pr.db.Query(ctx, sql, followingID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.Id, &post.User, &post.Image, &post.Caption, &post.CreatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	// renew cache
		bt, err := json.Marshal(posts)
		if err != nil {
			log.Println("Internal Server Error.\n Cause: ", err.Error())
		}
		if err := pr.rdb.Set(ctx, redisKey, string(bt), 1*time.Minute).Err(); err != nil {
			log.Println("Redis Error. \nCause: ", err.Error())
		}
	log.Printf("[REDIS TIMING] Served in %s using DB (cache miss)", time.Since(start))

	return posts, nil
}
