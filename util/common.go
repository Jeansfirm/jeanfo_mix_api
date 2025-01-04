package util

import (
	"context"
	"log"
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
