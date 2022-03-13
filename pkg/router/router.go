package router

import (
	"github.com/gin-gonic/gin"
	adminAssets "github.com/megrez/assets/admin"
	"github.com/megrez/pkg/middleware/cros"
	"github.com/megrez/pkg/router/admin"
	"net/http"
)

var DefaultTheme = "default"

func NewRouter(megrezDir string) (*gin.Engine, error) {
	g := gin.Default()
	g.Use(cros.Cors())
	//g.Use(checkinstall.CheckInstall())
	//
	//theme, err := model.GetOptionByKey(vo.OptionKeyBlogTheme)
	//if err != nil {
	//	theme = DefaultTheme
	//}
	//// load pongo2 for gin
	//g.HTMLRender = pongo2gin.TemplatePath(path.Join(megrezDir, "themes", theme))
	//// route for template
	//view.RouteView(g)
	//// route for admin ui
	g.StaticFS("/admin", http.FS(adminAssets.Static))
	// route for admin api
	admin.RouteAdminAPI(g)
	return g, nil
}
