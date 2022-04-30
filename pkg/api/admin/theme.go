package admin

import (
	"io/ioutil"
	"net/http"
	"path"
	"strings"

	"github.com/megrez/pkg/config"
	"github.com/megrez/pkg/entity/vo"
	"github.com/megrez/pkg/log"
	"github.com/megrez/pkg/model"
	"github.com/megrez/pkg/utils/dir"
	"github.com/megrez/pkg/utils/errmsg"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
)

func GetThemeConfig(c *gin.Context) {
	var cfg = &config.ThemeConfig{}
	home, err := dir.GetOrCreateMegrezHome()
	if err != nil {
		log.Error("get megrez home failed:", err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	themeID, err := model.GetOptionByKey(vo.OptionKeyBlogTheme)
	if err != nil {
		log.Error("get option theme failed:", err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	b, err := ioutil.ReadFile(path.Join(home, "themes", themeID, "config.yaml"))
	if err != nil {
		log.Error("read theme config file failed, ", err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	if err := yaml.Unmarshal(b, cfg); err != nil {
		log.Error("unmarshal config failed, ", err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	for _, tab := range cfg.Tabs {
		for _, item := range tab.Items {
			value, err := model.GetThemeOptionByThemeIDAndKey(themeID, item.Key)
			if err != nil {
				if err == gorm.ErrRecordNotFound {
					value = item.Default
				} else {
					log.Error("get theme option failed:", err.Error())
					c.JSON(http.StatusOK, errmsg.Error())
					return
				}
			}
			if item.Type == "multiSelect" || item.Type == "tags" {
				item.Value = strings.Split(value, ",")
			} else {
				item.Value = value
			}
		}
	}
	c.JSON(http.StatusOK, errmsg.Success(cfg))
}
