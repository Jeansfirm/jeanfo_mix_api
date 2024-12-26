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
	fmt.Println("test run in: ", string(output))

	appConfig := config.AppConfig
	redis := appConfig.Redis
	assert.NotEmpty(t, redis.Addr, "not redis configured")
	jwtSecret := appConfig.JWTSecret
	assert.NotEmpty(t, jwtSecret, "not jwt secret configured")

	fmt.Printf("jwt secret: %v, type of %T", jwtSecret, jwtSecret)
}
