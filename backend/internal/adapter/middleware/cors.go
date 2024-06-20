package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/progate-hackathon-ari/backend/pkg/log"
)

func AllowAllOrigins() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			requestAddr := c.Request().Header.Get("Origin")
			// no origin ignore
			if requestAddr == "" {
				log.Error(c.Request().Context(), "origin is empty")
				return echo.ErrUnauthorized
			}
			// ignore /healthz
			if c.Path() == "/healthz" {
				return next(c)
			}
			log.Info(c.Request().Context(), "origin", "origin", requestAddr)
			c.Response().Header().Set("Access-Control-Allow-Origin", requestAddr)
			c.Response().Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
			c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			c.Response().Header().Set("Access-Control-Max-Age", "3600")
			c.Response().Header().Set("Access-Control-Allow-Credentials", "true")

			return next(c)
		}
	}
}
