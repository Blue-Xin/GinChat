package utils

import "strings"

func IsNotBlank(str string) bool {
	return len(strings.TrimSpace(str)) != 0
}
func IsBlank(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}
