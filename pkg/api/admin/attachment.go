package admin

import "github.com/gin-gonic/gin"

func UploadAttachment(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
