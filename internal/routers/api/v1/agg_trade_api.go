package v1

import (
	"github.com/aggTrade/client"
	"github.com/aggTrade/global"
	"github.com/aggTrade/pkg/app"
	"github.com/aggTrade/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type AggTrade struct{}

func NewAggTrade() AggTrade {
	return AggTrade{}
}

// @Summary 取得 Redis 中最新一筆
// @Produce json
// @Param content query string false "交易對" maxlength(30))
// @Success 200 {object} model.AggTradeSwagger "成功"
// @Failure 400 {object} errcode.Error "請求錯誤"
// @Failure 500 {object} errcode.Error "內部錯誤"
// @Router /api/v1/AggTrade [get]
func (at AggTrade) Get(c *gin.Context) {
	response := app.NewResponse(c)

	content := c.Query("content")
	key := "streams=" + content + "@aggTrade"
	trade, err := client.GetValueByKey(global.RedisConn, key)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorGetAggTradeFail)
	}
	response.ToResponse(trade)
	return
}
