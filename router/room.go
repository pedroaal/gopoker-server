package router

import (
	"net/http"
	"server/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateRoomReq struct {
	Name     string
	Metric   string
	UserName string
}

func CreateRoom(ctx *gin.Context) {
	var data CreateRoomReq
	if err := ctx.BindJSON(&data); err != nil {
		return
	}

	roomId := uuid.NewString()
	userId := uuid.NewString()

	services.AddRoom(services.AddRoomReq{Id: roomId, Name: data.Name, Metric: data.Metric})
	services.AddUser(services.AddUserReq{RoomId: roomId, UserId: userId, UserName: data.UserName})

	ctx.JSON(http.StatusOK, gin.H{
		"roomId": roomId,
		"userId": userId,
	})
}

type JoinRoomReq struct {
	RoomId   string
	UserName string
}

func JoinRoom(ctx *gin.Context) {
	var data JoinRoomReq
	if err := ctx.BindJSON(&data); err != nil {
		return
	}

	userId := uuid.NewString()

	services.AddUser(services.AddUserReq{RoomId: data.RoomId, UserId: userId, UserName: data.UserName})

	ctx.JSON(http.StatusOK, gin.H{
		"userId": userId,
	})
}
