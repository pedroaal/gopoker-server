package services

import (
	"encoding/json"

	"github.com/google/uuid"
	"gopkg.in/olahol/melody.v1"
)

type MessageIn struct {
	Action   string `json:"action"`
	Name     string `json:"name"`
	Metric   string `json:"metric"`
	UserName string `json:"userName"`
	UserId   string `json:"userId"`
	RoomId   string `json:"roomId"`
	Content  string `json:"content"`
}

type Res struct {
	RoomId string `json:"roomId"`
	UserId string `json:"userId"`
	Room   Room   `json:"room"`
}

func CreateRoom(session *melody.Session, data MessageIn) Res {
	roomId := uuid.NewString()
	userId := uuid.NewString()

	session.Set("roomId", roomId)
	session.Set("userId", userId)

	AddRoom(AddRoomReq{Id: roomId, Name: data.Name, Metric: data.Metric})
	room := AddUser(AddUserReq{RoomId: roomId, UserId: userId, UserName: data.UserName})

	return Res{RoomId: roomId, UserId: userId, Room: room}
}

func JoinRoom(session *melody.Session, data MessageIn) Res {
	userId := uuid.NewString()

	session.Set("roomId", data.RoomId)
	session.Set("userId", userId)

	room := AddUser(AddUserReq{RoomId: data.RoomId, UserId: userId, UserName: data.UserName})

	return Res{RoomId: data.RoomId, UserId: userId, Room: room}
}

func DeleteUser(session *melody.Session) (Res, bool) {
	roomId, _ := session.Get("roomId")
	userId, _ := session.Get("userId")
	hasUsers := true

	room := RemoveUser(RemoveUserReq{RoomId: roomId.(string), UserId: userId.(string)})

	if len(room.Users) == 0 {
		RemoveRoom(RemoveRoomReq{RoomId: roomId.(string)})
		hasUsers = false
	}

	return Res{Room: room}, hasUsers
}

func SendMe(session *melody.Session, msgIn map[string]interface{}) {
	msgOut, _ := json.Marshal(msgIn)
	session.Write([]byte(msgOut))
}

func SendAll(wss *melody.Melody, session *melody.Session, msgIn map[string]interface{}) {
	msgOut, _ := json.Marshal(msgIn)
	sendToRoomId, _ := session.Get("roomId")
	wss.BroadcastFilter([]byte(msgOut), func(s *melody.Session) bool {
		currentRoomId, _ := s.Get("roomId")
		return currentRoomId == sendToRoomId
	})
}
