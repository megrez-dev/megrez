package checkinstall

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/megrez/pkg/model"
	"github.com/megrez/pkg/utils/errmsg"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strings"
)

func CheckInstall() gin.HandlerFunc {
	return func(c *gin.Context) {
		// exclude api for install
		if c.Request.URL.Path == "/api/admin/install" {
			c.Next()
			return
		}
		// exclude page for install
		if c.Request.URL.Path == "/admin/install" {
			c.Next()
			return
		}
		// exclude statics for admin
		if strings.HasPrefix(c.Request.URL.Path, "/admin/css") || strings.HasPrefix(c.Request.URL.Path, "/admin/js") {
			c.Next()
			return
		}
		isInstalledStr, err := model.GetOptionByKey(model.OptionKeyIsInstalled)
		if err == gorm.ErrRecordNotFound {
			log.Println("redirect to install page, origin path:", c.Request.URL.Path)
			// TODO: 判断完之后，后续 err 判空处理可能有 bug，所有 ErrRecordNotFound 都会有这个问题。
			if strings.HasPrefix(c.Request.URL.Path, "/api/admin") {
				fmt.Println("need redirect to install page")
				c.AbortWithStatusJSON(http.StatusOK, errmsg.Fail(errmsg.ErrorNotInstalled))
				return
			} else {
				c.Redirect(http.StatusFound, "/admin/install")
				return
			}
		} else if err != nil {
			if strings.HasPrefix(c.Request.URL.Path, "/api/admin") {
				c.AbortWithStatusJSON(http.StatusOK, errmsg.Error())
				return
			} else {
				c.Redirect(http.StatusInternalServerError, "/error")
				return
			}
		}
		if isInstalledStr == "true" {
			c.Next()
		} else {
			if strings.HasPrefix(c.Request.URL.Path, "/api/admin") {
				c.AbortWithStatusJSON(http.StatusOK, errmsg.Error())
				return
			} else {
				c.Redirect(http.StatusInternalServerError, "/error")
				return
			}
		}
	}
}
