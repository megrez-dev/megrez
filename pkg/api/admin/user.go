package admin

import (
	"github.com/gin-gonic/gin"
	dto "github.com/megrez/pkg/entity/dto/admin"
	"github.com/megrez/pkg/model"
	"github.com/megrez/pkg/utils/errmsg"
	"github.com/megrez/pkg/utils/jwt"
	"gorm.io/gorm"
	"log"
	"net/http"
)

// Login godoc
// @Summary login form
// @Schemes http https
// @Description login form
// @Accept application/json
// @Param  req body dto.LoginForm true "login form"
// @Success 200 {object} errmsg.Response{data=string}
// @Router /api/admin/login [post]
func Login(c *gin.Context) {
	var data dto.LoginForm
	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Println("decode json data failed, ", err.Error())
		c.JSON(http.StatusOK, errmsg.Fail(errmsg.ErrorInvalidParam))
		return
	}
	user, err := model.GetUserByUsername(data.Username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, errmsg.Fail(errmsg.ErrorPasswordWrong))
			return
		}
		log.Println("get user failed, ", err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	if user.Password != data.Password {
		c.JSON(http.StatusOK, errmsg.Fail(errmsg.ErrorPasswordWrong))
		return
	}
	token, err := jwt.GenerateToken(user.ID)
	if err != nil {
		log.Println("generate token failed, ", err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	c.JSON(http.StatusOK, errmsg.Success(token))
}
