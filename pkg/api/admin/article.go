package admin

import (
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/88250/lute"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"

	"github.com/megrez/pkg/entity/dto"
	"github.com/megrez/pkg/model"
	"github.com/megrez/pkg/utils/errmsg"
)

func CreateArticle(c *gin.Context) {
	var data dto.ArticleDTO
	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, errmsg.ErrorInvalidParam)
		return
	}
	article := data.Transfer2Model()
	article.PublishTime = time.Now()
	article.EditTime = time.Now()

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
			c.JSON(http.StatusOK, errmsg.Fail(errmsg.ErrorArticleSlugExist))
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

func UpdateArticle(c *gin.Context) {
	var data dto.ArticleDTO
	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, errmsg.ErrorInvalidParam)
		return
	}
	article := data.Transfer2Model()
	article.EditTime = time.Now()
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
			c.JSON(http.StatusOK, errmsg.Fail(errmsg.ErrorArticleSlugExist))
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
	err = model.UpdateArticleByID(&article)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}

	// list old categories and compare with new categories
	// delete old categories
	oldCategories, err := model.ListCategoriesByArticleID(article.ID)
	for _, oldCategory := range oldCategories {
		exist := false
		for _, newCategoryID := range data.Categories {
			if newCategoryID == oldCategory.ID {
				exist = true
				break
			}
		}
		if !exist {
			err := model.DeleteArticleCategory(article.ID, oldCategory.ID)
			if err != nil {
				log.Println(err.Error())
				c.JSON(http.StatusOK, errmsg.Error())
				return
			}
		}
	}
	// add new categories
	for _, newCategoryID := range data.Categories {
		exist := false
		for _, oldCategory := range oldCategories {
			if oldCategory.ID == newCategoryID {
				exist = true
				break
			}
		}
		if !exist {
			err := model.CreateArticleCategory(article.ID, newCategoryID)
			if err != nil {
				log.Println(err.Error())
				c.JSON(http.StatusOK, errmsg.Error())
				return
			}
		}
	}

	// list old tags and compare with new tags
	// delete old tags
	oldTags, err := model.ListTagsByArticleID(article.ID)
	for _, oldTag := range oldTags {
		exist := false
		for _, newTagID := range data.Tags {
			if newTagID == oldTag.ID {
				exist = true
				break
			}
		}
		if !exist {
			err := model.DeleteArticleTag(article.ID, oldTag.ID)
			if err != nil {
				log.Println(err.Error())
				c.JSON(http.StatusOK, errmsg.Error())
				return
			}
		}
	}
	// add new tags
	for _, newTagID := range data.Tags {
		exist := false
		for _, oldTag := range oldTags {
			if oldTag.ID == newTagID {
				exist = true
				break
			}
		}
		if !exist {
			err := model.CreateArticleTag(article.ID, newTagID)
			if err != nil {
				log.Println(err.Error())
				c.JSON(http.StatusOK, errmsg.Error())
				return
			}
		}
	}
	c.JSON(http.StatusOK, errmsg.Success(article))
}

func DeleteArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
	}

	err = model.DeleteArticleByID(uint(id))
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
	}
	err = model.DeleteArticleCategoriesByArticleID(uint(id))
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
	}
	err = model.DeleteArticleTagsByArticleID(uint(id))
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
	}

	c.JSON(http.StatusOK, errmsg.Success(nil))
}

func GetArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
	}
	article, err := model.GetArticleByID(uint(id))
	articleDTO := dto.ArticleDTO{}
	err = articleDTO.LoadFromModel(article)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
	}

	c.JSON(http.StatusOK, errmsg.Success(articleDTO))
}

func ListArticles(c *gin.Context) {
	pageNum, err := strconv.Atoi(c.Query("pageNum"))
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
	}
	articles, err := model.ListAllArticles(pageNum, pageSize)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
	}
	var articleDTOs []dto.ArticlesListDTO
	for _, article := range articles {
		articleDTO := dto.ArticlesListDTO{}
		err := articleDTO.LoadFromModel(article)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusOK, errmsg.Error())
		}
		articleDTOs = append(articleDTOs, articleDTO)
	}
	total, err := model.CountAllArticles()
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
	}
	pagination := dto.Pagination{
		List:     articleDTOs,
		Current:  pageNum,
		PageSize: pageSize,
		Total:    total,
	}
	c.JSON(http.StatusOK, errmsg.Success(pagination))
}
