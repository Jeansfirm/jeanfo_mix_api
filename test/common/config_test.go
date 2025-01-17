package common_test

import (
	"fmt"

	"jeanfo_mix/config"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	cmd := exec.Command("pwd")
	output, _ := cmd.CombinedOutput()
	fmt.Println("TestConfig run in: ", string(output))

	appConfig := config.GetConfig()
	redis := appConfig.Redis
	assert.NotEmpty(t, redis.Addr, "not redis configured")
	jwtSecret := appConfig.JWTSecret
	assert.NotEmpty(t, jwtSecret, "not jwt secret configured")
}
