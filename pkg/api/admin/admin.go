package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/megrez/pkg/entity/dto"
	"github.com/megrez/pkg/utils/errmsg"
	"log"
	"net/http"
)

func Login(c *gin.Context) {
	var data dto.LoginForm
	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Println("decode json data failed, ", err.Error())
		c.JSON(http.StatusOK, errmsg.Fail(errmsg.ErrorInvalidParam))
		return
	}
	c.JSON(http.StatusOK, errmsg.Success(nil))
}
