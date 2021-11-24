package routes

import "github.com/labstack/echo/v4"

func (route *GetRoutes) indexDashboard(c echo.Context) error {
	var err error
	if err != nil {
		return c.JSON(404, "Error Not Found")
	}
	return c.Render(200, "dashboard", map[string]interface{}{
		"title": "e-Tetika | Dashboard",
	})
}
