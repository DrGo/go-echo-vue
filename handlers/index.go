package handlers

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

var IndexPageRenderer *Template

func init(){
	IndexPageRenderer = &Template{
		templates: template.Must(template.ParseGlob("public/*.html")),
	}

}

func GetIndexPage(c echo.Context) error {
	
	return c.Render(http.StatusOK, "base.html","")
	
}
