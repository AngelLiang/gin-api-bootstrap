package util

type Response struct {
    Code int         `json:"code"`
    Msg  string      `json:"msg"`
    Data interface{} `json:"data"`
}

func MakeResponse(code int, msg string, data interface{}) Response {
	if code == 0 {
		code = 0
	}
	if msg == "" {
		msg = "操作成功"
	}

    return Response{
        Code: code,
        Msg:  msg,
        Data: data,
    }
}
