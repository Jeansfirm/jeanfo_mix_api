package util

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"time"

	"jeanfo_mix/config"

	"github.com/go-redis/redis/v8"
)

var (
	redisOnce     sync.Once
	redisInstance *redis.Client
)

// GetExeDir 获取可执行文件所在目录
func GetExeDir() string {
	ex, err := os.Executable()
	if err != nil {
		log.Fatalf("get exe dir failed: %s", err.Error())
	}

	return filepath.Dir(ex)
}

func GetProjRoot() string {
	projDir := config.AppConfig.Web.ProjRoot
	_, err := os.Stat(projDir)
	if err != nil {
		panic(fmt.Sprintf("ProjRoot From Config Not OK: %s - %s", projDir, err.Error()))
	}

	return projDir
}

// GetRedisClient 获取Redis客户端单例
func GetRedisClient() *redis.Client {
	redisOnce.Do(func() {
		cfg := config.GetConfig().Redis
		redisInstance = redis.NewClient(&redis.Options{
			Addr:         cfg.Addr,
			Password:     cfg.Password,
			DB:           cfg.DB,
			DialTimeout:  5 * time.Second,
			ReadTimeout:  3 * time.Second,
			WriteTimeout: 3 * time.Second,
			PoolSize:     20,
			MinIdleConns: 5,
		})

		// 测试连接
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := redisInstance.Ping(ctx).Err(); err != nil {
			log.Fatalf("redis connect failed: %v", err)
		}
	})

	return redisInstance
}

func GenRandomString(length int, onlyLowerCase bool) string {
	const lower = "abcdefghijklmnopqrstuvwxyz"
	const upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const digit = "0123456789"
	choices := lower + upper + digit

	if onlyLowerCase {
		choices = lower
	}
	cLen := len(choices)

	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)

	for i := range b {
		b[i] = choices[rand.Intn(cLen)]
	}

	return string(b)
}

func GenTimeBasedUUID(length int) string {
	now := time.Now()
	timePart := now.Format("20060102150405.000000")

	randomPart := ""
	randomLen := length - len(timePart)
	if randomLen > 0 {
		randomPart = GenRandomString(randomLen, true)
	}

	return fmt.Sprintf("%s.%s", timePart, randomPart)
}
