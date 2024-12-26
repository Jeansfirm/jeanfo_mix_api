package util_test

import (
	"fmt"
	"jeanfo_mix/util"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestJwt(t *testing.T) {
	data := map[string]interface{}{"name": "jeanfo", "age": 18}
	validSeconds := 1

	// test generate token
	token, err := util.JwtGenerateToken(data, validSeconds)
	fmt.Println("jwt token:", token)

	assert.Empty(t, err, func() string {
		if err != nil {
			return "generate jwt token fail: " + err.Error()
		}
		return ""
	})

	// test parse token
	data, err = util.JwtParseToken(token)
	assert.Empty(t, err, "parse jwt token fail: %v", err)
	fmt.Printf("parsed data from jwt token: %v\n", data)
	assert.Equal(t, data["age"], float64(18))

	isExpired := util.JwtTokenExpired(token)
	assert.False(t, isExpired, "jwt token should not be expired")

	// test token expire
	time.Sleep(time.Millisecond * 1010)
	isExpired = util.JwtTokenExpired(token)
	assert.True(t, isExpired, "jwt token should be expired")

	data, err = util.JwtParseToken(token)
	assert.NotEmpty(t, err, "jwt token should not parse ok when expired: %v", err)
}
