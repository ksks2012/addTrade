package logic

import (
	// "time"

	// "github.com/spf13/cast"
	"github.com/aggTrade/internal/model"
)

const (
	MsgTypeNormal    = iota // normal user message
	MsgTypeWelcome          // Current user welcome message
	MsgTypeUserEnter        // User enter
	MsgTypeUserLeave        // User exit
	MsgTypeError            // error message
)

func NewMessage(user *User, content string, clientTime string) *model.StreamMsg {
	message := &model.StreamMsg{

	}
	return message
}

func NewWelcomeMessage(user *User) *model.StreamMsg {
	return &model.StreamMsg{
	}
}

func NewUserEnterMessage(user *User) *model.StreamMsg {
	return &model.StreamMsg{
	}
}

func NewUserLeaveMessage(user *User) *model.StreamMsg {
	return &model.StreamMsg{
	}
}

func NewErrorMessage(content string) *model.StreamMsg {
	return &model.StreamMsg{
	}
}
