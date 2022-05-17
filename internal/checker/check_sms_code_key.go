package checker

import "github.com/gin-gonic/gin"

type SmsCodeChecker struct{}

func (c *SmsCodeChecker) DoCheck(ctx *gin.Context) error {
	return nil
}
