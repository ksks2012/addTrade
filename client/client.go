package client

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/aggTrade/global"
	"github.com/aggTrade/internal/model"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

func AggTrade() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	ws := global.AppSetting.WebSocketAddr + "?streams=" + global.AppSetting.ExchangeClass + "@aggTrade"
	fmt.Printf("websocket: %#v\n", ws)
	conn, _, err := websocket.Dial(ctx, ws, nil)
	if err != nil {
		log.Println("Dial error:", err)
		return
	}
	defer conn.Close(websocket.StatusInternalError, "Internal error!")

	ctx = context.Background()
	for {
		var message model.StreamMsg
		err = wsjson.Read(ctx, conn, &message)
		if err != nil {
			log.Println("receive msg error:", err)
			continue
		}
		val, _ := json.Marshal(message)
		if err := global.RedisConn.Set("streams=btcusdt@aggTrade", val, 0).Err(); err != nil {
			log.Println("client.Set failed", err)
		}
		global.SendMsg <- message
	}

	conn.Close(websocket.StatusNormalClosure, "")
}
