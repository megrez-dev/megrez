package router

import (
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	adminAssets "github.com/megrez/assets/admin"
	"github.com/megrez/pkg/entity/vo"
	"github.com/megrez/pkg/middleware/checkinstall"
	"github.com/megrez/pkg/middleware/cros"
	"github.com/megrez/pkg/middleware/pongo2gin"
	"github.com/megrez/pkg/model"
	"github.com/megrez/pkg/router/admin"
	"github.com/megrez/pkg/router/view"
	dirUtils "github.com/megrez/pkg/utils/dir"
	"go.uber.org/zap"
	"io/fs"
	"net/http"
	"path"
	"strings"
	"time"
)

var DefaultTheme = "default"

func NewRouter(logger *zap.Logger) (*gin.Engine, error) {
	g := gin.New()
	g.Use(cros.Cors())
	g.Use(checkinstall.CheckInstall())
	g.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	g.Use(ginzap.RecoveryWithZap(logger, true))

	theme, err := model.GetOptionByKey(vo.OptionKeyBlogTheme)
	if err != nil {
		theme = DefaultTheme
	}
	// load pongo2 for gin
	home, err := dirUtils.GetOrCreateMegrezHome()
	if err != nil {
		return nil, err
	}
	g.HTMLRender = pongo2gin.TemplatePath(path.Join(home, "themes", theme))
	// route for template
	view.RouteView(g)
	// route for admin API
	admin.RouteAdminAPI(g)
	// route for admin ui
	g.StaticFS("/admin", http.FS(adminAssets.Static))
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
			//显示刷新
			c.Writer.Flush()
			return
		}
	})
	return g, nil
}
