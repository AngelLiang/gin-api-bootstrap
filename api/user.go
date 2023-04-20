package api

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin" // 导入 gin 框架

	"gin_api_bootstrap/service"
	"gin_api_bootstrap/util"
)

// @Summary 获取用户列表
// @Schemes http
// @Description 获取用户列表
// @Tags 用户管理
// @Router /api/v1/user/list [get]
// @Accept json
// @Produce json
// @Success 200 {object} util.Response
// @Security ApiKeyAuth
func ListUserApi(c *gin.Context) {
	resp := util.MakeResponse(0, "操作成功", gin.H{})
	c.JSON(http.StatusOK, resp)
}


type GetUserDetailIn struct {
	Id int `form:"id"`
}

type GetUserDetailOut struct {
	Id int `json:"id"`
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
	userIdInt64 := util.Str2Int64(userId)
	user, err := service.GerUserById(userIdInt64)
	if err != nil {
		resp := util.MakeResponse(1, "请求失败", gin.H{})
		c.JSON(http.StatusOK, resp)
		return
	}

	fmt.Println(user)
	resp := util.MakeResponse(0, "操作成功", gin.H{})
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
