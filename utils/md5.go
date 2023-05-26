package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// 小写
func Md5Encode(data string) string {
	hash := md5.New()
	hash.Write([]byte(data))
	tempStr := hash.Sum(nil)
	return hex.EncodeToString(tempStr)
}

// 大写
func MD5Encode(data string) string {
	return strings.ToUpper(Md5Encode(data))
}

// 加密
func MakePassword(plainPwd, salt string) string {
	return Md5Encode(plainPwd + salt)
}

// 加密
func ValidPassword(password, salt, plainPwd string) bool {
	return Md5Encode(password+salt) == plainPwd
}
