package routes

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (route *GetRoutes) processLogin(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	if email == "admin" && password == "password" {
		//c.Set("email", email)
		err := c.Redirect(http.StatusTemporaryRedirect, "http://localhost:9000/dashboard")
		if err != nil {
			log.Println("error when redirect to dashboard", err)
			return c.JSON(500, "Internal Server Error, Please Contact Customer Service")
		}
	} else {
		log.Println("loggin failed")
		return c.Redirect(http.StatusTemporaryRedirect, "http://localhost:9000/login")
	}
	return nil
}

func (route *GetRoutes) processRegister(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	fullname := c.FormValue("fullname")
	phone := c.FormValue("phone")
	if email == "" || password == "" || fullname == "" || phone == "" {
		return c.Render(200, "register", map[string]interface{}{
			"title":        "Register | e-Tetika",
			"errorMessage": "Invalid Data",
		})
	}
	return c.Redirect(http.StatusTemporaryRedirect, "http://localhost:9000/dashboard")
}
