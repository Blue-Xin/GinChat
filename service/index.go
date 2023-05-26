package service

import "github.com/gin-gonic/gin"

// @Tags 首页
// @Success 200 {string} Helloworld
// @Router /index [get]
func GetIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "welcome",
	})
}
