package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	themesAssets "github.com/megrez/assets/themes"
	"github.com/megrez/pkg/entity/dto"
	"github.com/megrez/pkg/log"
	"github.com/megrez/pkg/model"
	dirUtils "github.com/megrez/pkg/utils/dir"
	"github.com/megrez/pkg/utils/errmsg"
	"gorm.io/gorm"
	"net/http"
	"os"
	"path"
	"time"
)

// Install godoc
// @Summary install blog form
// @Schemes http https
// @Description install blog form
// @Accept application/json
// @Param  req body dto.InstallBlogForm true "install blog form"
// @Success 200 {object} errmsg.Response{data=string}
// @Router /api/admin/install [post]
func Install(c *gin.Context) {
	var data dto.InstallBlogForm
	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Error("decode json data failed, ", err.Error())
		c.JSON(http.StatusOK, errmsg.Fail(errmsg.ErrorInvalidParam))
		return
	}
	tx := model.BeginTx()
	// check blog installed
	value, err := model.GetOptionByKey(model.OptionKeyIsInstalled)
	if err == nil && value == "true" {
		c.JSON(http.StatusOK, errmsg.FailMsg("不能重复安装"))
		return
	} else if err != nil && err != gorm.ErrRecordNotFound {
		log.Error("get option blog installed failed:", err.Error())
		tx.Rollback()
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	// set option blog birth
	err = model.SetOption(tx, model.OptionKeyBlogBirth, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		log.Error("set option blog birth failed:", err.Error())
		tx.Rollback()
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	// set option blog title
	err = model.SetOption(tx, model.OptionKeyBlogTitle, data.BlogTitle)
	if err != nil {
		log.Error("set option blog title failed:", err.Error())
		tx.Rollback()
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	// set option blog url
	err = model.SetOption(tx, model.OptionKeyBlogURL, data.BlogURL)
	if err != nil {
		log.Error("set option blog url failed:", err.Error())
		tx.Rollback()
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}

	// set option blog description
	err = model.SetOption(tx, model.OptionKeyBlogDescription, "平凡的日子里，也要闪闪发光✨")
	if err != nil {
		log.Error("set option blog description failed:", err.Error())
		tx.Rollback()
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}

	// set option blog themes
	err = model.SetOption(tx, model.OptionKeyBlogTheme, "default")
	if err != nil {
		log.Error("set option blog themes failed:", err.Error())
		tx.Rollback()
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}

	// check themes dir
	megrezHome, err := dirUtils.GetOrCreateMegrezHome()
	if err != nil {
		log.Error("get megrez home dir failed:", err.Error())
		tx.Rollback()
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	themesPath := path.Join(megrezHome, "themes")
	stat, err := os.Stat(themesPath)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(themesPath, os.ModePerm)
			if err != nil {
				log.Error("create themes dir failed:", err.Error())
				tx.Rollback()
				c.JSON(http.StatusOK, errmsg.Error())
				return
			}
		}
	} else {
		if !stat.IsDir() {
			log.Error("themes dir is not dir")
			tx.Rollback()
			c.JSON(http.StatusOK, errmsg.Error())
			return
		}
	}
	_, err = os.Stat(path.Join(themesPath, "default"))
	if err != nil {
		if os.IsNotExist(err) {
			copyErr := dirUtils.CopyDirFromFS(themesAssets.Static, "default", themesPath)
			if copyErr != nil {
				log.Error("copy default theme failed:", err.Error())
				tx.Rollback()
				c.JSON(http.StatusOK, errmsg.Error())
				return
			}
		} else {
			log.Error("copy default theme failed:", err.Error())
			tx.Rollback()
			c.JSON(http.StatusOK, errmsg.Error())
			return
		}
	} else {
		log.Info("default theme is exist")
	}

	// set default theme options
	cfgFile := path.Join(themesPath, "default", "config.yaml")
	open, err := os.Open(cfgFile)
	if err != nil {
		log.Error("open default theme config file failed:", err.Error())
		tx.Rollback()
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	defer func(open *os.File) {
		err := open.Close()
		if err != nil {
			log.Error("close default theme config file failed:", err.Error())
		}
	}(open)
	themeConfig, ok := getThemeConfig(open)
	if !ok {
		log.Error("get default theme config failed")
		tx.Rollback()
		c.JSON(http.StatusOK, errmsg.FailMsg("主题配置文件格式错误"))
		return
	}
	for _, tab := range themeConfig.Tabs {
		for _, item := range tab.Items {
			option := &model.ThemeOption{
				ThemeID: "default",
				Key:     item.Key,
				Value:   item.Default,
				Type:    item.Type,
			}
			err := model.CreateThemeOption(tx, option)
			if err != nil {
				os.RemoveAll(path.Join(themesPath, "default"))
				log.Errorf("init default theme option %s failed: %s", item.Key, err.Error())
				tx.Rollback()
				c.JSON(http.StatusOK, errmsg.Error())
				return
			}
		}
	}

	admin := model.User{
		ID:       1,
		Username: data.Username,
		Nickname: data.Nickname,
		// TODO: 默认 Avatar
		Avatar:      "https://avatars0.githubusercontent.com/u/8186664?s=460&v=4",
		Description: "平凡的日子里，也要闪闪发光✨",
		Email:       data.Email,
		Password:    data.Password,
		Status:      0,
	}
	err = model.CreateUser(tx, &admin)
	if err != nil {
		log.Error("create admin user failed:", err.Error())
		tx.Rollback()
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}

	// create default category
	category := model.Category{
		Name:        "默认分类",
		Slug:        "default",
		Description: "默认分类",
	}
	err = model.CreateCategory(tx, &category)
	if err != nil {
		log.Error("create default category failed:", err.Error())
		tx.Rollback()
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	// publish hello world article
	article := model.Article{
		Title: "Hello Megrez",
		Slug:  slug.Make("hello-megrez"),
		// TODO: init hallo world article
		Summary:         "Hello Megrez",
		OriginalContent: "Hello Megrez",
		FormatContent:   "Hello Megrez",
		AllowedComment:  true,
		PublishTime:     time.Now(),
	}
	err = model.CreateArticle(tx, &article)
	if err != nil {
		log.Error("create hello world article failed:", err.Error())
		tx.Rollback()
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}

	err = model.CreateArticleCategory(tx, article.ID, category.ID)
	if err != nil {
		log.Error("create articleCategory failed:", err.Error())
		tx.Rollback()
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	// publish hello world comment
	comment := model.Comment{
		ArticleID:  article.ID,
		Content:    "Welcome to Megrez!",
		Type:       model.CommentTypePage,
		Site:       "https://megrez.run",
		Email:      "admin@megrez.run",
		Author:     "MEGREZ",
		CreateTime: time.Now(),
	}
	err = model.CreateComment(tx, &comment)
	if err != nil {
		log.Error("create hello world comment failed:", err.Error())
		tx.Rollback()
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	// create default page for journal,link,about
	pages := make(map[string]model.Page)
	pages["about"] = model.Page{
		Name:            "关于",
		Slug:            "about",
		FormatContent:   "这是一个默认的关于页面，您可以在后台编辑。",
		OriginalContent: "这是一个默认的关于页面，您可以在后台编辑。",
		CreateTime:      time.Now(),
		Type:            model.PageTypeBuildIn,
	}
	pages["link"] = model.Page{
		Name:            "友链",
		Slug:            "links",
		FormatContent:   "这是一个默认的友链页面，您应当写下添加友链的条件，以及申请方式、申请格式等说明，您可以在后台编辑。",
		OriginalContent: "这是一个默认的友链页面，您应当写下添加友链的条件，以及申请方式、申请格式等说明，您可以在后台编辑。",
		CreateTime:      time.Now(),
		Type:            model.PageTypeBuildIn,
	}
	pages["journal"] = model.Page{
		Name:       "日志",
		Slug:       "journal",
		CreateTime: time.Now(),
		Type:       model.PageTypeBuildIn,
	}
	for _, page := range pages {
		err := model.CreatePage(tx, &page)
		if err != nil {
			log.Errorf("create page %s failed: %s\n", page.Name, err.Error())
			tx.Rollback()
			c.JSON(http.StatusOK, errmsg.Error())
			return
		}
	}
	// create default menu for journal,link,about
	menus := make(map[string]model.Menu)
	menus["about"] = model.Menu{
		Name:     "关于",
		Slug:     "about",
		PageID:   pages["about"].ID,
		Priority: 3,
	}
	menus["links"] = model.Menu{
		Name:     "友链",
		Slug:     "links",
		PageID:   pages["link"].ID,
		Priority: 2,
	}
	menus["journal"] = model.Menu{
		Name:     "日志",
		Slug:     "journal",
		PageID:   pages["journal"].ID,
		Priority: 1,
	}
	for _, menu := range menus {
		err := model.CreateMenu(tx, &menu)
		if err != nil {
			log.Errorf("create menu %s failed: %s\n", menu.Name, err.Error())
			tx.Rollback()
			c.JSON(http.StatusOK, errmsg.Error())
			return
		}
	}

	// set option installed
	err = model.SetOption(tx, model.OptionKeyIsInstalled, "true")
	if err != nil {
		log.Error("set option is installed failed, ", err.Error())
		tx.Rollback()
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	tx.Commit()
	c.JSON(http.StatusOK, errmsg.Success(nil))
}
