package service

import (
	"fmt"
	// "net/http"
	// "errors"
	// "github.com/gin-gonic/gin"
	"gin_api_bootstrap/util"

	"gin_api_bootstrap/model"
	"gin_api_bootstrap/query"
)

func GetUserList(p util.Pagination) ([]*model.User, int64, error) {
	var res []*model.User
	var count int64
	var err error
	ctx := query.DB.Statement.Context
	if (p.Page > 0 && p.PerPage > 0) {
		res, count, err = query.User.WithContext(ctx).FindByPage(p.PerPage, (p.Page-1)*p.PerPage)
	} else {
		res, err = query.User.WithContext(ctx).Find()
		count = query.User.WithContext(ctx).Count()
	}
	fmt.Println(res)
	return res, count, err
}

func GetUserById(userId int64) (*model.User, error) {
	// SELECT * FROM users WHERE id = 10;
	ctx := query.DB.Statement.Context
	user, err := query.User.WithContext(ctx).Where(query.User.ID.Eq(userId)).First()
	return user, err
}
