package routes

import "github.com/labstack/echo/v4"

func (route *GetRoutes) checkMethodLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if c.Request().Method == "POST" {
			return c.Render(200, "login", map[string]interface{}{
				"title": "Login | e-Tetika",
			})
		}
		return next(c)
	}
}
