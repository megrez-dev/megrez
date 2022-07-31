package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/megrez/pkg/api/admin"
	"github.com/megrez/pkg/middleware/jwt"
)

func RouteAdminAPI(g *gin.Engine) {
	auth := g.Group("api/admin")
	auth.Use(jwt.Jwt())

	// api for install
	auth.POST("install", admin.Install)

	// api for upload
	auth.POST("upload", admin.UploadAttachment)
	auth.POST("upload/qcloudcos/ping", admin.PingQcloudCos)
	auth.GET("attachments", admin.ListAttachments)

	// api for themes
	auth.GET("theme/current/config", admin.GetCurrentThemeConfig)
	auth.PUT("theme/current/config", admin.UpdateCurrentThemeConfig)
	auth.GET("theme/current/id", admin.GetCurrentThemeID)
	auth.POST("theme/install", admin.InstallTheme)
	auth.DELETE("theme/:id", admin.DeleteTheme)
	auth.GET("themes", admin.ListThemes)

	//api for admin
	auth.POST("login", admin.Login)

	// api for article
	auth.POST("article", admin.CreateArticle)
	auth.PUT("article/:id", admin.UpdateArticle)
	auth.DELETE("article/:id", admin.DeleteArticle)
	auth.GET("articles", admin.ListArticles)
	auth.GET("article/:id", admin.GetArticle)

	// api for comment
	auth.POST("comment", admin.CreateComment)
	// auth.PUT("comment/:id", admin.UpdateComment)
	auth.DELETE("comment/:id", admin.DeleteComment)
	auth.GET("comments", admin.ListComments)
	// auth.GET("comment/:id", admin.GetComment)

	// // api for category
	auth.POST("category", admin.CreateCategory)
	auth.PUT("category/:id", admin.UpdateCategory)
	auth.DELETE("category/:id", admin.DeleteCategory)
	auth.GET("categories", admin.ListCategories)
	auth.GET("category/:id", admin.GetCategory)

	// // api for tag
	auth.POST("tag", admin.CreateTag)
	// auth.PUT("tag/:id", admin.UpdateTag)
	// auth.DELETE("tag/:id", admin.DeleteTag)
	auth.GET("tags", admin.ListTags)
	// auth.GET("tag/:id", admin.GetTag)

	// api for link
	auth.POST("link", admin.CreateLink)
	auth.PUT("link/:id", admin.UpdateLink)
	auth.DELETE("link/:id", admin.DeleteLink)
	auth.GET("links", admin.ListLinks)

	// api for option
	auth.PUT("option/:key", admin.SetOption)

	// api for setting
	auth.GET("settings", admin.GetSettings)
	auth.PUT("settings", admin.UpdateSettings)

	// // api for journal
	auth.POST("journal", admin.CreateJournal)
	// auth.PUT("journal/:id", admin.UpdateJournal)
	// auth.DELETE("journal/:id", admin.DeleteJournal)
	auth.GET("journals", admin.ListJournals)
	// auth.GET("journal/:id", admin.GetJournal)

	// // api for page
	// auth.POST("page", admin.CreatePage)
	// auth.PUT("page/:id", admin.UpdatePage)
	// auth.DELETE("page/:id", admin.DeletePage)
	// auth.GET("pages", admin.ListPages)
	// auth.GET("page/:id", admin.GetPage)

	// // api for menu
	// auth.POST("menu", admin.CreateMenu)
	// auth.PUT("menu/:id", admin.UpdateMenu)
	// auth.DELETE("menu/:id", admin.DeleteMenu)
	// auth.GET("menus", admin.ListMenu)
	// auth.GET("menu/:id", admin.GetMenu)

	// // api for option
	// auth.PUT("option/:key", admin.SetOption)
	// auth.DELETE("option/:key", admin.DeleteOption)
	// auth.GET("option/:key", admin.GetOption)
}
