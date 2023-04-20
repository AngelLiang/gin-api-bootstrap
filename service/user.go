package service

import (
	// "net/http"
	// "errors"
	// "github.com/gin-gonic/gin"

	"gin_api_bootstrap/model"
	"gin_api_bootstrap/query"
)

type Pagination struct {
    Page      int     `form:"current"`
    PerPage   int     `form:"size"`
}

// func GetUserList(p Pagination) {
// 	qs := query.User
// 	if err := qs.Limit(p.PerPage).Offset((p.Page-1)*p.PerPage).Find(&users).Error; err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{
// 			"err_msg":  err.Error(),
// 		})
// 	}
// }

func GerUserById(userId int64) (*model.User, error) {
	qs := query.User
	// SELECT * FROM users WHERE id = 10;
	// ctx := model.DB.context
	ctx := model.DB.Statement.Context
	user, err := qs.WithContext(ctx).Where(qs.ID.Eq(userId)).First()
	// user, err := qs.Where(qs.ID.Eq(userId)).First()
	return user, err
}
