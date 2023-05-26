package models

import (
	"GinChat/utils"
	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Identity      string
	Name          string `json:"name"`
	PassWord      string `json:"password"`
	Phone         string `valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email         string `valid:"email"`
	ClientIp      string
	ClientPort    string
	Salt          string
	LoginTime     uint
	HeartbeatTime uint
	LoginOutTime  uint
	IsLogout      bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func FindUserByPwdAndName(name, password string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("name=? and pass_word=?", name, password).First(&user)
	return user
}

func FindByName(name string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("name=?", name).First(&user)
	return user
}
func FindByPhone(phone string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("phone=?", phone).First(&user)
	return user
}
func FindByEmail(email string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("email=?", email).First(&user)
	return user
}

func GetUserList() []*UserBasic {
	date := make([]*UserBasic, 10)
	utils.DB.Find(&date)
	return date
}

func CreateUser(user UserBasic) *gorm.DB {
	return utils.DB.Create(&user)
}
func DeleteUser(user UserBasic) *gorm.DB {
	return utils.DB.Delete(&user)
}
func UpdateUser(user UserBasic) *gorm.DB {
	return utils.DB.Model(&user).Updates(UserBasic{Name: user.Name, PassWord: user.PassWord, Email: user.Email, Phone: user.Phone})
}
