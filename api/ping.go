package api

import (
	"net/http"
	"github.com/gin-gonic/gin" // 导入 gin 框架

	"gin_api_bootstrap/util"
)

// @Summary ping
// @Schemes http
// @Description Ping 测试接口
// @Tags 通用
// @Router /api/v1/ping [get]
// @Accept json
// @Produce json
// @Success 200 {object} util.Response
func Ping(c *gin.Context) {
	util.Log().Debug("ping")
	resp := util.MakeResponse(0, "操作成功", gin.H{})
	c.JSON(http.StatusOK, resp)
}

// @Summary 检查Apikey
// @Schemes http
// @Description 检查Apikey
// @Tags 通用
// @Router /api/v1/checkApikey [get]
// @Accept json
// @Produce json
// @Success 200 {object} util.Response
// @Security ApiKeyAuth
func CheckApikeyApi(c *gin.Context) {
	resp := util.MakeResponse(0, "操作成功", gin.H{})
	c.JSON(http.StatusOK, resp)
}
