package usecase

import (
	"github.com/gorilla/websocket"
	"github.com/progate-hackathon-ari/backend/internal/entities/model"
)

var Rooms map[string]*RoomSesison

func init() {
	Rooms = make(map[string]*RoomSesison, 1000)
}

type RoomSesison struct {
	Players map[*Client]bool
	Master  string
}

func IsMaster(roomID, connectionID string) bool {
	room, ok := Rooms[roomID]
	if !ok {
		return false
	}

	return room.Master == connectionID
}

func NewRoomSession(roomID string, masterID string) {
	Rooms[roomID] = &RoomSesison{
		Players: make(map[*Client]bool),
		Master:  masterID,
	}
}

func AddClient(ws *websocket.Conn, info *model.ConnectedPlayer, roomID string) {
	client := &Client{
		ws:   ws,
		info: info,
	}
	if _, ok := Rooms[roomID]; !ok {
		NewRoomSession(roomID, info.ConnectionID)
	}

	Rooms[roomID].Players[client] = true
}

func BroadcastInRoom(roomID string, message []byte) error {
	for client := range Rooms[roomID].Players {
		if err := client.ws.WriteMessage(websocket.TextMessage, message); err != nil {
			return err
		}
	}
	return nil
}
