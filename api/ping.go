package api

import (
	"net/http"
	"github.com/gin-gonic/gin" // 导入 gin 框架
)

// @Summary ping
// @Schemes http
// @Description Ping 测试接口
// @Tags 通用
// @Router /api/v1/ping [get]
// @Accept json
// @Produce json
// @Success 200 {object} util.JSONResult
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"message":  "请求成功",
		"data": gin.H{},
	})
}
