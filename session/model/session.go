package model

import (
	"time"

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
