package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/progate-hackathon-ari/backend/internal/adapter/gateway/bedrock"
	"github.com/progate-hackathon-ari/backend/internal/adapter/gateway/repository"
	"github.com/progate-hackathon-ari/backend/internal/adapter/gateway/s3"
	"github.com/progate-hackathon-ari/backend/internal/adapter/handler"
	"github.com/progate-hackathon-ari/backend/internal/container"
	"github.com/progate-hackathon-ari/backend/internal/usecase"
)

type router struct {
	echo *echo.Echo
}

func NewRouter() http.Handler {
	echo := echo.New()
	router := &router{
		echo: echo,
	}

	router.health()

	repo := container.Invoke[repository.DataAccess]()
	s3 := container.Invoke[s3.S3]()
	bedrock := container.Invoke[bedrock.Bedrock]()

	router.echo.GET("/game", handler.SocketGameRoom(repo, s3, bedrock))

	corsRoute := router.echo.Group("")

	corsRoute.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))
	{
		cri := container.Invoke[*usecase.RoomInteractor]()
		// create room
		corsRoute.POST("/room", handler.CreateRoom(cri))
		corsRoute.POST("/room/:room_id", handler.UpdateRoom(cri))

	}

	return router.echo
}

func (r *router) health() {
	r.echo.GET("/healthz", func(c echo.Context) error {
		return c.String(http.StatusOK, `{"status:":"ok"}`)
	})
}
