package jwt

import (
	"github.com/gin-gonic/gin"
	"github.com/megrez/pkg/utils/errmsg"
	"github.com/megrez/pkg/utils/jwt"
	"net/http"
	"time"
)

func Jwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		// exclude api for install and login
		if c.Request.URL.Path == "/api/admin/install" || c.Request.URL.Path == "/api/admin/login" {
			c.Next()
		} else {
			token := c.Request.Header.Get("Megrez-Token")
			if token == "" {
				c.AbortWithStatusJSON(http.StatusOK, errmsg.Fail(errmsg.ErrorTokenNotExist))
				return
			}
			claims, err := jwt.ParseToken(token)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusOK, errmsg.Fail(errmsg.ErrorTokenInvalid))
				return
			}
			if claims.VerifyExpiresAt(time.Now(), true) == false {
				c.AbortWithStatusJSON(http.StatusOK, errmsg.Fail(errmsg.ErrorTokenInvalid))
				return
			}
			c.Set("uid", claims.ID)
			c.Next()
		}
	}
}
