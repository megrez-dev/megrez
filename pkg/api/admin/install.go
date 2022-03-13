package admin

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	themeAssets "github.com/megrez/assets/theme"
	"github.com/megrez/pkg/entity/dto"
	"github.com/megrez/pkg/entity/vo"
	"github.com/megrez/pkg/model"
	dirUtils "github.com/megrez/pkg/utils/dir"
	"github.com/megrez/pkg/utils/errmsg"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"time"
)

func Install(c *gin.Context) {
	var data dto.InstallBlogForm
	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Println("decode json data failed, ", err.Error())
		c.JSON(http.StatusOK, errmsg.Fail(errmsg.ErrorInvalidParam))
		return
	}
	// set option blog birth
	err = model.SetOption(vo.OptionKeyBlogBirth, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		log.Println("set option blog birth failed:", err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	// set option blog title
	err = model.SetOption(vo.OptionKeyBlogTitle, data.BlogTitle)
	if err != nil {
		log.Println("set option blog title failed:", err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	// set option blog url
	err = model.SetOption(vo.OptionKeyBlogURL, data.BlogURL)
	if err != nil {
		log.Println("set option blog url failed:", err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}

	// set option blog description
	err = model.SetOption(vo.OptionKeyBlogDescription, "平凡的日子里，也要闪闪发光✨")
	if err != nil {
		log.Println("set option blog description failed:", err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}

	// set option blog theme
	err = model.SetOption(vo.OptionKeyBlogTheme, "default")
	if err != nil {
		log.Println("set option blog theme failed:", err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}

	// check themes dir
	megrezHome, err := dirUtils.GetOrCreateMegrezHome()
	if err != nil {
		log.Println("get megrez home dir failed:", err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	themesPath := path.Join(megrezHome, "themes")
	stat, err := os.Stat(themesPath)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(themesPath, os.ModePerm)
			if err != nil {
				log.Println("create themes dir failed:", err.Error())
				c.JSON(http.StatusOK, errmsg.Error())
				return
			}
		}
	} else {
		if !stat.IsDir() {
			log.Println("themes dir is not dir")
			c.JSON(http.StatusOK, errmsg.Error())
			return
		}
	}

	// copy default theme template to megrez home dir
	dirs, err := themeAssets.Static.ReadDir("default")
	if err != nil {
		log.Println("read default theme template failed:", err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	dest, err := os.Open(themesPath)
	if err != nil {
		log.Println("open themes dir failed:", err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	defer dest.Close()
	for _, dir := range dirs {
		_, err = io.Copy(dest, bytes.NewReader(dir))
		if err != nil {
			log.Println("copy default theme template failed:", err.Error())
			c.JSON(http.StatusOK, errmsg.Error())
			return
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
	err = model.CreateUser(&admin)
	if err != nil {
		log.Println("create admin user failed:", err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}

	// create default category
	category := model.Category{
		Name:        "默认分类",
		Slug:        "default",
		Description: "默认分类",
	}
	err = model.CreateCategory(&category)
	if err != nil {
		log.Println("create default category failed:", err.Error())
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
	err = model.CreateArticle(&article)
	if err != nil {
		log.Println("create hello world article failed:", err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}

	err = model.CreateArticleCategory(article.ID, category.ID)
	if err != nil {
		log.Println("create articleCategory failed:", err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	// publish hello world comment
	comment := model.Comment{
		ArticleID:  article.ID,
		Content:    "Welcome to Megrez!",
		Type:       1,
		Site:       "https://megrez.run",
		Mail:       "admin@megrez.run",
		Author:     "MEGREZ",
		CreateTime: time.Now(),
	}
	err = model.CreateComment(&comment)
	if err != nil {
		log.Println("create hello world comment failed:", err.Error())
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
	}
	pages["link"] = model.Page{
		Name:            "友链",
		Slug:            "links",
		FormatContent:   "这是一个默认的友链页面，您应当写下添加友链的条件，以及申请方式、申请格式等说明，您可以在后台编辑。",
		OriginalContent: "这是一个默认的友链页面，您应当写下添加友链的条件，以及申请方式、申请格式等说明，您可以在后台编辑。",
		CreateTime:      time.Now(),
	}
	pages["journal"] = model.Page{
		Name:       "日志",
		Slug:       "journal",
		CreateTime: time.Now(),
	}
	for _, page := range pages {
		err := model.CreatePage(&page)
		if err != nil {
			log.Printf("create page %s failed: %s\n", page.Name, err.Error())
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
		err := model.CreateMenu(&menu)
		if err != nil {
			log.Printf("create menu %s failed: %s\n", menu.Name, err.Error())
			c.JSON(http.StatusOK, errmsg.Error())
			return
		}
	}

	// set option installed
	err = model.SetOption(vo.OptionKeyIsInstalled, "true")
	if err != nil {
		if err != nil {
			log.Println("set option is installed failed, ", err.Error())
			c.JSON(http.StatusOK, errmsg.Error())
			return
		}
	}

	c.JSON(http.StatusOK, errmsg.Success(nil))
}
