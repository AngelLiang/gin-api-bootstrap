package serializer

type AddUserIn struct {
	// 姓名
	Name string `json:"name"`
	// 年龄
	Age int64 `json:"age"`
	// 余额
	Balance float64 `json:"balance"`
}
