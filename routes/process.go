package routes

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (route *GetRoutes) processLogin(c echo.Context) error {
	username := c.FormValue("email")
	password := c.FormValue("password")
	if username == "admin" && password == "password" {
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
