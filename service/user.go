package service

import (
	// "fmt"
	// "net/http"
	// "errors"
	// "github.com/gin-gonic/gin"
	"gin_api_bootstrap/util"
	"gin_api_bootstrap/serializer"

	"gin_api_bootstrap/model"
	"gin_api_bootstrap/query"
)

// 获取用户列表
func GetUserList(p util.Pagination) ([]*model.User, int64, error) {
	var records []*model.User
	var count int64
	var err error
	ctx := query.DB.Statement.Context
	if (p.Page > 0 && p.PerPage > 0) {
		records, count, err = query.User.WithContext(ctx).FindByPage(p.PerPage, (p.Page-1)*p.PerPage)
	} else {
		records, err = query.User.WithContext(ctx).Find()
		if err != nil {
			return nil, -1, err
		}
		count, err = query.User.WithContext(ctx).Count()
		if err != nil {
			return nil, -1, err
		}
	}
	// fmt.Println(res)
	return records, count, err
}

// 获取用户
func GetUserById(userId int64) (*model.User, error) {
	// SELECT * FROM users WHERE id = 10;
	ctx := query.DB.Statement.Context
	user, err := query.User.WithContext(ctx).Where(query.User.ID.Eq(userId)).First()
	return user, err
}

func UserIsExistByName(name string) bool {
	ctx := query.DB.Statement.Context
	_, err := query.User.WithContext(ctx).Where(query.User.Name.Eq(name)).First()
	if err != nil {
		//record not found
		// fmt.Printf("%+v \n", err)
		return false
	}
	return true
}

func UserIsExistByNameNotId(name string, id int64) bool {
	ctx := query.DB.Statement.Context
	_, err := query.User.WithContext(ctx).Where(
		query.User.Name.Eq(name),
	).Not(query.User.ID.Eq(id)).First()
	if err != nil {
		//record not found
		// fmt.Printf("%+v \n", err)
		return false
	}
	return true
}

// 添加用户
func AddUser(in serializer.AddUserIn) (error) {
	user := model.User{Name: in.Name, Age: in.Age, Balance: in.Balance}
	ctx := query.DB.Statement.Context
	err := query.User.WithContext(ctx).Create(&user)
	return err
}

// 更新用户
func UpdateUser(user model.User, in serializer.UpdateUserIn) error {
	user.Name = in.Name
	user.Age = in.Age
	user.Balance = in.Balance
	ctx := query.DB.Statement.Context
	err := query.User.WithContext(ctx).Save(&user)
	return err
}

// 删除用户
func DeleteUser(userId int64) error {
	ctx := query.DB.Statement.Context
	_, err := query.User.WithContext(ctx).Where(query.User.ID.Eq(userId)).Delete()
	return err
}
