package service

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/go-redis/redis/v8"

	"github.com/utopizza/go-passport/session/model"
)

func CreateSession(ctx context.Context, userId int64) (*model.SessionData, error) {
	rdb := GetRDB()
	if rdb == nil {
		return nil, redis.ErrClosed
	}

	sessionData := model.NewSessionData(userId)
	val, err := json.Marshal(sessionData)
	if err != nil {
		return nil, err
	}
	sessionJson := string(val)

	set, err := rdb.SetNX(ctx, sessionData.Key, sessionJson, sessionData.ExpireTime).Result()
	if err != nil {
		return nil, err
	}
	if !set {
		return nil, errors.New("failed to set session in redis")
	}

	return sessionData, nil
}

func ReadSession(ctx context.Context, sessionKey string) (*model.SessionData, error) {
	rdb := GetRDB()
	if rdb == nil {
		return nil, redis.ErrClosed
	}

	val, err := rdb.Get(ctx, sessionKey).Result()
	if err != nil && err != redis.Nil {
		return nil, err
	}
	if err == redis.Nil {
		return nil, nil
	}

	var sessionData model.SessionData
	if err := json.Unmarshal([]byte(val), &sessionData); err != nil {
		return nil, err
	}

	return &sessionData, nil
}

func DeleteSession(ctx context.Context, sessionKey string) error {
	rdb := GetRDB()
	if rdb == nil {
		return redis.ErrClosed
	}

	_, err := rdb.Del(ctx, sessionKey).Result()
	if err != nil {
		return err
	}

	return nil
}
