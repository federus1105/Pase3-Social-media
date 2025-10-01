package routers

import (
	"github.com/federus1105/socialmedia/internals/handlers"
	"github.com/federus1105/socialmedia/internals/middleware"
	"github.com/federus1105/socialmedia/internals/repositories"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func InitPostListRouter(router *gin.Engine, db *pgxpool.Pool, rdb *redis.Client) {
	postingan := router.Group("/listpostingan")
	pr := repositories.NewPostListRepository(db, rdb)
	ph := handlers.NewPostListHandler(pr)

	postingan.GET("/:following_id", middleware.VerifyToken, middleware.AuthMiddleware(), ph.GetUserPosts)
}
