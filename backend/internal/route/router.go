package route

import (
	"github.com/labstack/echo/v4"
	"github.com/progate-hackathon-ari/backend/internal/handler"
)

func NewRoute() *echo.Echo {
	engine := echo.New()

	// ws handler
	engine.GET("/ws", handler.Echo)

	// common route
	{
		// health
		engine.GET("/healthz", func(c echo.Context) error {
			return c.JSON(200, "ok")
		})
	}

	return engine
}
