package handler

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
)

func Echo(ctx echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		for {
			// Write
			err := websocket.Message.Send(ws, "Hello, Client!")
			if err != nil {
				ctx.Logger().Error(err)
			}

			// Read
			msg := ""
			err = websocket.Message.Receive(ws, &msg)
			if err != nil {
				ctx.Logger().Error(err)
			}
			fmt.Printf("%s\n", msg)
		}
	}).ServeHTTP(ctx.Response(), ctx.Request())
	return nil
}
