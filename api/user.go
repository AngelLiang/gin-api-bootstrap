package api

import (
	"fmt"
	"time"
	"net/http"
	"github.com/gin-gonic/gin" // 导入 gin 框架

	"gin_api_bootstrap/service"
	"gin_api_bootstrap/util"
)


type UserListOut struct {
	ID int64 `json:"id"`
	Name string `json:"name"`
	Age int64 `json:"age"`
	Balance float64 `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type UserListDataOut struct {
	Current int `json:"current"`
	Size int    `json:"size"`
	Count int64  `json:"count"`
	Records []UserListOut `json:"records"`
}

// @Summary 获取用户列表
// @Schemes http
// @Description 获取用户列表
// @Tags 用户管理
// @Router /api/v1/user/list [get]
// @Accept json
// @Produce json
// @Param current query int false "当前页数"
// @Param size query int false "每页数量"
// @Success 200 {object} util.Response{data=UserListDataOut}
// @Security ApiKeyAuth
func ListUserApi(c *gin.Context) {
	var p util.Pagination
	if err := c.ShouldBindQuery(&p); err != nil {
		resp := util.MakeResponse(0, "传参错误", gin.H{})
		c.JSON(http.StatusOK, resp)
		return
	}
	fmt.Println(p)
	userList, count, err := service.GetUserList(p)
	if err != nil {
		resp := util.MakeResponse(0, "传参错误", gin.H{})
		c.JSON(http.StatusOK, resp)
		return
	}
	fmt.Println(userList)
	out := UserListDataOut{
		Current: p.Page,
		Size: p.PerPage,
		Count: count,
	}
	for _, item := range userList {
		record := UserListOut{
			ID: item.ID,
			Name: item.Name,
			Age: item.Age,
			Balance: item.Balance,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		}
		out.Records = append(out.Records, record)
	}

	resp := util.MakeResponse(0, "操作成功", out)
	c.JSON(http.StatusOK, resp)
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

// @Summary 获取用户详情
// @Schemes http
// @Description 获取用户详情
// @Tags 用户管理
// @Router /api/v1/user/detail [get]
// @Accept json
// @Produce json
// @Success 200 {object} util.Response{data=GetUserDetailOut}
// @Security ApiKeyAuth
// @Param id query int true "用户id"
func GetUserDetailApi(c *gin.Context) {
	userId := c.Query("id")
	if userId == "" {
		resp := util.MakeResponse(1, "请求失败，参数丢失", gin.H{})
		c.JSON(http.StatusOK, resp)
		return
	}
	user, err := service.GetUserById(util.Str2Int64(userId))
	if err != nil {
		fmt.Println(err)
		resp := util.MakeResponse(1, err.Error(), gin.H{})
		c.JSON(http.StatusOK, resp)
		return
	}

	fmt.Println(user)
	out := GetUserDetailOut{
		ID:user.ID,
		Name:user.Name,
		Age:user.Age,
		Balance:user.Balance,
		CreatedAt:user.CreatedAt,
		UpdatedAt:user.UpdatedAt,
	}

	resp := util.MakeResponse(0, "操作成功", out)
	c.JSON(http.StatusOK, resp)
}


// @Summary 添加用户
// @Schemes http
// @Description 添加用户
// @Tags 用户管理
// @Router /api/v1/user/add [post]
// @Accept json
// @Produce json
// @Success 200 {object} util.Response
// @Security ApiKeyAuth
func AddUserApi(c *gin.Context) {
	resp := util.MakeResponse(0, "操作成功", gin.H{})
	c.JSON(http.StatusOK, resp)
}



// @Summary 更新用户信息
// @Schemes http
// @Description 更新用户信息
// @Tags 用户管理
// @Router /api/v1/user/update [put]
// @Accept json
// @Produce json
// @Success 200 {object} util.Response
// @Security ApiKeyAuth
func UpdateUserApi(c *gin.Context) {
	resp := util.MakeResponse(0, "操作成功", gin.H{})
	c.JSON(http.StatusOK, resp)
}
