package routers

import (
	"net/http"

	"github.com/federus1105/socialmedia/internals/middleware"
	"github.com/federus1105/socialmedia/internals/models"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"

	docs "github.com/federus1105/socialmedia/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter(db *pgxpool.Pool, rdb *redis.Client) *gin.Engine {
	router := gin.Default()
	router.Use(gin.Recovery())
	router.Use(middleware.MyLogger)

	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Static("/img", "public")

	InitAuthRouter(router, db)
	InitPostRouter(router, db)
	InitCommentRouter(router, db)
	InitFollowRouter(router, db)
	InitPostListRouter(router, db, rdb)
	InitLikeRouter(router, db)

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, models.Response{
			Message: "Rute Salah",
			Status:  "Rute Tidak Ditemukan",
		})
	})
	return router

}
