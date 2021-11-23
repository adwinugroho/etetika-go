package routes

import (
	"html/template"
	"io"
	"log"

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
	// read directory
	t := &Template{
		templates: template.Must(template.ParseGlob("views/pages/*.html")),
	}
	e.Renderer = t
	route := &GetRoutes{services}
	// init front page
	frontPageRoute := e.Group("/")
	frontPageRoute.GET("", route.index2)
	frontPageRoute.GET("about", route.about)
	frontPageRoute.GET("blog", route.blog)
	frontPageRoute.GET("cart", route.cart)
	frontPageRoute.GET("checkout", route.checkout)
	frontPageRoute.GET("contact", route.contact)
	frontPageRoute.GET("faq", route.faq)
	frontPageRoute.GET("event", route.event)
	frontPageRoute.GET("privacy", route.privacy)
	frontPageRoute.GET("product", route.product)
	// init dashboard page
	dashboardRoute := e.Group("/dashboard")
	dashboardRoute.GET("", route.index2)

}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tplErr := t.templates.ExecuteTemplate(w, name, data)
	if tplErr != nil {
		log.Printf("error cause:%+v\n", tplErr)
		return t.templates.ExecuteTemplate(w, "404", map[string]interface{}{
			"title": "e-Tetika | Error Halaman Tidak Ditemukan! :(",
		})
	} else {
		return nil
	}
}

func (route *GetRoutes) index2(c echo.Context) error {
	var err error
	if err != nil {
		return c.JSON(404, "Error Not Found")
	}
	return c.Render(200, "home", map[string]interface{}{
		"title": "e-Tetika | Wadah Etam Berkespresi",
	})
}

func (route *GetRoutes) about(c echo.Context) error {
	return c.Render(200, "about", map[string]interface{}{
		"title": "About | e-Tetika",
	})
}

func (route *GetRoutes) blog(c echo.Context) error {
	return c.Render(200, "blog", map[string]interface{}{
		"title": "Blog | e-Tetika",
	})
}

func (route *GetRoutes) cart(c echo.Context) error {
	return c.Render(200, "cart", map[string]interface{}{
		"title": "Cart | e-Tetika",
	})
}

func (route *GetRoutes) checkout(c echo.Context) error {
	return c.Render(200, "checkout", map[string]interface{}{
		"title": "Checkout | e-Tetika",
	})
}

func (route *GetRoutes) contact(c echo.Context) error {
	return c.Render(200, "contact", map[string]interface{}{
		"title": "Contact | e-Tetika",
	})
}

func (route *GetRoutes) event(c echo.Context) error {
	return c.Render(200, "event", map[string]interface{}{
		"title": "Event | e-Tetika",
	})
}

func (route *GetRoutes) faq(c echo.Context) error {
	return c.Render(200, "faq", map[string]interface{}{
		"title": "Frequently Asked Question | e-Tetika",
	})
}

func (route *GetRoutes) privacy(c echo.Context) error {
	return c.Render(200, "privacy", map[string]interface{}{
		"title": "Privacy | e-Tetika",
	})
}

func (route *GetRoutes) product(c echo.Context) error {
	return c.Render(200, "product", map[string]interface{}{
		"title": "Product | e-Tetika",
	})
}
