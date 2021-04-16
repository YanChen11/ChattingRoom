package library

import (
	"ChattingRoom/structure"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var (
	upgrader 	websocket.Upgrader
	Clients		map[*websocket.Conn]string
	Broadcast	chan structure.Message
)

func init() {
	upgrader = websocket.Upgrader{
		ReadBufferSize:    1024,
		WriteBufferSize:   1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	Clients = make(map[*websocket.Conn]string)
	Broadcast = make(chan structure.Message)
}

// 处理连接
func Chat(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	var msg structure.Message

	for {
		err = ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error %v", err)
			delete(Clients, ws)
			break
		}
		if _, ok := Clients[ws]; !ok {
			Clients[ws] = msg.User
		}
		fmt.Printf("%+v", Clients)
		fmt.Println("服务端接收消息", msg)
		if msg.Info == "" {
			continue
		}

		Broadcast <- msg
	}
}

// 分发消息
func DistributeMsg() {
	fmt.Println("distribute")
	for {
		msg := <- Broadcast
		fmt.Println("准备向客户端发送消息", msg)

		for ws, user := range Clients {
			fmt.Println(user)
			if user == msg.User {
				continue
			}
			err := ws.WriteJSON(msg)
			if err != nil {
				fmt.Printf("error %v", err)
				_ = ws.Close()
				fmt.Printf("user %s disconnected", Clients[ws])
				delete(Clients, ws)
			}
		}
	}
}