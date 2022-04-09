package session

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

const (
	DefaultExpireTime = time.Minute * 1
)

type SessionData struct {
	Key        string
	Version    string
	UserId     int64
	ExpireTime time.Duration
}

func NewSessionData(userId int64) *SessionData {
	return &SessionData{
		Key:        uuid.New().String(),
		Version:    "1.0",
		UserId:     userId,
		ExpireTime: DefaultExpireTime,
	}
}

func CreateSession(ctx context.Context, userId int64) (*SessionData, error) {
	rdb := GetRDB()
	if rdb == nil {
		return nil, redis.ErrClosed
	}

	sessionData := NewSessionData(userId)
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

func ReadSession(ctx context.Context, sessionKey string) (*SessionData, error) {
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

	var sessionData SessionData
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
