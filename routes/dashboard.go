package routes

import (
	"context"
	"net/http"

	"github.com/go-session/session"
	"github.com/labstack/echo/v4"
)

func (route *GetRoutes) indexDashboard(c echo.Context) error {
	// ctx := context.Background()
	// ctxValue := context.WithValue(ctx, userContext, route.user)
	return c.Render(200, "dashboard", map[string]interface{}{
		"title":    "e-Tetika | Dashboard",
		"email":    route.user.Email,
		"role":     route.user.Role,
		"fullname": "Adwin Nugroho Siswoyo",
	})
}

func (route *GetRoutes) listEvent(c echo.Context) error {
	return c.Render(200, "event_list", map[string]interface{}{
		"title":    "e-Tetika | List Event",
		"email":    route.user.Email,
		"role":     route.user.Role,
		"fullname": "Adwin Nugroho Siswoyo",
	})
}

func (route *GetRoutes) listProduct(c echo.Context) error {
	return c.Render(200, "product_list", map[string]interface{}{
		"title":    "e-Tetika | List Product",
		"email":    route.user.Email,
		"role":     route.user.Role,
		"fullname": "Adwin Nugroho Siswoyo",
	})
}

func (route *GetRoutes) listUser(c echo.Context) error {
	return c.Render(200, "user_list", map[string]interface{}{
		"title":    "e-Tetika | List User",
		"email":    route.user.Email,
		"role":     route.user.Role,
		"fullname": "Adwin Nugroho Siswoyo",
	})
}

func (route *GetRoutes) listTicket(c echo.Context) error {
	return c.Render(200, "ticket_list", map[string]interface{}{
		"title":    "e-Tetika | List Ticket",
		"email":    route.user.Email,
		"role":     route.user.Role,
		"fullname": "Adwin Nugroho Siswoyo",
	})
}

func (route *GetRoutes) logout(c echo.Context) error {
	err := session.Destroy(context.Background(), c.Response(), c.Request())
	if err != nil {
		return c.JSON(500, "Internal Server Error")
	}
	return c.Redirect(http.StatusTemporaryRedirect, "http://localhost:9000/login")
}

func (route *GetRoutes) manageEvent(c echo.Context) error {
	return c.Render(200, "event_manage", map[string]interface{}{
		"title":    "e-Tetika | Manage Event",
		"email":    route.user.Email,
		"role":     route.user.Role,
		"fullname": "Adwin Nugroho Siswoyo",
	})
}
