package routes

import (
	"context"
	"log"
	"net/http"

	"github.com/go-session/session"
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
		// log.Printf("loggin failed cause:%+v\n", err)
		store, err := session.Start(context.Background(), c.Response(), c.Request())
		if err != nil {
			return c.JSON(500, "Internal Server Error")
		}
		store.Set("err", "Invalid Username/Password")
		err = store.Save()
		if err != nil {
			return c.JSON(500, "Internal Server Error")
		}
		return c.Redirect(http.StatusTemporaryRedirect, "http://localhost:9000/login")
	}
	return nil
}

func (route *GetRoutes) processRegister(c echo.Context) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	fullname := c.FormValue("fullName")
	phone := c.FormValue("phone")
	if email == "" || password == "" || fullname == "" || phone == "" {
		store, err := session.Start(context.Background(), c.Response(), c.Request())
		if err != nil {
			return c.JSON(500, "Internal Server Error")
		}
		store.Set("err", "Invalid Data")
		err = store.Save()
		if err != nil {
			return c.JSON(500, "Internal Server Error")
		}
		return c.Redirect(http.StatusTemporaryRedirect, "http://localhost:9000/register")
	}
	return c.Redirect(http.StatusTemporaryRedirect, "http://localhost:9000/dashboard")
}
