package client

import (
	"github.com/aggTrade/internal/model"
	"encoding/json"
	"log"

	"github.com/aggTrade/global"
	"github.com/go-redis/redis"
)

func NewRedisClient() *redis.Client { // 實體化redis.Client 並返回實體的位址
	client := redis.NewClient(&redis.Options{
		Addr:     global.RedisSetting.Host,
		Password: global.RedisSetting.Password, // no password set
		DB:       global.RedisSetting.DBName,   // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Fatalf("open redis error: %v", err)
	}
	return client
}

func GetValueByKey(conn *redis.Client, key string) (interface{}, error) {
	val, err := conn.Get(key).Result()
	if err != nil {
		log.Println("GetValueByKey failed", err)
		return nil, err
	}
	msg := model.StreamMsg{}
	str := []byte(val)
	err = json.Unmarshal(str, &msg)
	if err != nil {
		log.Println("json marshal failed", err)
		return nil, err
	}
	return msg, nil
}
