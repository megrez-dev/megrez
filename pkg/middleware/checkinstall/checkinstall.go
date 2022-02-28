package checkinstall

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/megrez/pkg/entity/vo"
	"github.com/megrez/pkg/model"
	"github.com/megrez/pkg/utils/errmsg"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

func CheckInstall() gin.HandlerFunc {
	return func(c *gin.Context) {
		// exclude api for install
		if c.Request.URL.Path == "/api/admin/install" {
			c.Next()
		}
		isInstalledStr, err := model.GetOptionByKey(vo.OptionKeyIsInstalled)
		if err == gorm.ErrRecordNotFound {
			if strings.HasPrefix(c.Request.URL.Path, "/api") {
				fmt.Println("need redirect to install page")
				c.AbortWithStatusJSON(http.StatusOK, errmsg.Fail(errmsg.ErrorNotInstalled))
				return
			} else {
				c.Redirect(http.StatusFound, "/admin/install")
				return
			}
		} else if err != nil {
			if strings.HasPrefix(c.Request.URL.Path, "/api") {
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
			if strings.HasPrefix(c.Request.URL.Path, "/api") {
				c.AbortWithStatusJSON(http.StatusOK, errmsg.Error())
				return
			} else {
				c.Redirect(http.StatusInternalServerError, "/error")
				return
			}
		}
	}
}
