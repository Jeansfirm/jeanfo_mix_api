package util

import (
	"regexp"
)

// IsValidPassword 验证密码是否合法
func IsValidPassword(password string) bool {
	// 至少一个大写字母、一个小写字母和一个数字，长度至少6
	var passwordRegex = regexp.MustCompile(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[A-Za-z\d]{6,}$`)
	return passwordRegex.MatchString(password)
}
