package api

import (
	// "fmt"
	// "time"
	"net/http"
	"github.com/gin-gonic/gin" // 导入 gin 框架

	"gin_api_bootstrap/serializer"
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
// @Param current query int false "当前页数"
// @Param size query int false "每页数量"
// @Success 200 {object} util.Response{data=serializer.UserRecord}
// @Security ApiKeyAuth
func ListUserApi(c *gin.Context) {
	var p util.Pagination
	if err := c.ShouldBindQuery(&p); err != nil {
		resp := util.MakeResponse(0, "传参错误", gin.H{})
		c.JSON(http.StatusOK, resp)
		return
	}
	// util.Log().Debug(p)
	userList, count, err := service.GetUserList(p)
	if err != nil {
		resp := util.MakeResponse(0, "传参错误", gin.H{})
		c.JSON(http.StatusOK, resp)
		return
	}
	// util.Log().Debug(userList)
	out := serializer.UserListOut{
		Current: p.Page,
		Size: p.PerPage,
		Count: count,
	}
	for _, item := range userList {
		record := serializer.UserRecord{
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

// @Summary 获取用户详情
// @Schemes http
// @Description 获取用户详情
// @Tags 用户管理
// @Router /api/v1/user/detail [get]
// @Accept json
// @Produce json
// @Success 200 {object} util.Response{data=serializer.GetUserDetailOut}
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
		util.Log().Error(err)
		resp := util.MakeResponse(1, err.Error(), gin.H{})
		c.JSON(http.StatusOK, resp)
		return
	}

	util.Log().Debug(user)
	out := serializer.GetUserDetailOut{
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
// @Param string body serializer.AddUserIn false "用户详情"
func AddUserApi(c *gin.Context) {
	var jsonIn serializer.AddUserIn
	if err := c.ShouldBindJSON(&jsonIn); err != nil {
		resp := util.MakeResponse(0, "传参错误", gin.H{})
		c.JSON(http.StatusOK, resp)
		return
	}

	if isExist := service.UserIsExistByName(jsonIn.Name); isExist == true {
		resp := util.MakeResponse(1, "存在相同的姓名", gin.H{})
		c.JSON(http.StatusOK, resp)
		return
	}

	err := service.AddUser(jsonIn)
	if err != nil {
		util.Log().Error(err)
		resp := util.MakeResponse(1, err.Error(), gin.H{})
		c.JSON(http.StatusOK, resp)
		return
	}

	resp := util.MakeResponse(0, "操作成功", gin.H{})
	c.JSON(http.StatusOK, resp)
}



// @Summary 更新用户信息
// @Schemes http
// @Description 更新用户信息
// @Tags 用户管理
// @Router /api/v1/user/update/{id} [put]
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "ID"
// @Param string body serializer.UpdateUserIn false "用户详情"
// @Success 200 {object} util.Response
func UpdateUserApi(c *gin.Context) {
	userId := c.Param("id")
	if userId == "" {
		resp := util.MakeResponse(1, "请求失败，参数丢失", gin.H{})
		c.JSON(http.StatusOK, resp)
		return
	}

	// 获取传参
	var jsonIn serializer.UpdateUserIn
	if err := c.ShouldBindJSON(&jsonIn); err != nil {
		resp := util.MakeResponse(0, "传参错误", gin.H{})
		c.JSON(http.StatusOK, resp)
		return
	}

	if isExist := service.UserIsExistByNameNotId(jsonIn.Name, util.Str2Int64(userId)); isExist == true {
		resp := util.MakeResponse(1, "存在相同的姓名", gin.H{})
		c.JSON(http.StatusOK, resp)
		return
	}

	// 获取对象
	user, err := service.GetUserById(util.Str2Int64(userId))
	if err != nil {
		resp := util.MakeResponse(1, err.Error(), gin.H{})
		c.JSON(http.StatusOK, resp)
		return
	}

	// 更新数据
	err = service.UpdateUser(*user, jsonIn)
	if err != nil {
		resp := util.MakeResponse(1, err.Error(), gin.H{})
		c.JSON(http.StatusOK, resp)
		return
	}

	resp := util.MakeResponse(0, "操作成功", gin.H{})
	c.JSON(http.StatusOK, resp)
}



// @Summary 删除用户
// @Schemes http
// @Description 删除用户
// @Tags 用户管理
// @Router /api/v1/user/delete/{id} [delete]
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "ID"
// @Success 200 {object} util.Response
func DeleteUserApi(c *gin.Context) {
	userId := c.Param("id")
	if userId == "" {
		resp := util.MakeResponse(1, "请求失败，参数丢失", gin.H{})
		c.JSON(http.StatusOK, resp)
		return
	}

	// 获取对象
	_, err := service.GetUserById(util.Str2Int64(userId))
	if err != nil {
		resp := util.MakeResponse(1, err.Error(), gin.H{})
		c.JSON(http.StatusOK, resp)
		return
	}

	// 删除对象
	if err := service.DeleteUser(util.Str2Int64(userId)); err != nil {
		resp := util.MakeResponse(1, err.Error(), gin.H{})
		c.JSON(http.StatusOK, resp)
		return
	}

	resp := util.MakeResponse(0, "操作成功", gin.H{})
	c.JSON(http.StatusOK, resp)
}
