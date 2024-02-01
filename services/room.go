package services

import "fmt"

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Room struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Metric string `json:"metric"`
	Users  []User `json:"users"`
}

var rooms = make(map[string]Room)

type AddRoomReq struct {
	Id     string
	Name   string
	Metric string
}

func AddRoom(data AddRoomReq) Room {
	room := Room{
		Id:     data.Id,
		Name:   data.Name,
		Metric: data.Metric,
		Users:  []User{},
	}

	rooms[data.Id] = room

	return room
}

type AddUserReq struct {
	RoomId   string
	UserId   string
	UserName string
}

func AddUser(data AddUserReq) Room {
	room := rooms[data.RoomId]

	room.Users = append(room.Users, User{
		Id:   data.UserId,
		Name: data.UserName,
	})

	rooms[data.RoomId] = room

	return room
}

type RemoveUserReq struct {
	RoomId string
	UserId string
}

func RemoveUser(data RemoveUserReq) Room {
	room := rooms[data.RoomId]

	var users []User
	for _, user := range room.Users {
		if user.Id != data.UserId {
			users = append(users, user)
		}
	}

	room.Users = users

	rooms[data.RoomId] = room

	return room
}

type RemoveRoomReq struct {
	RoomId string
}

func RemoveRoom(data RemoveRoomReq) bool {
	delete(rooms, data.RoomId)
	fmt.Println("Rooms:", len(rooms))
	return true
}
