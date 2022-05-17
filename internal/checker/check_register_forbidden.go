package checker

import (
	"github.com/gin-gonic/gin"
)

type RegisterForbiddenChecker struct{}

func (c *RegisterForbiddenChecker) DoCheck(ctx *gin.Context) error {
	return nil
}
