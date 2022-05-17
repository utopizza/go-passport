package checker

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestEngine(t *testing.T) {
	ctx := new(gin.Context)

	err := new(Engine).Execute(ctx, SMS_CODE_CHECKER, SMS_CODE_KEY_CHECKER)
	if err != nil {
		t.Error(err)
	}

	if !new(Engine).IsChecked(ctx, SMS_CODE_CHECKER, SMS_CODE_KEY_CHECKER) {
		t.Error()
	}
}
