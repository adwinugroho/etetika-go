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
	indexRoute.GET("", route.index)
	indexRoute.GET("about", route.about)
	indexRoute.GET("contact", route.contact)
	indexRoute.GET("faq", route.faq)
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func (route *GetRoutes) index(c echo.Context) error {
	return c.Render(200, "home", map[string]interface{}{
		"title": "e-Tetika | Wadah Etam Berkespresi",
	})
}

func (route *GetRoutes) about(c echo.Context) error {
	return c.Render(200, "about", map[string]interface{}{
		"title": "About | e-Tetika",
	})
}

func (route *GetRoutes) contact(c echo.Context) error {
	return c.Render(200, "contact", map[string]interface{}{
		"title": "Contact | e-Tetika",
	})
}

func (route *GetRoutes) faq(c echo.Context) error {
	return c.Render(200, "faq", map[string]interface{}{
		"title": "Frequently Asked Question | e-Tetika",
	})
}
