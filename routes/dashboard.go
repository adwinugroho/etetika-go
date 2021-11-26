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
		"title": "e-Tetika | Dashboard",
		"email": route.user.Email,
	})
}

func (route *GetRoutes) listEvent(c echo.Context) error {
	return c.Render(200, "list_event", map[string]interface{}{
		"title": "e-Tetika | List Event",
		"email": route.user.Email,
	})
}

func (route *GetRoutes) logout(c echo.Context) error {
	err := session.Destroy(context.Background(), c.Response(), c.Request())
	if err != nil {
		return c.JSON(500, "Internal Server Error")
	}
	return c.Redirect(http.StatusTemporaryRedirect, "http://localhost:9000/login")
}
