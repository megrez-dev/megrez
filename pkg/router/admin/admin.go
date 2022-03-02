package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/megrez/pkg/api/admin"
)

func RouteAdmin(g *gin.Engine) {
	auth := g.Group("api/admin")
	// TODO: Jwt middleware
	// admin.Use(middleware.JwtToken())

	//api for install
	auth.POST("install", admin.Install)

	//api for admin
	auth.POST("login", admin.Login)

	// api for article
	auth.POST("article", admin.CreateArticle)
	auth.PUT("article/:id", admin.UpdateArticle)
	auth.DELETE("article/:id", admin.DeleteArticle)
	auth.GET("articles", admin.ListArticles)
	auth.GET("article/:id", admin.GetArticle)

	// // api for comment
	// auth.POST("comment", admin.CreateComment)
	// auth.PUT("comment/:id", admin.UpdateComment)
	// auth.DELETE("comment/:id", admin.DeleteComment)
	// auth.GET("comment", admin.ListComments)
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

	// // api for link
	// auth.POST("link", admin.CreateLink)
	// auth.PUT("link/:id", admin.UpdateLink)
	// auth.DELETE("link/:id", admin.DeleteLink)
	// auth.GET("link", admin.ListLinks)
	// auth.GET("link/:id", admin.GetLink)

	// // api for journal
	// auth.POST("journal", admin.CreateJournal)
	// auth.PUT("journal/:id", admin.UpdateJournal)
	// auth.DELETE("journal/:id", admin.DeleteJournal)
	// auth.GET("journal", admin.ListJournals)
	// auth.GET("journal/:id", admin.GetJournal)

	// // api for page
	// auth.POST("page", admin.CreatePage)
	// auth.PUT("page/:id", admin.UpdatePage)
	// auth.DELETE("page/:id", admin.DeletePage)
	// auth.GET("page", admin.ListPages)
	// auth.GET("page/:id", admin.GetPage)

	// // api for menu
	// auth.POST("menu", admin.CreateMenu)
	// auth.PUT("menu/:id", admin.UpdateMenu)
	// auth.DELETE("menu/:id", admin.DeleteMenu)
	// auth.GET("menu", admin.ListMenu)
	// auth.GET("menu/:id", admin.GetMenu)

	// // api for option
	// auth.PUT("option/:key", admin.SetOption)
	// auth.DELETE("option/:key", admin.DeleteOption)
	// auth.GET("option/:key", admin.GetOption)
}
