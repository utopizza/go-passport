package model

import (
	"time"
)

type AccountInfo struct {
	Id             int64 `gorm:"primary_key"`
	UserId         int64
	AccountGroupId int64
	ScreenName     string
	AvatarUrl      string
	Description    string
	Password       string
	Status         int
	Gender         int
	Extra          string
	Idc            string
	CreateTime     time.Time
	ModifyTime     time.Time
}

func (AccountInfo) TableName() string {
	return "account_info"
}
