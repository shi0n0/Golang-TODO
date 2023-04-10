package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/shi0n0/Golang-TODO/handler"
)

var e = createMux()

func createMux() *echo.Echo {
    e := echo.New()
    return e
}

type TemplateRender struct {
	templates *template.Template
}

func (t *TemplateRender) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}


func main() {
	http.Handle("/", e)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())

	render := &TemplateRender{
		templates: template.Must(template.ParseGlob("view/*.html")),
	}
	e.Renderer = render

	e.GET("/", handler.ShowHTML)
	
	e.Start(":8000")
}