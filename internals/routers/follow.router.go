package routers

import (
	"github.com/federus1105/socialmedia/internals/handlers"
	"github.com/federus1105/socialmedia/internals/middleware"
	"github.com/federus1105/socialmedia/internals/repositories"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitFollowRouter(router *gin.Engine, db *pgxpool.Pool) {
	follow := router.Group("/follow")
	fr := repositories.NewFollowRepository(db)
	fh := handlers.NewFollowHandler(fr)

	follow.POST("/:id_user", middleware.VerifyToken, middleware.AuthMiddleware(), fh.Follow)
}
	