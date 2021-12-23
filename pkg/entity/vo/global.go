package vo

import (
	"fmt"
	"log"
	"strconv"

	"github.com/megrez/pkg/dao"
	"github.com/megrez/pkg/entity/po"
)

type Global struct {
	BlogURL         string
	BlogTitle       string
	BlogDescription string
	IPCRecord       string
	Menus           []*Menu
	LatestArticles  []*LatestArticl
	LatestComments  []*LatestComment
}

type LatestArticl struct {
	Title string
	URL   string
}

type LatestComment struct {
	Content string
	URL     string
}

func GetLatestArticleFromPO(article *po.Article) (*LatestArticl, error) {
	latestArticle := &LatestArticl{
		Title: article.Title,
	}
	dao, err := dao.GetDAO()
	if err != nil {
		return latestArticle, err
	}
	blogURL, err := dao.GetOptionByKey(OptionKeyBlogURL)
	if err != nil {
		return latestArticle, err
	}
	latestArticle.URL = blogURL + "/" + strconv.Itoa(int(article.ID))
	return latestArticle, nil
}

func GetLatestCommentFromPO(comment *po.Comment) (*LatestComment, error) {
	dao, err := dao.GetDAO()
	if err != nil {
		return nil, err
	}
	latestComment := &LatestComment{
		Content: comment.Content,
	}
	indexPageSizeStr, err := dao.GetOptionByKey(OptionComentsPageSize)
	if err != nil {
		return nil, err
	}
	indexPageSize, err := strconv.Atoi(indexPageSizeStr)
	if err != nil {
		return nil, err
	}
	var rootComments []po.Comment
	if comment.Type == 1 {
		rootComments, err = dao.ListRootCommentsByArticleID(comment.ArticleID, 0, 0)
		if err != nil {
			return nil, err
		}
	} else if comment.Type == 2 {
		rootComments, err = dao.ListRootCommentsByArticleID(comment.ArticleID, 0, 0)
		if err != nil {
			return nil, err
		}
	}
	var index int
	for i, rootComment := range rootComments {
		log.Println("rootComment.ID", rootComment.ID)
		log.Println("po.RootID", comment.RootID)
		log.Println("po.ID", comment.ID)
		log.Println("=========================")
		if rootComment.ID == comment.ID || rootComment.ID == comment.RootID {
			index = i + 1
			break
		}
	}
	page := (index + indexPageSize - 1) / indexPageSize
	url := fmt.Sprintf("/article/%d/comment-page/%d#comment-%d", comment.ArticleID, page, comment.ID)
	latestComment.URL = url

	return latestComment, nil
}

func GetGlobalOption() (Global, error) {
	global := Global{}
	dao, err := dao.GetDAO()
	if err != nil {
		return global, err
	}
	blogTitle, err := dao.GetOptionByKey(OptionKeyBlogTitle)
	if err == nil {
		global.BlogTitle = blogTitle
	}
	blogURL, err := dao.GetOptionByKey(OptionKeyBlogURL)
	if err == nil {
		global.BlogURL = blogURL
	}
	blogDescription, err := dao.GetOptionByKey(OptionKeyBlogDescription)
	if err == nil {
		global.BlogDescription = blogDescription
	}
	ipcRecord, err := dao.GetOptionByKey(OptionKeyIPCRecord)
	if err == nil {
		global.IPCRecord = ipcRecord
	}

	// list menus
	menuPOs, err := dao.ListAllMenus()
	if err != nil {
		return global, err
	}
	menus := []*Menu{}
	for _, mensPO := range menuPOs {
		menus = append(menus, GetMenuFromPO(mensPO))
	}
	global.Menus = menus

	// latest articles
	articlePOs, err := dao.ListLatestArticles()
	if err == nil {
		latestArticles := []*LatestArticl{}
		for _, articlePO := range articlePOs {
			latestArticle, err := GetLatestArticleFromPO(&articlePO)
			if err == nil {
				latestArticles = append(latestArticles, latestArticle)
			}
		}
		global.LatestArticles = latestArticles
	}

	// latest comments
	commentPOs, err := dao.ListLatestComments()
	if err == nil {
		latestComments := []*LatestComment{}
		for _, commentPO := range commentPOs {
			latestComment, err := GetLatestCommentFromPO(&commentPO)
			if err == nil {
				latestComments = append(latestComments, latestComment)
			}
		}
		global.LatestComments = latestComments
	}
	return global, nil
}
