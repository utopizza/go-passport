package service

import (
	"context"

	"gorm.io/gorm"

	"github.com/utopizza/go-passport/account/model"
)

func CreateAccountInfo(ctx context.Context, accountInfo *model.AccountInfo) error {
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

func ReadAccountInfoByName(ctx context.Context, name string) (*model.AccountInfo, error) {
	db := GetDB()
	if db == nil {
		return nil, gorm.ErrInvalidDB
	}
	var accountInfo model.AccountInfo
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
	ret := db.Delete(&model.AccountInfo{}, id)
	if ret.Error != nil {
		return ret.Error
	}
	return nil
}
