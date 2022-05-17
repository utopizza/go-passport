package checker

import "github.com/gin-gonic/gin"

type Engine struct{}

func (e *Engine) Execute(ctx *gin.Context, checkerNames ...string) error {
	return nil
}

func (e *Engine) IsChecked(ctx *gin.Context, checkerNames ...string) bool {
	return true
}
