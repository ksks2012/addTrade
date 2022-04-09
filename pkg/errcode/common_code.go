package errcode

var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(10000000, "服務內部錯誤")
	InvalidParams             = NewError(10000001, "導入參數錯誤")
	NotFound                  = NewError(10000002, "找不到")
	UnauthorizedAuthNotExist  = NewError(10000003, "鑑權失敗，找不到對應的AppKey和AppSecret")
	UnauthorizedTokenError    = NewError(10000004, "鑑權失敗，Token錯誤")
	UnauthorizedTokenTimeout  = NewError(10000005, "鑑權失敗，Token超時")
	UnauthorizedTokenGenerate = NewError(10000006, "鑑權失敗，Token生成失敗")
	TooManyRequests           = NewError(10000007, "請求過多")
)
