package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/megrez/pkg/entity/dto"
	"github.com/megrez/pkg/utils/errmsg"
)

func GetSettings(c *gin.Context) {
	settings := &dto.Setting{}
	err := settings.LoadFromModel()
	if err != nil {
		c.JSON(200, errmsg.Error())
		return
	}
	c.JSON(200, errmsg.Success(settings))
}

func UpdateSettings(c *gin.Context) {
	settings := &dto.Setting{}
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
