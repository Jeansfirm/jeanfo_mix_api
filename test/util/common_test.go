package util_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"jeanfo_mix/config"
	"jeanfo_mix/util"

	"github.com/stretchr/testify/assert"
)

func TestRedisClient(t *testing.T) {
	// 测试单例模式
	client1 := util.GetRedisClient()
	client2 := util.GetRedisClient()
	// 打印实例信息
	fmt.Printf("Client1: %p Client2: %p ClientType: %T\n", client1, client2, client1)
	assert.Equal(t, client1, client2, "should return same instance")

	// 测试连接有效性
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := client1.Ping(ctx).Err()
	assert.Nil(t, err, "should connect to redis server")

	// 测试配置加载
	cfg := config.GetConfig().Redis
	assert.NotEmpty(t, cfg.Addr, "redis address should not be empty")
	// assert.NotEqual(t, 0, cfg.DB, "redis db should not be 0")

	// 测试基本操作
	key := "test_key"
	value := "test_value"
	err = client1.Set(ctx, key, value, 10*time.Second).Err()
	assert.Nil(t, err, "should set value successfully")

	gotValue, err := client1.Get(ctx, key).Result()
	assert.Nil(t, err, "should get value successfully")
	assert.Equal(t, value, gotValue, "should get correct value")

	// 清理测试数据
	err = client1.Del(ctx, key).Err()
	assert.Nil(t, err, "should delete key successfully")
}
