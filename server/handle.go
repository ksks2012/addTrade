package server

import (
	"net/http"

	"github.com/aggTrade/logic"
)

var rootDir string

func RegisterHandle() {
	// Processing broadcast message
	go logic.Broadcaster.Start()

	http.HandleFunc("/stream", streamHandleFunc)

}
