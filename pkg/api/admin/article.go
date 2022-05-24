package admin

import (
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"

	"github.com/88250/lute"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"

	"github.com/megrez/pkg/entity/dto"
	"github.com/megrez/pkg/log"
	"github.com/megrez/pkg/model"
	"github.com/megrez/pkg/utils/errmsg"
)

func CreateArticle(c *gin.Context) {
	var data dto.ArticleDTO
	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Error(err.Error())
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
				log.Error(err)
				c.JSON(http.StatusOK, errmsg.Error())
				return
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
				log.Error(err)
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

	tx := model.BeginTx()
	// create article
	err = model.CreateArticle(tx, &article)
	if err != nil {
		log.Error(err)
		tx.Rollback()
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}

	// insert article category table
	for _, categoryID := range data.Categories {
		err := model.CreateArticleCategory(tx, article.ID, categoryID)
		if err != nil {
			log.Error(err)
			tx.Rollback()
			c.JSON(http.StatusOK, errmsg.Error())
			return
		}
	}
	// insert article tag table
	for _, tagID := range data.Tags {
		err := model.CreateArticleTag(tx, article.ID, tagID)
		if err != nil {
			log.Error(err)
			tx.Rollback()
			c.JSON(http.StatusOK, errmsg.Error())
			return
		}
	}
	tx.Commit()
	c.JSON(http.StatusOK, errmsg.Success(article))
}

func UpdateArticle(c *gin.Context) {
	var data dto.ArticleDTO
	err := c.ShouldBindJSON(&data)
	if err != nil {
		log.Error(err)
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
				log.Error(err)
				c.JSON(http.StatusOK, errmsg.Error())
				return
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
				log.Error(err)
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

	tx := model.BeginTx()
	// create article
	err = model.UpdateArticleByID(tx, &article)
	if err != nil {
		log.Error(err)
		tx.Rollback()
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
			err := model.DeleteArticleCategory(tx, article.ID, oldCategory.ID)
			if err != nil {
				log.Error(err)
				tx.Rollback()
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
			err := model.CreateArticleCategory(tx, article.ID, newCategoryID)
			if err != nil {
				log.Error(err)
				tx.Rollback()
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
			err := model.DeleteArticleTag(tx, article.ID, oldTag.ID)
			if err != nil {
				log.Error(err)
				tx.Rollback()
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
			err := model.CreateArticleTag(tx, article.ID, newTagID)
			if err != nil {
				log.Error(err)
				tx.Rollback()
				c.JSON(http.StatusOK, errmsg.Error())
				return
			}
		}
	}
	tx.Commit()
	c.JSON(http.StatusOK, errmsg.Success(article))
}

func DeleteArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}

	tx := model.BeginTx()
	err = model.DeleteArticleByID(tx, uint(id))
	if err != nil {
		log.Error(err)
		tx.Rollback()
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	err = model.DeleteArticleCategoriesByArticleID(tx, uint(id))
	if err != nil {
		log.Error(err)
		tx.Rollback()
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	err = model.DeleteArticleTagsByArticleID(tx, uint(id))
	if err != nil {
		log.Error(err)
		tx.Rollback()
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	tx.Commit()
	c.JSON(http.StatusOK, errmsg.Success(nil))
}

func GetArticle(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	article, err := model.GetArticleByID(uint(id))
	articleDTO := dto.ArticleDTO{}
	err = articleDTO.LoadFromModel(article)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}

	c.JSON(http.StatusOK, errmsg.Success(articleDTO))
}

func ListArticles(c *gin.Context) {
	var pageNum, pageSize int
	var err error
	if c.Query("pageNum") == "" {
		pageNum = 1
	} else {
		pageNum, err = strconv.Atoi(c.Query("pageNum"))
		if err != nil {
			log.Error(err)
			c.JSON(http.StatusOK, errmsg.Fail(errmsg.ErrorInvalidParam))
			return
		}
	}
	if c.Query("pageSize") == "" {
		pageSize = 10
	} else {
		pageSize, err = strconv.Atoi(c.Query("pageSize"))
		if err != nil {
			log.Error(err)
			c.JSON(http.StatusOK, errmsg.Fail(errmsg.ErrorInvalidParam))
			return
		}
	}
	articles, err := model.ListAllArticles(pageNum, pageSize)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	var articleDTOs []dto.ArticlesListDTO
	for _, article := range articles {
		articleDTO := dto.ArticlesListDTO{}
		err := articleDTO.LoadFromModel(article)
		if err != nil {
			log.Error(err)
			c.JSON(http.StatusOK, errmsg.Error())
			return
		}
		articleDTOs = append(articleDTOs, articleDTO)
	}
	total, err := model.CountAllArticles()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusOK, errmsg.Error())
		return
	}
	pagination := dto.Pagination{
		List:     articleDTOs,
		Current:  pageNum,
		PageSize: pageSize,
		Total:    total,
	}
	c.JSON(http.StatusOK, errmsg.Success(pagination))
}
