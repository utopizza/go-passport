package checker

import "github.com/gin-gonic/gin"

type SmsCodeKeyChecker struct{}

func (c *SmsCodeKeyChecker) DoCheck(ctx *gin.Context) error {
	return nil
}
