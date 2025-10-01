package routers

import (
	"github.com/federus1105/socialmedia/internals/handlers"
	"github.com/federus1105/socialmedia/internals/middleware"
	"github.com/federus1105/socialmedia/internals/repositories"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitPostRouter(router *gin.Engine, db *pgxpool.Pool) {
	postingan := router.Group("/postingan")
	pr := repositories.NewPostRepository(db)
	ph := handlers.NewPostHandler(pr)

	postingan.POST("", middleware.VerifyToken, middleware.AuthMiddleware(), ph.CreatePost)
}
