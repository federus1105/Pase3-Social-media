package routers

import (
	"github.com/federus1105/socialmedia/internals/handlers"
	"github.com/federus1105/socialmedia/internals/middleware"
	"github.com/federus1105/socialmedia/internals/repositories"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitCommentRouter(router *gin.Engine, db *pgxpool.Pool) {
	comment := router.Group("/comment")
	cr := repositories.NewCommentRepository(db)
	ch := handlers.NewCommentHandler(cr)

	comment.POST("/:postingan_id", middleware.VerifyToken, middleware.AuthMiddleware(), ch.CreateComment)
}
