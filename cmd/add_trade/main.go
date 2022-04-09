package main

import (
	"flag"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/aggTrade/internal/routers"
	"github.com/aggTrade/client"
	"github.com/aggTrade/global"
	"github.com/aggTrade/pkg/logger"
	"github.com/aggTrade/pkg/setting"
	"github.com/aggTrade/server"
)

var (
	port    string
	runMode string
	cfg     string
)

func init() {
	err := setupFlag()
	if err != nil {
		log.Fatalf("init.setupFlag err: %v", err)
	}
	err = setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	err = setupRedis()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}

}

func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	// stopChannel := make(chan os.Signal, 1)
	// signal.Notify(stopChannel, os.Interrupt, unix.SIGTERM)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	go s.ListenAndServe()

	// Subscript data from websockt
	go client.AggTrade()

	// Set sebSocket server
	server.RegisterHandle()
	http.ListenAndServe(global.AppSetting.WebSocketServerPort, nil)
}

func setupFlag() error {
	flag.StringVar(&port, "port", "", "啟動通訊埠")
	flag.StringVar(&runMode, "mode", "", "啟動模式")
	flag.StringVar(&cfg, "config", "etc/", "指定要使用的設定檔路徑")
	flag.Parse()

	return nil
}

func setupSetting() error {
	s, err := setting.NewSetting(strings.Split(cfg, ",")...)
	if err != nil {
		return err
	}
	err = s.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Redis", &global.RedisSetting)
	if err != nil {
		return err
	}

	if port != "" {
		global.ServerSetting.HttpPort = port
	}
	if runMode != "" {
		global.ServerSetting.RunMode = runMode
	}

	global.AppSetting.DefaultContextTimeout *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	return nil
}

func setupRedis() error {
	global.RedisConn = client.NewRedisClient()

	return nil
}

func setupLogger() error {
	fileName := global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   500,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)

	return nil
}
