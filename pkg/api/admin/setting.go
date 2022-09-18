package admin

import (
	"github.com/gin-gonic/gin"
	admindto "github.com/megrez/pkg/entity/dto/admin"
	"github.com/megrez/pkg/utils/errmsg"
)

func GetSettings(c *gin.Context) {
	settings := &admindto.Setting{}
	err := settings.LoadFromModel()
	if err != nil {
		c.JSON(200, errmsg.Error())
		return
	}
	c.JSON(200, errmsg.Success(settings))
}

func UpdateSettings(c *gin.Context) {
	settings := &admindto.Setting{}
	err := c.ShouldBindJSON(settings)
	if err != nil {
		c.JSON(200, errmsg.Fail(errmsg.ErrorInvalidParam))
		return
	}
	err = settings.SaveToModel()
	if err != nil {
		c.JSON(200, errmsg.Error())
		return
	}
	c.JSON(200, errmsg.Success(nil))
}
