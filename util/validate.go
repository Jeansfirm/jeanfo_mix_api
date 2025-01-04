package util

import (
	"regexp"
)

// IsValidPassword 验证密码是否合法
func IsValidPassword(password string) bool {
	// 至少一个大写字母、一个小写字母和一个数字，长度至少6
	var (
		hasLower  = regexp.MustCompile(`[a-z]`)
		hasUpper  = regexp.MustCompile(`[A-Z]`)
		hasNumber = regexp.MustCompile(`\d`)
		// hasSpecial = regexp.MustCompile(`[@$!%*?&]`)
	)

	if len(password) < 8 {
		return false
	}
	return hasLower.MatchString(password) &&
		hasUpper.MatchString(password) &&
		hasNumber.MatchString(password)
}
