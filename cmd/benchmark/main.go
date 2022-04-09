package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/aggTrade/internal/model"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

var (
	userNum       int           // number of users
	loginInterval time.Duration // User login time interval
	msgInterval   time.Duration // The interval for sending messages to the same user
)

func init() {
	flag.IntVar(&userNum, "u", 500, "Number of logged-in users")
	flag.DurationVar(&loginInterval, "l", 5e9, "User login time interval")
	flag.DurationVar(&msgInterval, "m", 1*time.Minute, "User sending message interval")
}

func main() {
	flag.Parse()

	for i := 0; i < userNum; i++ {
		go UserConnect("user" + strconv.Itoa(i))
		time.Sleep(loginInterval)
	}

	select {}
}

func UserConnect(nickname string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	conn, _, err := websocket.Dial(ctx, "ws://127.0.0.1:2022/stream?token=" + nickname, nil)
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
		fmt.Printf("Received server response: %#v\n", message.Data)
	}

	conn.Close(websocket.StatusNormalClosure, "")
}
