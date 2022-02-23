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
	var data dto.CreateArticleDTO
	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, errmsg.ErrorInvalidParam)
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
	var data model.Article
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
	}
	err = c.ShouldBindJSON(&data)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, gin.H{})
	}

	err = model.UpdateArticleByID(uint(id), &data)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
	}

	c.JSON(http.StatusOK, errmsg.Success(nil))
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

	c.JSON(http.StatusOK, errmsg.Success(nil))
}

func GetArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
	}
	article, err := model.GetArticleByID(uint(id))
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusOK, errmsg.Error())
	}

	c.JSON(http.StatusOK, errmsg.Success(article))
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
	var articleDTOs []dto.ListArticlesDTO
	for _, article := range articles {
		articleDTO, err := dto.LoadFromModel(article)
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
