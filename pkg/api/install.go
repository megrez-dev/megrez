package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"github.com/megrez/pkg/entity/dto"
	"github.com/megrez/pkg/entity/vo"
	"github.com/megrez/pkg/model"
	"github.com/megrez/pkg/utils/errmsg"
	"log"
	"net/http"
	"time"
)

func Install(c *gin.Context) {
	var data dto.InstallBlogDTO
	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Println("decode install json data failed, ", err.Error())
		c.JSON(http.StatusOK, errmsg.ErrorInvalidParam)
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

	// set option blog description
	err = model.SetOption(vo.OptionKeyBlogDescription, "平凡的日子里，也要闪闪发光✨")
	if err != nil {
		log.Println("set option blog description failed:", err.Error())
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
		ArticleID: article.ID,
		Content:   "Welcome to Megrez!",
		Type:      1,
		Site:      "https://megrez.run",
		Mail:      "admin@megrez.run",
		Author:    "MEGREZ",
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
	}
	pages["link"] = model.Page{
		Name:            "友链",
		Slug:            "links",
		FormatContent:   "这是一个默认的友链页面，您应当写下添加友链的条件，以及申请方式、申请格式等说明，您可以在后台编辑。",
		OriginalContent: "这是一个默认的友链页面，您应当写下添加友链的条件，以及申请方式、申请格式等说明，您可以在后台编辑。",
	}
	pages["journal"] = model.Page{
		Name: "日志",
		Slug: "journal",
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

	c.JSON(http.StatusOK, errmsg.Success(article))
}
