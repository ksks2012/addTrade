package errcode

var (
	ErrorGetAggTradeFail = NewError(20010001, "從 Redis 獲取 AggTrade 失敗")
)
