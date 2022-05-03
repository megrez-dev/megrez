package admin

import (
	"archive/zip"
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/megrez/pkg/config"
	"github.com/megrez/pkg/entity/vo"
	"github.com/megrez/pkg/log"
	"github.com/megrez/pkg/model"
	dirUtils "github.com/megrez/pkg/utils/dir"
	"github.com/megrez/pkg/utils/errmsg"
	zipUtils "github.com/megrez/pkg/utils/zip"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
	"gorm.io/gorm"
)

func GetCurrentThemeConfig(c *gin.Context) {
	var cfg = &config.ThemeConfig{}
	home, err := dirUtils.GetOrCreateMegrezHome()
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
	for i, tab := range cfg.Tabs {
		for j, item := range tab.Items {
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
				cfg.Tabs[i].Items[j].Value = strings.Split(value, ";")
			} else {
				cfg.Tabs[i].Items[j].Value = value
			}
		}
	}
	c.JSON(http.StatusOK, errmsg.Success(cfg))
}

func UpdateCurrentThemeConfig(c *gin.Context) {
	var cfg config.ThemeConfig
	err := c.ShouldBindJSON(&cfg)
	if err != nil {
		log.Error("bind json failed:", err.Error())
		c.JSON(http.StatusOK, errmsg.Fail(errmsg.ErrorInvalidParam))
		return
	}
	currentThemeID, err := model.GetOptionByKey(vo.OptionKeyBlogTheme)
	if err != nil {
		log.Error("get option blog theme failed:", err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	for _, tab := range cfg.Tabs {
		for _, item := range tab.Items {
			var strValue string
			if values, ok := item.Value.([]interface{}); ok {
				if len(values) == 0 {
					strValue = item.Default
				} else {
					var strValues []string
					for _, value := range values {
						strValues = append(strValues, value.(string))
					}
					strValue = strings.Join(strValues, ";")
				}
			} else if value, ok := item.Value.(string); ok {
				if value == "" {
					strValue = item.Default
				} else {
					strValue = value
				}
			}
			err := model.UpdateThemeOption(currentThemeID, item.Key, strValue)
			if err != nil {
				log.Errorf("update theme option %s failed: %s", item.Key, err.Error())
				c.JSON(http.StatusOK, errmsg.Error())
				return
			}
		}
	}
	c.JSON(http.StatusOK, errmsg.Success(nil))
}

func InstallTheme(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		log.Error("get file from request failed: ", err)
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	fileName := file.Filename
	ext := path.Ext(fileName)
	if ext != ".zip" {
		c.JSON(http.StatusOK, errmsg.FailMsg("仅支持zip格式"))
		return
	}
	open, err := file.Open()
	if err != nil {
		log.Error("open zip file failed: ", err)
		c.JSON(http.StatusOK, errmsg.FailMsg("打开压缩文件失败"))
		return
	}
	defer open.Close()
	reader, err := zip.NewReader(open, file.Size)
	if err != nil {
		log.Error("new zip reader failed: ", err)
		c.JSON(http.StatusOK, errmsg.FailMsg("解压失败"))
		return
	}
	infoFile, err := reader.Open("theme.yaml")
	if err != nil {
		log.Error("open theme.yaml failed: ", err)
		c.JSON(http.StatusOK, errmsg.FailMsg("打开主题信息文件失败"))
		return
	}
	defer infoFile.Close()
	themeInfo, ok := getThemeInfo(infoFile)
	if !ok {
		c.JSON(http.StatusOK, errmsg.FailMsg("主题信息文件格式错误"))
		return
	}
	if themeInfo.ID == "" {
		c.JSON(http.StatusOK, errmsg.FailMsg("主题ID不能为空"))
		return
	}
	cfgFile, err := reader.Open("config.yaml")
	if err != nil {
		log.Error("open config.yaml failed: ", err)
		c.JSON(http.StatusOK, errmsg.FailMsg("打开主题配置文件失败"))
		return
	}
	defer cfgFile.Close()
	themeConfig, ok := getThemeConfig(cfgFile)
	if !ok {
		c.JSON(http.StatusOK, errmsg.FailMsg("主题配置文件格式错误"))
		return
	}
	home, err := dirUtils.GetOrCreateMegrezHome()
	if err != nil {
		log.Error("get megrez home failed:", err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	if _, err := os.Stat(path.Join(home, "themes", themeInfo.ID)); err != nil {
		if err != nil {
			if !os.IsNotExist(err) {
				log.Errorf("get theme %s dir failed: %s", themeInfo.ID, err.Error())
				c.JSON(http.StatusOK, errmsg.Error())
				return
			}
		} else {
			log.Errorf("theme %s already exists", themeInfo.ID)
			c.JSON(http.StatusOK, errmsg.FailMsg("主题已存在"))
			return
		}
	}
	if err := os.MkdirAll(path.Join(home, "themes", themeInfo.ID), 0755); err != nil {
		log.Error("create theme dir failed:", err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	err = zipUtils.UnZip(reader, path.Join(home, "themes", themeInfo.ID))
	if err != nil {
		os.RemoveAll(path.Join(home, "themes", themeInfo.ID))
		log.Error("unzip theme failed:", err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	for _, tab := range themeConfig.Tabs {
		for _, item := range tab.Items {
			option := &model.ThemeOption{
				ThemeID: themeInfo.ID,
				Key:     item.Key,
				Value:   item.Default,
				Type:    item.Type,
			}
			err := model.CreateThemeOption(option)
			if err != nil {
				os.RemoveAll(path.Join(home, "themes", themeInfo.ID))
				log.Errorf("init theme option %s failed: %s", themeInfo.ID, err.Error())
				c.JSON(http.StatusOK, errmsg.Error())
				return
			}
		}
	}
	c.JSON(http.StatusOK, errmsg.Success(nil))
}

func GetCurrentThemeID(c *gin.Context) {
	themeID, err := model.GetOptionByKey(vo.OptionKeyBlogTheme)
	if err != nil {
		log.Error("get option blog theme failed:", err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	c.JSON(http.StatusOK, errmsg.Success(themeID))
}

func getThemeConfig(file fs.File) (config.ThemeConfig, bool) {
	var cfg config.ThemeConfig
	if err := yaml.NewDecoder(file).Decode(&cfg); err != nil {
		return cfg, false
	}
	return cfg, true
}

func getThemeInfo(file fs.File) (config.ThemeInfo, bool) {
	var cfg config.ThemeInfo
	if err := yaml.NewDecoder(file).Decode(&cfg); err != nil {
		return cfg, false
	}
	return cfg, true
}
