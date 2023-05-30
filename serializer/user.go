package serializer

import (
	"time"
)

type UserRecord struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Age int64 `json:"age"`
	Balance float64 `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserListOut struct {
	Current int `json:"current"`
	Size int    `json:"size"`
	Count int64  `json:"count"`
	Records []UserRecord `json:"records"`
}

type GetUserDetailIn struct {
	Id int `form:"id"`
}

type GetUserDetailOut struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Age int64 `json:"age"`
	Balance float64 `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}


type AddUserIn struct {
	// 姓名
	Name string `json:"name"`
	// 年龄
	Age int64 `json:"age"`
	// 余额
	Balance float64 `json:"balance"`
}


type UpdateUserIn struct {
	// 姓名
	Name string `json:"name"`
	// 年龄
	Age int64 `json:"age"`
	// 余额
	Balance float64 `json:"balance"`
}
