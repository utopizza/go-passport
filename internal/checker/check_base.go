package checker

import "github.com/gin-gonic/gin"

type BaseChecker struct{}

func (c *BaseChecker) DoCheck(ctx *gin.Context) error {
	return nil
}
