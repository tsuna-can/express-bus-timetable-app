package infrastructure

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func APIKeyAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {

	expectedAPIKey := os.Getenv("API_KEY")

	return func(c echo.Context) error {
		apiKey := c.Request().Header.Get("X-API-Key")

		if apiKey != expectedAPIKey {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid or missing API key")
		}

		return next(c)
	}
}
