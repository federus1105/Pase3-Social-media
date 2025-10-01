package routers

import (
	"github.com/federus1105/socialmedia/internals/handlers"
	"github.com/federus1105/socialmedia/internals/middleware"
	"github.com/federus1105/socialmedia/internals/repositories"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitLikeRouter(router *gin.Engine, db *pgxpool.Pool) {
	like := router.Group("/like")
	lr := repositories.NewLikeRepository(db)
	lh := handlers.NewLikeHandler(lr)

	like.POST("/:post_id", middleware.VerifyToken, middleware.AuthMiddleware(), lh.LikePost)
}
