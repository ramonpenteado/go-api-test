package server

import (
	"test/api/internal/app/http/healthcheck"

	postsRoutes "test/api/internal/app/routes/posts"
	userRoutes "test/api/internal/app/routes/user"

	"github.com/gin-gonic/gin"
)

func Init() {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/healthcheck", healthcheck.HealthcheckHandler)

	userGroup := router.Group("/user")
	userRoutes.UserRoutes(userGroup)

	postsGroup := router.Group("/posts")
	postsRoutes.PostsRoutes(postsGroup)

	router.Run(":8080")
}
