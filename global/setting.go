package global

import (
	"github.com/aggTrade/internal/model"
	"github.com/aggTrade/pkg/logger"
	"github.com/aggTrade/pkg/setting"

	"github.com/go-redis/redis"
)

var (
	MessageQueueLen = 1024
	SendMsg         chan model.StreamMsg
	RedisConn       *redis.Client
)

var (
	ServerSetting *setting.ServerSettingS
	AppSetting    *setting.AppSettingS
	RedisSetting  *setting.RedisSettingS
	Logger        *logger.Logger
)

func init() {
	SendMsg = make(chan model.StreamMsg, 10)
}
