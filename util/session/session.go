package session_util

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"jeanfo_mix/util"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

type _SessionMeta struct {
	SessionID       string `json:"SessionID"`
	SessionExpireAt int64  `json:"SessionExpireAt"`
	SessionLoaded   bool
	SessionDeleted  bool
}

type SessionData struct {
	_SessionMeta

	UserID   int    `json:"UserID"`
	UserName string `json:"UserName"`
	Role     int    `json:"Role"`
}

const (
	SessionKeyPrefix string = "JMP-SESS"
	SessionTTL              = 7 * 24 * 60 * 60 // seconds
)

func genSessionID() string {
	return uuid.New().String()[24:]
}

func (sd *SessionData) RedisKey() string {
	redisKey := fmt.Sprintf("%s-%d-%s-%s", SessionKeyPrefix, sd.UserID, sd.UserName, sd.SessionID)
	return redisKey
}

func (sd *SessionData) Save() error {
	return sd.save(false)
}

func (sd *SessionData) SaveResetTTL() error {
	return sd.save(true)
}

func (sd *SessionData) save(resetTTL bool) error {
	if sd.SessionID == "" {
		sd.SessionID = genSessionID()
	}

	timeNow := time.Now().Unix()
	if sd.SessionExpireAt == 0 || resetTTL {
		sd.SessionExpireAt = timeNow + SessionTTL
	}
	realTTL := sd.SessionExpireAt - timeNow

	redisKey := sd.RedisKey()
	data, err := json.Marshal(sd)
	if err != nil {
		return errors.New("fail to marshal session data to json: " + err.Error())
	}

	ctx := context.Background()
	err = util.GetRedisClient().Set(ctx, redisKey, data, time.Duration(realTTL)*time.Second).Err()
	if err != nil {
		return errors.New("fail to save session data to redis: " + err.Error())
	}
	sd.SessionLoaded = true
	sd.SessionDeleted = false

	return nil
}

func (sd *SessionData) Load() error {
	redisKey := sd.RedisKey()
	redisClient := util.GetRedisClient()

	ctx := context.Background()
	data, err := redisClient.Get(ctx, redisKey).Result()
	if err != nil {
		if err == redis.Nil {
			return errors.New("session不存在或者登录已过期")
		} else {
			return errors.New("redis读取session失败: " + err.Error())
		}
	}

	err = json.Unmarshal([]byte(data), sd)
	if err != nil {
		return errors.New("fail to json unmarshal redis session data: " + err.Error())
	}
	sd.SessionLoaded = true
	sd.SessionDeleted = false

	return nil
}

func (sd *SessionData) Delete() error {
	ctx := context.Background()
	err := util.GetRedisClient().Del(ctx, sd.RedisKey()).Err()
	if err != nil {
		return errors.New("redis删除session失败: " + err.Error())
	}
	sd.SessionDeleted = true
	sd.SessionLoaded = false

	return nil
}

func ClearUserSession(userID int) {
	// redisKeyPrefix := fmt.Sprintf("%s-%d-", SessionKeyPrefix, userID)
}

func ClearAllSession() error {
	redisKeyPrefix := SessionKeyPrefix + "-*"
	redisClient := util.GetRedisClient()

	//使用redis scan命令查找所有符合前缀的keys
	var cursor uint64
	ctx := context.Background()
	for {
		keys, cursor := redisClient.Scan(ctx, cursor, redisKeyPrefix, 100).Val()
		if len(keys) > 0 {
			fmt.Println("clearing these session keys: ", keys)
			_, err := redisClient.Del(ctx, keys...).Result()
			if err != nil {
				return err
			}
		}

		if cursor == 0 {
			break
		}
	}

	return nil
}
