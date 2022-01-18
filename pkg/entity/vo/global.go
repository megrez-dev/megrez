package vo

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/megrez/pkg/model"
)

type Global struct {
	BlogURL         string
	BlogTitle       string
	BlogDescription string
	BlogCover       string
	HeaderLogo      string
	FooterLogo      string
	PayQRCode       string
	IPCRecord       string
	BlogAge         float64
	Github          string
	QQ              string
	Email           string
	Telegram        string
	BlogBirth       time.Time
	Menus           []*Menu
	Categories      []*BriefCategory
	LatestArticles  []*LatestArticl
	LatestComments  []*LatestComment
	RandomColor     func() string
}

type LatestArticl struct {
	Title string
	URL   string
}

type LatestComment struct {
	Content string
	URL     string
}

func GetLatestArticleFromPO(article *model.Article) (*LatestArticl, error) {
	latestArticle := &LatestArticl{
		Title: article.Title,
	}
	latestArticle.URL = "/article/" + strconv.Itoa(int(article.ID))
	return latestArticle, nil
}

func GetLatestCommentFromPO(comment *model.Comment) (*LatestComment, error) {
	latestComment := &LatestComment{
		Content: comment.Content,
	}
	indexPageSizeStr, err := model.GetOptionByKey(OptionComentsPageSize)
	if err != nil {
		return nil, err
	}
	indexPageSize, err := strconv.Atoi(indexPageSizeStr)
	if err != nil {
		return nil, err
	}
	var rootComments []model.Comment
	// comment for article
	if comment.Type == 1 {
		rootComments, err = model.ListRootCommentsByArticleID(comment.ArticleID, 0, 0)
		if err != nil {
			return nil, err
		}
		var index int
		for i, rootComment := range rootComments {
			if rootComment.ID == comment.ID || rootComment.ID == comment.RootID {
				index = i + 1
				break
			}
		}
		pagination := (index + indexPageSize - 1) / indexPageSize
		url := fmt.Sprintf("/article/%d/comment-page/%d#comment-%d", comment.ArticleID, pagination, comment.ID)
		latestComment.URL = url
	} else if comment.Type == 2 {
		// comment for page
		rootComments, err = model.ListRootCommentsByPageID(comment.PageID, 0, 0)
		if err != nil {
			return nil, err
		}
		var index int
		for i, rootComment := range rootComments {
			if rootComment.ID == comment.ID || rootComment.ID == comment.RootID {
				index = i + 1
				break
			}
		}
		pagination := (index + indexPageSize - 1) / indexPageSize
		page, err := model.GetPageByID(comment.PageID)
		if err != nil {
			return nil, err
		}
		url := fmt.Sprintf("/%s/comment-page/%d#comment-%d", page.Slug, pagination, comment.ID)
		latestComment.URL = url
	}

	return latestComment, nil
}

func GetGlobalOption() (Global, error) {
	global := Global{}
	blogTitle, err := model.GetOptionByKey(OptionKeyBlogTitle)
	if err == nil {
		global.BlogTitle = blogTitle
	}
	blogURL, err := model.GetOptionByKey(OptionKeyBlogURL)
	if err == nil {
		global.BlogURL = blogURL
	}
	blogDescription, err := model.GetOptionByKey(OptionKeyBlogDescription)
	if err == nil {
		global.BlogDescription = blogDescription
	}
	blogCover, err := model.GetOptionByKey(OptionKeyBlogCover)
	if err == nil {
		global.BlogCover = blogCover
	}
	headerLogo, err := model.GetOptionByKey(OptionKeyHeaderLogo)
	if err == nil {
		global.HeaderLogo = headerLogo
	}
	footerLogo, err := model.GetOptionByKey(OptionKeyFooterLogo)
	if err == nil {
		global.FooterLogo = footerLogo
	}
	payQRCode, err := model.GetOptionByKey(OptionKeyPayQRCode)
	if err == nil {
		global.PayQRCode = payQRCode
	}
	ipcRecord, err := model.GetOptionByKey(OptionKeyIPCRecord)
	if err == nil {
		global.IPCRecord = ipcRecord
	}

	blogBirthStr, err := model.GetOptionByKey(OptionKeyBlogBirth)
	if err == nil {
		blogBirth, err := time.Parse("2006-01-02 15:04:05", blogBirthStr)
		if err == nil {
			global.BlogBirth = blogBirth
		}
	}

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
	if err == nil {
		var latestArticles []*LatestArticl
		for _, articlePO := range articlePOs {
			latestArticle, err := GetLatestArticleFromPO(&articlePO)
			if err == nil {
				latestArticles = append(latestArticles, latestArticle)
			}
		}
		global.LatestArticles = latestArticles
	}

	// latest comments
	commentPOs, err := model.ListLatestComments()
	if err == nil {
		var latestComments []*LatestComment
		for _, commentPO := range commentPOs {
			latestComment, err := GetLatestCommentFromPO(&commentPO)
			if err == nil {
				latestComments = append(latestComments, latestComment)
			}
		}
		global.LatestComments = latestComments
	}
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
