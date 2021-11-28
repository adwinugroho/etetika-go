package routes

import (
	"context"
	"html/template"
	"io"
	"log"

	"github.com/adwinugroho/etetika-go/config"
	"github.com/adwinugroho/etetika-go/models/request"
	"github.com/go-session/session"
	"github.com/labstack/echo/v4"
)

type (
	getRoutesDao interface {
		// Implements from services
		GetDataUserByEmail(email string) interface{}
	}

	GetRoutes struct {
		services getRoutesDao
		user     *request.User
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
	route := &GetRoutes{services, nil}
	// init front page
	frontPageRoute := e.Group("/")
	frontPageRoute.Use(route.checkSessionUser)
	frontPageRoute.GET("", route.index2)
	frontPageRoute.GET("about", route.about)
	frontPageRoute.GET("blog", route.blog)
	frontPageRoute.GET("cart", route.cart)
	frontPageRoute.GET("checkout", route.checkout)
	frontPageRoute.GET("contact", route.contact)
	frontPageRoute.GET("faq", route.faq)
	frontPageRoute.GET("event", route.event)
	frontPageRoute.POST("login", route.login)
	frontPageRoute.GET("login", route.login)
	frontPageRoute.GET("privacy", route.privacy)
	frontPageRoute.GET("product", route.product)
	frontPageRoute.POST("register", route.register)
	frontPageRoute.GET("register", route.register)
	// init dashboard page
	dashboardRoute := e.Group("/dashboard")
	// dashboardRoute.Use(route.accessDashboard)
	dashboardRoute.POST("", route.indexDashboard, route.validateDashboard)
	dashboardRoute.GET("", route.indexDashboard, route.accessDashboard)
	//dashboardRoute.POST("/event", route.listEvent, route.accessDashboard)
	dashboardRoute.GET("/event/list", route.listEvent, route.accessDashboard)
	dashboardRoute.GET("/logout", route.logout)
	dashboardRoute.GET("/ticket/list", route.listTicket, route.accessDashboard)
	// init process routes
	processRoute := e.Group("/process")
	processRoute.POST("/login", route.processLogin)
	processRoute.POST("/register", route.processRegister)

}

var userContext config.UserContext

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

func (route *GetRoutes) login(c echo.Context) error {
	// log.Println(c.Get("err"))
	store, err := session.Start(context.Background(), c.Response(), c.Request())
	if err != nil {
		return c.JSON(500, "Internal Server Error")
	}
	getSession, _ := store.Get("err")
	store.Delete("err")
	return c.Render(200, "login", map[string]interface{}{
		"title":      "Login | e-Tetika",
		"err":        getSession,
		"email_user": route.user.Email,
	})
}

func (route *GetRoutes) loginPost(c echo.Context) error {
	// log.Println(c.Get("err"))
	var email string
	if c.Get("email") != nil {
		email = c.Get("email").(string)
	}
	return c.Render(200, "login", map[string]interface{}{
		"title":      "Login | e-Tetika",
		"err":        c.Get("err"),
		"email_user": email,
	})
	/*
		{{ if ne .err nil }}
			<h3>Error cause {{ .err }}</h3>
		{{ end }}
	*/
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

func (route *GetRoutes) register(c echo.Context) error {
	store, err := session.Start(context.Background(), c.Response(), c.Request())
	if err != nil {
		return c.JSON(500, "Internal Server Error")
	}
	getSession, _ := store.Get("err")
	store.Delete("err")
	return c.Render(200, "register", map[string]interface{}{
		"title":      "Register | e-Tetika",
		"err":        getSession,
		"email_user": route.user.Email,
	})
}
