package healthcheck

import "github.com/gin-gonic/gin"

func HealthcheckHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "OK"})
}
