package vo

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/megrez/pkg/log"
	"gorm.io/gorm"

	"github.com/megrez/pkg/model"
)

type Global struct {
	BlogURL         string
	BlogTitle       string
	BlogDescription string
	IPCRecord       string
	Github          string
	QQ              string
	Email           string
	Telegram        string
	BlogBirth       time.Time
	Menus           []*Menu
	Categories      []*BriefCategory
	LatestArticles  []*LatestArticle
	LatestComments  []*LatestComment
	ThemeOptions    map[string]string
	RandomColor     func() string
}

type LatestArticle struct {
	Title string
	URL   string
}

type LatestComment struct {
	Content string
	URL     string
}

func GetLatestArticleFromPO(article *model.Article) (*LatestArticle, error) {
	latestArticle := &LatestArticle{
		Title: article.Title,
	}
	latestArticle.URL = "/article/" + strconv.Itoa(int(article.ID))
	return latestArticle, nil
}

func GetLatestCommentFromPO(comment *model.Comment) (*LatestComment, error) {
	latestComment := &LatestComment{
		Content: comment.Content,
	}
	commentsPageSizeStr, err := model.GetOptionByKey(OptionComentsPageSize)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			commentsPageSizeStr = "10"
		} else {
			log.Error(err)
			return nil, err
		}
	}
	commentsPageSize, err := strconv.Atoi(commentsPageSizeStr)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	var rootComments []model.Comment
	// comment for article
	if comment.Type == 1 {
		rootComments, err = model.ListRootCommentsByArticleID(comment.ArticleID, 0, 0)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		var index int
		for i, rootComment := range rootComments {
			if rootComment.ID == comment.ID || rootComment.ID == comment.RootID {
				index = i + 1
				break
			}
		}
		pagination := (index + commentsPageSize - 1) / commentsPageSize
		url := fmt.Sprintf("/article/%d/comment-page/%d#comment-%d", comment.ArticleID, pagination, comment.ID)
		latestComment.URL = url
	} else if comment.Type == 2 {
		// comment for page
		rootComments, err = model.ListRootCommentsByPageID(comment.PageID, 0, 0)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		var index int
		for i, rootComment := range rootComments {
			if rootComment.ID == comment.ID || rootComment.ID == comment.RootID {
				index = i + 1
				break
			}
		}
		pagination := (index + commentsPageSize - 1) / commentsPageSize
		page, err := model.GetPageByID(comment.PageID)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		url := fmt.Sprintf("/%s/comment-page/%d#comment-%d", page.Slug, pagination, comment.ID)
		latestComment.URL = url
	}

	return latestComment, nil
}

// @return map[string]map[string]string map[tab]map[key]value
func GetThemeOptions() map[string]string {
	theme, err := model.GetOptionByKey(OptionKeyBlogTheme)
	if err != nil {
		log.Error(err)
		return nil
	}
	options, err := model.ListThemeOptionsByThemeID(theme)
	if err != nil {
		log.Error(err)
		return nil
	}
	m := make(map[string]string)
	for _, option := range options {
		m[option.Key] = option.Value
	}
	return m
}

func GetGlobalOption() (Global, error) {
	global := Global{}
	blogTitle, err := model.GetOptionByKey(OptionKeyBlogTitle)
	if err != nil && err != gorm.ErrRecordNotFound {
		return global, err
	}
	global.BlogTitle = blogTitle
	blogURL, err := model.GetOptionByKey(OptionKeyBlogURL)
	if err != nil && err != gorm.ErrRecordNotFound {
		return global, err
	}
	global.BlogURL = blogURL
	blogDescription, err := model.GetOptionByKey(OptionKeyBlogDescription)
	if err != nil && err != gorm.ErrRecordNotFound {
		return global, err
	}
	global.BlogDescription = blogDescription
	ipcRecord, err := model.GetOptionByKey(OptionKeyIPCRecord)
	if err != nil && err != gorm.ErrRecordNotFound {
		return global, err
	}
	global.IPCRecord = ipcRecord
	blogBirthStr, err := model.GetOptionByKey(OptionKeyBlogBirth)
	if err != nil && err != gorm.ErrRecordNotFound {
		return global, err
	}
	blogBirth, err := time.Parse("2006-01-02 15:04:05", blogBirthStr)
	if err != nil {
		return global, err
	}
	global.BlogBirth = blogBirth
	// list menus
	menuPOs, err := model.ListAllMenus()
	if err != nil {
		return global, err
	}
	var menus []*Menu
	for _, mensPO := range menuPOs {
		menus = append(menus, GetMenuFromPO(mensPO))
	}
	global.Menus = menus

	// list category
	categoryPOs, err := model.ListAllCategories()
	if err != nil {
		return global, err
	}
	var categories []*BriefCategory
	for _, categoryPO := range categoryPOs {
		categories = append(categories, GetBriefCategoryFromPO(categoryPO))
	}
	global.Categories = categories

	// latest articles
	articlePOs, err := model.ListLatestArticles()
	if err != nil {
		return global, err
	}
	var latestArticles []*LatestArticle
	for _, articlePO := range articlePOs {
		latestArticle, err := GetLatestArticleFromPO(&articlePO)
		if err != nil {
			return global, err
		}
		latestArticles = append(latestArticles, latestArticle)
	}
	global.LatestArticles = latestArticles

	// latest comments
	commentPOs, err := model.ListLatestComments()
	if err != nil {
		return global, err
	}
	var latestComments []*LatestComment
	for _, commentPO := range commentPOs {
		latestComment, err := GetLatestCommentFromPO(&commentPO)
		if err != nil {
			return global, err
		}
		latestComments = append(latestComments, latestComment)
	}
	themeOptions := GetThemeOptions()
	global.ThemeOptions = themeOptions
	global.LatestComments = latestComments
	global.RandomColor = randomColor
	return global, nil
}

func randomColor() string {
	i := rand.Intn(6)
	switch i {
	case 0:
		return "blue"
	case 1:
		return "purple"
	case 2:
		return "green"
	case 3:
		return "yellow"
	case 4:
		return "red"
	case 5:
		return "orange"
	}
	return "red"
}
