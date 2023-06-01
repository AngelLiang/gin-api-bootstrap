package serializer

import (
	"gin_api_bootstrap/model"
	"time"
)

type UserListRecord struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Age int64 `json:"age"`
	Balance float64 `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserListOut struct {
	PaginationOut
	Records []UserListRecord `json:"records"`
}

func BuildUserListOut(page int, perPage int, count int64, userList []*model.User) *UserListOut {
	out := UserListOut{
		PaginationOut: PaginationOut{
			Current: page,
			Size:    perPage,
			Count:   count,
		},
	}
	for _, item := range userList {
		record := UserListRecord{
			ID: item.ID,
			Name: item.Name,
			Age: item.Age,
			Balance: item.Balance,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		}
		out.Records = append(out.Records, record)
	}
	return &out
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
