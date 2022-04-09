package logic

import (
	"expvar"
	"fmt"
	"log"

	"github.com/aggTrade/global"
	"github.com/aggTrade/internal/model"
)

func init() {
	expvar.Publish("message_queue", expvar.Func(calcMessageQueueLen))
}

func calcMessageQueueLen() interface{} {
	fmt.Println("===len=:", len(Broadcaster.messageChannel))
	return len(Broadcaster.messageChannel)
}

// logic/broadcast.go
// broadcaster
type broadcaster struct {
	// all users
	users map[string]*User

	// All channels are managed in a unified manner, which can avoid external misuse

	enteringChannel chan *User
	leavingChannel  chan *User
	messageChannel  chan model.StreamMsg
}

var Broadcaster = &broadcaster{
	users: make(map[string]*User),

	enteringChannel: make(chan *User),
	leavingChannel:  make(chan *User),
	messageChannel:  make(chan model.StreamMsg, global.MessageQueueLen),
}

// logic/broadcast.go

// Start starts the broadcaster
// needs to be run in a new goroutine because it will not return
func (b *broadcaster) Start() {
	for {
		select {
		case user := <-b.enteringChannel:
			// new user enters
			b.users[user.Token] = user
		case user := <-b.leavingChannel:
			// user leaves
			delete(b.users, user.Token)
			// Avoid goroutine leaks
			user.CloseMessageChannel()
		case msg := <- global.SendMsg:
			// data from websocket api
			for _, user := range b.users {
				user.MessageChannel <- msg
			}
		}
	}
}

func (b *broadcaster) UserEntering(u *User) {
	b.enteringChannel <- u
}

func (b *broadcaster) UserLeaving(u *User) {
	b.leavingChannel <- u
}

func (b *broadcaster) Broadcast(msg model.StreamMsg) {
	if len(b.messageChannel) >= global.MessageQueueLen {
		log.Println("broadcast queue full, message dropped")
	}
	b.messageChannel <- msg
}
