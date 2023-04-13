package util

type JSONResult struct {
	// 业务码
    Code    int          `json:"code" `
	// 响应信息
    Message string       `json:"message"`
	// 响应数据
    Data    interface{}  `json:"data"`
}
