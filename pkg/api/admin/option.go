package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/megrez/pkg/log"
	"github.com/megrez/pkg/model"
	"github.com/megrez/pkg/utils/errmsg"
	"net/http"
)

// SetOption godoc
// @Summary set blog option
// @Schemes http https
// @Description set blog option
// @Accept application/json
// @Param Authorization header string false "Authorization"
// @Param key path string false "option key"
// @Success 200 {object} errmsg.Response{}
// @Router /api/admin/option/{key} [put]
func SetOption(c *gin.Context) {
	key := c.Param("key")
	type valueJson struct {
		Value string `json:"value"`
	}
	j := &valueJson{}
	err := c.ShouldBindJSON(j)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, errmsg.ErrorInvalidParam)
		return
	}
	if key == "" {
		c.JSON(http.StatusOK, errmsg.Fail(errmsg.ErrorInvalidParam))
		return
	}
	tx := model.BeginTx()
	err = model.SetOption(tx, key, j.Value)
	if err != nil {
		log.Error("set option error: %v", err)
		tx.Rollback()
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	tx.Commit()
	c.JSON(http.StatusOK, errmsg.Success(nil))
}
