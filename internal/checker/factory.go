package checker

import "github.com/gin-gonic/gin"

const (
	BASE_CHECKER               = "base_checker"
	REGISTER_FORBIDDEN_CHECKER = "register_forbidden_checker"
	SMS_CODE_CHECKER           = "sms_code_checker"
	SMS_CODE_KEY_CHECKER       = "sms_code_key_checker"
)

type Checker interface {
	DoCheck(ctx *gin.Context) error
}

func GetChecker(ctx *gin.Context, checkerName string) Checker {
	switch checkerName {
	case BASE_CHECKER:
		return &BaseChecker{}
	case REGISTER_FORBIDDEN_CHECKER:
		return &RegisterForbiddenChecker{}
	case SMS_CODE_CHECKER:
		return &SmsCodeChecker{}
	case SMS_CODE_KEY_CHECKER:
		return &SmsCodeKeyChecker{}
	}
	return &BaseChecker{}
}
