package account

import (
	"context"
	"testing"
)

func TestCreateAccountInfo(t *testing.T) {
	ctx := context.Background()

	// create
	accountInfo := &AccountInfo{
		Id:         1,
		UserId:     1,
		ScreenName: "utopizza",
		Password:   "321",
		Status:     0,
		Gender:     0,
	}
	if err := CreateAccountInfo(ctx, accountInfo); err != nil {
		t.Fatal(err)
	}

	// read
	ret, err := ReadAccountInfoByName(ctx, "utopizza")
	if err != nil {
		t.Fatal(err)
	}
	if ret == nil || ret.Password != "321" {
		t.Error("read failed")
	}

	// delete
	if err := DeleteAccountInfoById(ctx, 1); err != nil {
		t.Error(err)
	}
}
