package service

import (
	"GinChat/models"
	"GinChat/utils"
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
)

// @Tags 获取用户
// @Success 200 {string} json{"code":,"message":}
// @Router /user/getUserList [get]
func GetUserList(c *gin.Context) {
	date := make([]*models.UserBasic, 10)
	date = models.GetUserList()
	c.JSON(200, gin.H{
		"message": date,
	})
}

// @Tags 新增用户
// @Success 200 {string} json{"code":,"message":}
// @param name query string false "用户名"
// @param password query string false "密码"
// @param repassword query string false "确认密码"
// @Router /user/createUser [get]
func CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	salt := fmt.Sprintf("%6d", rand.Int31())
	user.Salt = salt
	user.Name = c.Query("name")
	data := models.FindByName(user.Name)
	if utils.IsNotBlank(data.Name) {
		c.JSON(-1, gin.H{
			"message": "用户名已存在",
		})
		return
	}
	password := c.Query("password")
	repasswoed := c.Query("repassword")
	if repasswoed != password {
		c.JSON(-1, gin.H{
			"message": "两次密码不一致",
		})
		return
	}
	//user.PassWord = password
	user.PassWord = utils.MakePassword(password, salt)
	c.JSON(200, gin.H{
		"message": "用户创建成功",
	})
	models.CreateUser(user)
}

// @Tags 删除用户
// @Success 200 {string} json{"code":,"message":}
// @param id query string false "id"
// @Router /user/deleteUser [get]
func DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(-1, gin.H{
			"message": "id错误",
		})
	}
	user.ID = uint(id)
	c.JSON(200, gin.H{
		"message": "删除成功",
	})
	models.DeleteUser(user)
}

// @Tags 修改用户
// @Accept multipart/form-data
// @Success 200 {string} json{"code":,"message":}
// @Param name formData string true "用户名"
// @Param password formData string true "密码"
// @Param email formData string true "邮箱"
// @Param phone formData string true "电话"
// @Param id formData string true "id"
// @Router /user/updateUser [post]
func UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))

	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.PassWord = c.PostForm("password")
	user.Phone = c.PostForm("phone")
	user.Email = c.PostForm("email")
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(-1, gin.H{
			"message": "修改失败，参数不对",
		})
		return
	}
	models.UpdateUser(user)
	c.JSON(200, gin.H{
		"message": "修改成功",
	})
}

// @Tags 修改用户
// @Accept multipart/form-data
// @Success 200 {string} json{"code":,"message":}
// @Param name formData string true "用户名"
// @Param password formData string true "密码"
// @Param email formData string true "邮箱"
// @Param phone formData string true "电话"
// @Param id formData string true "id"
// @Router /user/updateUser [post]
func FindUserPwdAndName(c *gin.Context) {
	name := c.Query("name")
	password := c.Query("password")
	user := models.FindByName(name)
	if utils.IsBlank(user.Name) {
		c.JSON(-1, gin.H{
			"message": "该用户不存在",
		})
		return
	}
	flag := utils.ValidPassword(password, user.Salt, user.PassWord)
	if !flag {
		c.JSON(-1, gin.H{
			"message": "密码不正确",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "登录成功",
	})
}
