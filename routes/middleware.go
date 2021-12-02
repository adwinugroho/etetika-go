package routes

import (
	"context"
	"log"

	"github.com/adwinugroho/etetika-go/models/request"
	"github.com/go-session/session"
	"github.com/labstack/echo/v4"
)

func (route *GetRoutes) checkSessionUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		store, err := session.Start(context.Background(), c.Response(), c.Request())
		if err != nil {
			return c.JSON(500, "Internal Server Error")
		}
		getSession, _ := store.Get("email")
		var email string
		if getSession != nil {
			email = getSession.(string)
		}
		log.Printf("email from checkSessionUser:%v\n", email)
		c.Set("email", email)
		route.user = new(request.User)
		route.user.Email = email
		return next(c)
	}
}

func (route *GetRoutes) validateDashboard(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		email := c.FormValue("email")
		var role string
		if email == "admin@etetika.com" {
			role = "admin"
		} else if email == "user@etetika.com" {
			role = "user"
		}
		route.user = new(request.User)
		store, err := session.Start(context.Background(), c.Response(), c.Request())
		if err != nil {
			return c.JSON(500, "Internal Server Error")
		}
		store.Set("email", email)
		err = store.Save()
		if err != nil {
			return c.JSON(500, "Internal Server Error")
		}
		route.user.Email = email
		route.user.Role = role
		log.Println("email validate dashboard", route.user.Email)
		log.Println("role validate dashboard", route.user.Role)
		return next(c)
	}
}

func (route *GetRoutes) accessDashboard(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		store, err := session.Start(context.Background(), c.Response(), c.Request())
		if err != nil {
			return c.JSON(500, "Internal Server Error")
		}
		getSession, _ := store.Get("email")
		var email string
		if getSession != nil {
			email = getSession.(string)
		}
		var role string
		if email == "admin@etetika.com" {
			role = "admin"
		} else if email == "user@etetika.com" {
			role = "user"
		}
		log.Printf("email from accessDashboard:%v\n", email)
		route.user = new(request.User)
		route.user.Email = email
		route.user.Role = role
		return next(c)
	}
}
