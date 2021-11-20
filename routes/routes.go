package routes

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type (
	getRoutesDao interface {
		// Implements from services
	}

	GetRoutes struct {
		services getRoutesDao
	}
)

type Template struct {
	templates *template.Template
}

func Init(e *echo.Echo, services getRoutesDao) {
	t := &Template{
		templates: template.Must(template.ParseGlob("views/front-end/*.html")),
	}
	e.Renderer = t
	route := &GetRoutes{services}
	// init index
	indexRoute := e.Group("/")
	indexRoute.GET("", route.indexHandler)
	indexRoute.GET("about", route.aboutHandler)
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func (route *GetRoutes) indexHandler(c echo.Context) error {
	return c.Render(200, "home", map[string]interface{}{
		"name": "e-Tetika home",
	})
}

func (route *GetRoutes) aboutHandler(c echo.Context) error {
	return c.Render(200, "about", map[string]interface{}{
		"name": "About | e-Tetika",
	})
}
