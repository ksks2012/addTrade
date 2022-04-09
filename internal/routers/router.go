package routers

import (
	// "net/http"
	"time"

	"github.com/aggTrade/global"

	v1 "github.com/aggTrade/internal/routers/api/v1"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/aggTrade/internal/middleware"
	"github.com/aggTrade/pkg/limiter"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(
	limiter.LimiterBucketRule{
		Key:          "/auth",
		FillInterval: time.Second,
		Capacity:     10,
		Quantum:      10,
	},
)

func NewRouter() *gin.Engine {
	r := gin.New()
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
	} else {
		r.Use(middleware.AccessLog())
	}
	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	aggTrade_api := v1.NewAggTrade()

	apiv1 := r.Group("/api/v1")
	apiv1.GET("/AggTrade", aggTrade_api.Get)

	return r
}
