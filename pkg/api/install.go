package api

import (
	"github.com/88250/lute"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"github.com/megrez/pkg/entity/dto"
	"github.com/megrez/pkg/model"
	"github.com/megrez/pkg/utils/errmsg"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

func Install(c *gin.Context) {
	// create default category
	category := model.Category{
		Name:        "默认分类",
		Slug:        "default",
		Description: "默认分类",
	}
	err := model.CreateCategory(&category)
	if err != nil {
		log.Println("create default category failed:", err.Error())
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
	}

	err = model.CreateArticleCategory(article.ID, category.ID)
	if err != nil {
		log.Println("create articleCategory failed:", err.Error())
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
	}
	// create page for journal, link, about
	pages := []model.Page{
		{
			Name:    "关于",
			Slug:    "about",
			Cover:   "关于",
			Content: "关于",
		},
		{
			Name:        "日志",
			Slug:        "journal",
			Description: "日志",
			Content:     "日志",
		},
		{
			Title:       "友链",
			Slug:        "link",
			Description: "友链",
			Content:     "友链",
		},
	}
	// create default menu for journal,link,about
	menuAbout := model.Menu{
		Name:        "关于",
		Slug:        "about",
		Description: "关于",
		ParentID:    0,
		Order:       0,
	}
	// create default page for journal,link,about
	//
	var data dto.CreateArticleDTO
	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, errmsg.ERROR_INVALID_PARAM)
		return
	}
	article := data.Transfer2Model()
	article.PublishTime = time.Now()

	// check and generate slug
	if article.Slug == "" {
		article.Slug = slug.Make(article.Title)
		// ensure slug unique
		exist, err := model.GetArticleBySlug(article.Slug)
		if err != nil {
			if err != gorm.ErrRecordNotFound {
				log.Println(err.Error())
				c.JSON(http.StatusOK, errmsg.Error())
			}
		}
		if exist.ID != 0 {
			c.JSON(http.StatusOK, errmsg.Fail(errmsg.ERROR_ARTICLE_SLUG_EXIST))
			return
		}
	}
	// check and generate summary
	if article.Summary == "" {
		l := lute.New()
		length := len([]rune(l.HTML2Text(article.FormatContent)))
		if length > 500 {
			article.Summary = string([]rune(l.HTML2Text(article.FormatContent))[:500])
		} else {
			article.Summary = l.HTML2Text(article.FormatContent)
		}
	}

	// check to generate seo keywords
	if article.SeoKeywords == "" {
		for _, tagID := range data.Tags {
			tag, err := model.GetTagByID(tagID)
			if err != nil {
				log.Println(err.Error())
				c.JSON(http.StatusOK, errmsg.Error())
				return
			}
			if article.SeoKeywords == "" {
				article.SeoKeywords = tag.Name
			} else {
				article.SeoKeywords = article.SeoKeywords + ";" + tag.Name
			}
		}
	}
	// check to generate seo description
	if article.SeoDescription == "" {
		article.SeoDescription = article.Summary
	}

	// create article
	err = model.CreateArticle(&article)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}

	// insert article category table
	for _, categoryID := range data.Categories {
		err := model.CreateArticleCategory(article.ID, categoryID)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusOK, errmsg.Error())
			return
		}
	}
	// insert article tag table
	for _, tagID := range data.Tags {
		err := model.CreateArticleTag(article.ID, tagID)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusOK, errmsg.Error())
			return
		}
	}
	c.JSON(http.StatusOK, errmsg.Success(article))
}
