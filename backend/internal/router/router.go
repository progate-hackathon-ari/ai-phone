package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/progate-hackathon-ari/backend/internal/container"
	"github.com/progate-hackathon-ari/backend/internal/external/bedrock"
	"github.com/progate-hackathon-ari/backend/internal/external/s3"
	"github.com/progate-hackathon-ari/backend/internal/handler"
	"github.com/progate-hackathon-ari/backend/internal/middleware"
	"github.com/progate-hackathon-ari/backend/internal/repository"
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

	corsRoute.Use(middleware.AllowAllOrigins())
	{
		cri := container.Invoke[*usecase.CreateRoomInteractor]()
		// create room
		corsRoute.POST("/room", handler.CreateRoom(cri))

	}

	return router.echo
}

func (r *router) health() {
	r.echo.GET("/healthz", func(c echo.Context) error {
		return c.String(http.StatusOK, `{"status:":"ok"}`)
	})
}
