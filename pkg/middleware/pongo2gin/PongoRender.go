package pongo2gin

import (
	"net/http"
	"path"

	"github.com/flosch/pongo2/v4"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

//PongoRender struct init
type PongoRender struct {
	TmplDir string
}

//TemplatePath html files path
func TemplatePath(tmplDir string) *PongoRender {
	return &PongoRender{
		TmplDir: tmplDir,
	}
}

//Instance init
func (p *PongoRender) Instance(name string, data interface{}) render.Render {
	var template *pongo2.Template
	fileName := path.Join(p.TmplDir, name)

	if gin.Mode() == gin.DebugMode {
		template = pongo2.Must(pongo2.FromFile(fileName))
	} else {
		template = pongo2.Must(pongo2.FromCache(fileName))
	}

	return &PongoHTML{
		Template: template,
		Name:     name,
		Data:     data.(pongo2.Context),
	}
}

//PongoHTML strcut
type PongoHTML struct {
	Template *pongo2.Template
	Name     string
	Data     pongo2.Context
}

//Render for gin interface  render override
func (p *PongoHTML) Render(w http.ResponseWriter) error {
	p.WriteContentType(w)
	return p.Template.ExecuteWriter(p.Data, w)
}

//WriteContentType  for gin interface  WriteContentType override
func (p *PongoHTML) WriteContentType(w http.ResponseWriter) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = []string{"text/html; charset=utf-8"}
	}
}
