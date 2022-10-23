package api

import (
	"github.com/gin-gonic/gin"
	"github.com/megrez/pkg/api/openapi"
)

func RouteOpenAPI(g *gin.Engine) {
	api := g.Group("api")

	// api for comment
	// api.POST("comment", openapi.CreateComment)
	// auth.PUT("comment/:id", admin.UpdateComment)
	// api.DELETE("comment/:id", openapi.DeleteComment)
	api.GET(":type/:id/comments", openapi.ListComments)
	// auth.GET("comment/:id", admin.GetComment)
}
