package router

import (
	"io/fs"
	"net/http"
	"path"
	"strings"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	adminAssets "github.com/megrez/assets/admin"
	"github.com/megrez/pkg/middleware/checkinstall"
	"github.com/megrez/pkg/middleware/cros"
	"github.com/megrez/pkg/middleware/pongo2gin"
	"github.com/megrez/pkg/model"
	api "github.com/megrez/pkg/router/api"
	"github.com/megrez/pkg/router/view"
	dirUtils "github.com/megrez/pkg/utils/dir"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

var DefaultTheme = "default"

func NewRouter(logger *zap.Logger, debug bool) (*gin.Engine, error) {
	if debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	g := gin.New()
	g.Use(cros.Cors())
	g.Use(checkinstall.CheckInstall())
	g.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	g.Use(ginzap.RecoveryWithZap(logger, true))

	theme, err := model.GetOptionByKey(model.OptionKeyBlogTheme)
	if err != nil {
		theme = DefaultTheme
	}
	// load pongo2 for gin
	home, err := dirUtils.GetOrCreateMegrezHome()
	if err != nil {
		return nil, err
	}
	g.HTMLRender = pongo2gin.TemplatePath(path.Join(home, "themes", theme))
	g.StaticFS("/themes/"+theme, http.Dir(path.Join(home, "themes", theme)))
	// route for template
	view.RouteView(g)
	// route for admin API
	api.RouteAdminAPI(g)
	// route for open API
	api.RouteOpenAPI(g)
	// route for admin ui
	g.StaticFS("/admin", http.FS(adminAssets.Static))
	// route for upload attachments
	uploadHome, err := dirUtils.GetOrCreateUploadHome()
	if err != nil {
		return nil, err
	}
	g.Static("/upload", uploadHome)
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	g.NoRoute(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/admin") {
			//设置响应状态
			c.Writer.WriteHeader(http.StatusOK)
			indexHTML, err := fs.ReadFile(adminAssets.Static, "index.html")
			if err != nil {
				c.Redirect(http.StatusInternalServerError, "/error")
				return
			}
			_, err = c.Writer.Write(indexHTML)
			if err != nil {
				c.Redirect(http.StatusInternalServerError, "/error")
				return
			}
			c.Writer.Header().Add("Accept", "text/html")
			c.Writer.Flush()
			return
		}
	})
	return g, nil
}
