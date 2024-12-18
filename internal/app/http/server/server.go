package server

import (
	"test/api/internal/app/http/healthcheck"

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

	router.Run(":8080")
}
