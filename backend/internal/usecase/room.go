package usecase

import (
	"encoding/json"
	"time"

	"github.com/gorilla/websocket"
	"github.com/progate-hackathon-ari/backend/internal/entities/model"
)

var Rooms map[string]*RoomSesison

func init() {
	Rooms = make(map[string]*RoomSesison, 1000)
}

type RoomSesison struct {
	Players map[string]*Client
	Master  string
}

func IsAnswered(roomID string) bool {
	room, ok := Rooms[roomID]
	if !ok {
		return false
	}

	// 全員が答えてたらtrue
	for _, client := range room.Players {
		if !client.IsAnswered {
			return false
		}
	}

	return true
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
		Players: make(map[string]*Client),
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

	Rooms[roomID].Players[client.info.ConnectionID] = client
}

func BroadcastInRoom(roomID string, message []byte) error {
	for _, client := range Rooms[roomID].Players {
		if err := client.ws.WriteMessage(websocket.TextMessage, message); err != nil {
			return err
		}
	}
	return nil
}

type Countdown struct {
	IsDone bool `json:"isDone"`
	Count  int  `json:"count"`
}

// int(秒)カウントした後に
func Counter(roomID string, count int) error {
	for i := range count {
		data, err := json.Marshal(&Countdown{IsDone: false, Count: count - i})
		if err != nil {
			return err
		}

		if err := BroadcastInRoom(roomID, data); err != nil {
			return err
		}

		time.Sleep(1 * time.Second)
	}

	data, err := json.Marshal(&Countdown{IsDone: true, Count: 0})
	if err != nil {
		return err
	}

	return BroadcastInRoom(roomID, data)
}
