package admin

import (
	"archive/zip"
	"io/fs"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/megrez/pkg/config"
	"github.com/megrez/pkg/log"
	"github.com/megrez/pkg/model"
	dirUtils "github.com/megrez/pkg/utils/dir"
	"github.com/megrez/pkg/utils/errmsg"
	zipUtils "github.com/megrez/pkg/utils/zip"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

// UpdateCurrentThemeConfig godoc
// @Summary update current theme config
// @Schemes http https
// @Description update current theme config
// @Accept application/json
// @Param Authorization header string false "Authorization"
// @Param req body config.ThemeConfig true "body"
// @Success 200 {object} errmsg.Response{}
// @Router /api/admin/theme/current/config [put]
func UpdateCurrentThemeConfig(c *gin.Context) {
	var cfg config.ThemeConfig
	err := c.ShouldBindJSON(&cfg)
	if err != nil {
		log.Error("bind json failed:", err.Error())
		c.JSON(http.StatusOK, errmsg.Fail(errmsg.ErrorInvalidParam))
		return
	}
	currentThemeID, err := model.GetOptionByKey(model.OptionKeyBlogTheme)
	if err != nil {
		log.Error("get option blog theme failed:", err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	tx := model.BeginTx()
	for _, tab := range cfg.Tabs {
		for _, item := range tab.Items {
			var strValue string
			if values, ok := item.Value.([]interface{}); ok {
				var strValues []string
				for _, value := range values {
					strValues = append(strValues, value.(string))
				}
				strValue = strings.Join(strValues, ";")
			} else if value, ok := item.Value.(string); ok {
				strValue = value
			}
			err := model.UpdateThemeOption(tx, currentThemeID, item.Key, strValue)
			if err != nil {
				log.Errorf("update theme option %s failed: %s", item.Key, err.Error())
				tx.Rollback()
				c.JSON(http.StatusOK, errmsg.Error())
				return
			}
		}
	}
	tx.Commit()
	c.JSON(http.StatusOK, errmsg.Success(nil))
}

// InstallTheme godoc
// @Summary install theme via upload zip file
// @Schemes http https
// @Description install theme via upload zip file
// @Accept multipart/form-data
// @Param Authorization header string false "Authorization"
// @Param  file formData file true "file"
// @Success 200 {object} errmsg.Response{}
// @Router /api/admin/theme/install [post]
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
	defer func(open multipart.File) {
		err := open.Close()
		if err != nil {
			log.Error("close zip file failed: ", err)
		}
	}(open)
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
	defer func(infoFile fs.File) {
		err := infoFile.Close()
		if err != nil {
			log.Error("close theme.yaml failed: ", err)
		}
	}(infoFile)
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
	defer func(cfgFile fs.File) {
		err := cfgFile.Close()
		if err != nil {
			log.Error("close config.yaml failed: ", err)
		}
	}(cfgFile)
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
	if _, err := os.Stat(path.Join(home, "themes", themeInfo.ID)); err == nil {
		log.Errorf("theme %s already exists", themeInfo.ID)
		c.JSON(http.StatusOK, errmsg.FailMsg("主题文件夹已存在"))
		return
	} else if !os.IsNotExist(err) {
		log.Errorf("get theme %s dir failed: %s", themeInfo.ID, err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	tx := model.BeginTx()
	for _, tab := range themeConfig.Tabs {
		for _, item := range tab.Items {
			option := &model.ThemeOption{
				ThemeID: themeInfo.ID,
				Key:     item.Key,
				Value:   item.Default,
				Type:    item.Type,
			}
			if err := model.CreateThemeOption(tx, option); err != nil {
				log.Errorf("init theme option %s failed: %s", themeInfo.ID, err.Error())
				tx.Rollback()
				c.JSON(http.StatusOK, errmsg.Error())
				return
			}
		}
	}
	if err := os.MkdirAll(path.Join(home, "themes", themeInfo.ID), 0755); err != nil {
		log.Error("create theme dir failed:", err.Error())
		tx.Rollback()
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	if err := zipUtils.UnZip(reader, path.Join(home, "themes", themeInfo.ID)); err != nil {
		if e := os.RemoveAll(path.Join(home, "themes", themeInfo.ID)); e != nil {
			log.Error("remove theme dir failed:", e.Error())
		}
		log.Error("unzip theme failed:", err.Error())
		tx.Rollback()
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	tx.Commit()
	c.JSON(http.StatusOK, errmsg.Success(nil))
}

func UpdateTheme(c *gin.Context) {
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
	defer func(open multipart.File) {
		err := open.Close()
		if err != nil {
			log.Error("close zip file failed: ", err)
		}
	}(open)
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
	defer func(infoFile fs.File) {
		err := infoFile.Close()
		if err != nil {
			log.Error("close theme.yaml failed: ", err)
		}
	}(infoFile)
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
	defer func(cfgFile fs.File) {
		err := cfgFile.Close()
		if err != nil {
			log.Error("close config.yaml failed: ", err)
		}
	}(cfgFile)
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
		if os.IsNotExist(err) {
			log.Errorf("theme %s not exists", themeInfo.ID)
			c.JSON(http.StatusOK, errmsg.FailMsg("主题文件夹不存在"))
			return
		} else {
			log.Errorf("get theme %s dir failed: %s", themeInfo.ID, err.Error())
			c.JSON(http.StatusOK, errmsg.Error())
			return
		}
	}
	tx := model.BeginTx()
	for _, tab := range themeConfig.Tabs {
		for _, item := range tab.Items {
			option := &model.ThemeOption{
				ThemeID: themeInfo.ID,
				Key:     item.Key,
				Value:   item.Default,
				Type:    item.Type,
			}
			if err := model.CreateThemeOptionIfNotExists(tx, option); err != nil {
				log.Errorf("init theme option %s failed: %s", themeInfo.ID, err.Error())
				tx.Rollback()
				c.JSON(http.StatusOK, errmsg.Error())
				return
			}
		}
	}
	// TODO: 怎么样在不删除原来的文件夹的情况下更新主题
	if err := os.MkdirAll(path.Join(home, "themes", themeInfo.ID), 0755); err != nil {
		log.Error("create theme dir failed:", err.Error())
		tx.Rollback()
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	if err := zipUtils.UnZip(reader, path.Join(home, "themes", themeInfo.ID)); err != nil {
		if e := os.RemoveAll(path.Join(home, "themes", themeInfo.ID)); e != nil {
			log.Error("remove theme dir failed:", e.Error())
		}
		log.Error("unzip theme failed:", err.Error())
		tx.Rollback()
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	tx.Commit()
	c.JSON(http.StatusOK, errmsg.Success(nil))
}

// DeleteTheme godoc
// @Summary delete theme by theme id
// @Schemes http https
// @Description delete theme by theme id
// @Accept application/json
// @Param Authorization header string false "Authorization"
// @Param  id path string true "theme id"
// @Success 200 {object} errmsg.Response{}
// @Router /api/admin/theme/{id} [delete]
func DeleteTheme(c *gin.Context) {
	home, err := dirUtils.GetOrCreateMegrezHome()
	if err != nil {
		log.Error("get megrez home dir failed")
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	themeID := c.Param("id")
	if themeID == "" {
		c.JSON(http.StatusOK, errmsg.Fail(errmsg.ErrorInvalidParam))
		return
	}
	err = os.RemoveAll(path.Join(home, "themes", themeID))
	if err != nil {
		log.Errorf("delete theme %s dir failed: %s", c.Param("id"), err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	tx := model.BeginTx()
	err = model.DeleteThemeOptionsByID(tx, themeID)
	if err != nil {
		log.Error("delete theme options failed: ", err.Error())
		tx.Rollback()
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	tx.Commit()
	c.JSON(http.StatusOK, errmsg.Success(nil))
}

// GetCurrentThemeConfig godoc
// @Summary get current theme config
// @Schemes http https
// @Description get current theme config
// @Accept application/json
// @Param Authorization header string false "Authorization"
// @Success 200 {object} errmsg.Response{data=config.ThemeConfig}
// @Router /api/admin/theme/current/config [get]
func GetCurrentThemeConfig(c *gin.Context) {
	var cfg = &config.ThemeConfig{}
	home, err := dirUtils.GetOrCreateMegrezHome()
	if err != nil {
		log.Error("get megrez home failed:", err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	themeID, err := model.GetOptionByKey(model.OptionKeyBlogTheme)
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
				log.Error("get theme option failed:", err.Error())
				c.JSON(http.StatusOK, errmsg.Error())
				return
			}
			if item.Type == config.ItemTypeMultiSelect || item.Type == config.ItemTypeTags {
				if value == "" {
					cfg.Tabs[i].Items[j].Value = []string{}
				} else {
					cfg.Tabs[i].Items[j].Value = strings.Split(value, ";")
				}
			} else {
				cfg.Tabs[i].Items[j].Value = value
			}
		}
	}
	c.JSON(http.StatusOK, errmsg.Success(cfg))
}

// ListThemes godoc
// @Summary list themes
// @Schemes http https
// @Description list themes
// @Accept application/json
// @Param Authorization header string false "Authorization"
// @Success 200 {object} errmsg.Response{data=[]config.ThemeInfo}
// @Router /api/admin/themes [get]
func ListThemes(c *gin.Context) {
	home, err := dirUtils.GetOrCreateMegrezHome()
	if err != nil {
		log.Error("get megrez home failed:", err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	themes, err := ioutil.ReadDir(path.Join(home, "themes"))
	if err != nil {
		log.Error("read themes dir failed:", err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	var themeList []config.ThemeInfo
	for _, theme := range themes {
		if !theme.IsDir() {
			continue
		}
		infoFile, err := os.Open(path.Join(home, "themes", theme.Name(), "theme.yaml"))
		if err != nil {
			log.Errorf("open theme dir %s info file failed: %s", theme.Name(), err.Error())
			c.JSON(http.StatusOK, errmsg.Error())
			return
		}
		defer func(infoFile *os.File) {
			err := infoFile.Close()
			if err != nil {
				log.Error("close theme info file failed:", err.Error())
			}
		}(infoFile)
		themeInfo, ok := getThemeInfo(infoFile)
		if !ok {
			log.Errorf("read theme dir %s info failed: %s", theme.Name(), err.Error())
			c.JSON(http.StatusOK, errmsg.Error())
			return
		}
		if themeInfo.Cover == "" {
			if _, err := os.Stat(path.Join(home, "themes", theme.Name(), "cover.jpg")); err == nil {
				themeInfo.Cover = path.Join("themes", theme.Name(), "cover.jpg")
			}
		}
		themeList = append(themeList, themeInfo)
	}
	c.JSON(http.StatusOK, errmsg.Success(themeList))
}

// GetCurrentThemeID godoc
// @Summary get current theme id
// @Schemes http https
// @Description get current theme id
// @Accept application/json
// @Param Authorization header string false "Authorization"
// @Success 200 {object} errmsg.Response{data=string}
// @Router /api/admin/theme/current/id [get]
func GetCurrentThemeID(c *gin.Context) {
	themeID, err := model.GetOptionByKey(model.OptionKeyBlogTheme)
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
		log.Error("read theme info file failed:", err.Error())
		return cfg, false
	}
	return cfg, true
}
