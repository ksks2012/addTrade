package server

import (
	"log"
	"net/http"

	"nhooyr.io/websocket"

	"github.com/aggTrade/logic"
)

func streamHandleFunc(w http.ResponseWriter, req *http.Request) {
	conn, err := websocket.Accept(w, req, nil)
	if err != nil {
		log.Println("websocket accept error:", err)
		return
	}

	token := req.FormValue("token")
	user := logic.NewUser(conn, token, req.RemoteAddr)

	go user.SendMessage(req.Context())

	logic.Broadcaster.UserEntering(user)
	log.Println("token:", token, "joins chat")

	err = user.ReceiveMessage(req.Context())

	// Execute different Close according to the error during reading
	if err == nil {
		conn.Close(websocket.StatusNormalClosure, "END")
	} else {
		log.Println("read from client error:", err)
		conn.Close(websocket.StatusInternalError, "Read from client error")
	}
}
