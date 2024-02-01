package main

import (
	"encoding/json"
	"fmt"
	"server/services"

	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

func main() {
	server := gin.Default()
	wss := melody.New()

	// server.POST("/create-room", router.CreateRoom)
	// server.POST("/join-room", router.CreateRoom)
	// server.POST("/vote", router.CreateRoom)
	// server.POST("/clear-votes", router.CreateRoom)

	server.GET("/ws", func(ctx *gin.Context) {
		wss.HandleRequest(ctx.Writer, ctx.Request)
	})

	wss.HandleMessage(func(session *melody.Session, msgIn []byte) {
		var data services.MessageIn

		if error := json.Unmarshal(msgIn, &data); error != nil {
			fmt.Println("error: ", error)
		}
		fmt.Println("msgIn:", data)

		switch action := data.Action; action {
		case "create":
			res := services.CreateRoom(session, data)
			services.SendMe(session, map[string]interface{}{"type": "session", "roomId": res.RoomId, "userId": res.UserId})
			services.SendAll(wss, session, map[string]interface{}{"type": "room", "room": res.Room})
		case "join":
			res := services.JoinRoom(session, data)
			services.SendMe(session, map[string]interface{}{"type": "session", "roomId": res.RoomId, "userId": res.UserId})
			services.SendAll(wss, session, map[string]interface{}{"type": "room", "room": res.Room})
		case "vote":
			services.SendAll(wss, session, map[string]interface{}{"type": "vote", "content": data.Content, "userId": data.UserId})
		case "clear":
			services.SendAll(wss, session, map[string]interface{}{"type": "clear"})
		default:
			fmt.Println("action not found")
		}
	})

	wss.HandleDisconnect(func(session *melody.Session) {
		if res, ok := services.DeleteUser(session); ok {
			services.SendAll(wss, session, map[string]interface{}{"type": "room", "room": res.Room})
		}
	})

	server.Run(":3000")
}
