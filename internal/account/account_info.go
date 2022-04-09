package account

import (
	"context"

	"gorm.io/gorm"
)

type AccountInfo struct {
	Id          int64 `gorm:"primary_key"`
	UserId      int64
	ScreenName  string
	AvatarUrl   string
	Description string
	Password    string
	Status      int
	Gender      int
	Extra       string
}

func (AccountInfo) TableName() string {
	return "account_info"
}

func CreateAccountInfo(ctx context.Context, accountInfo *AccountInfo) error {
	db := GetDB()
	if db == nil {
		return gorm.ErrInvalidDB
	}
	ret := db.Create(&accountInfo)
	if ret.Error != nil {
		return ret.Error
	}
	return nil
}

func ReadAccountInfoByName(ctx context.Context, name string) (*AccountInfo, error) {
	db := GetDB()
	if db == nil {
		return nil, gorm.ErrInvalidDB
	}
	var accountInfo AccountInfo
	ret := db.Find(&accountInfo, "screen_name = ?", name)
	if ret.Error != nil {
		return nil, ret.Error
	}
	return &accountInfo, nil
}

func DeleteAccountInfoById(ctx context.Context, id int64) error {
	db := GetDB()
	if db == nil {
		return gorm.ErrInvalidDB
	}
	ret := db.Delete(&AccountInfo{}, id)
	if ret.Error != nil {
		return ret.Error
	}
	return nil
}
