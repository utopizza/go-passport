package router

import (
	"github.com/gin-gonic/gin"

	"github.com/utopizza/go-passport/goapi/service"
)

func GetRouter() *gin.Engine {
	router := gin.New()

	// 中间件
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// 手机号+短信验证码 注册登录
	router.POST("/passport/mobile/send_code/", nil)
	router.POST("/passport/mobile/sms_login/", nil)
	router.POST("/passport/mobile/sms_login_only/", nil)
	router.POST("/passport/mobile/sms_login_continue/", nil)

	// 邮箱+邮箱验证码 注册登录
	router.POST("/passport/email/send_code/", nil)
	router.POST("/passport/email/register/", nil)
	router.POST("/passport/email/register_verify/", nil)
	router.POST("/passport/email/register_verify/login/", nil)
	router.POST("/passport/email/code_login/", nil)

	// 三方 注册登录
	router.POST("/passport/auth/login/", nil)
	router.POST("/passport/auth/login_only/", nil)

	// 帐密登录（手机号/用户名/邮箱+密码）
	router.POST("/passport/user/login/", service.UserLogin)
	router.GET("/passport/password/check/", nil)

	// 二维码扫码登录
	router.GET("/passport/web/get_qr_code/", nil)
	router.GET("/passport/web/check_qr_connect/", nil)
	router.POST("/passport/mobile/scan_qr_code/", nil)
	router.POST("/passport/mobile/confirm_qr_code/", nil)

	// 登出
	router.POST("/passport/user/logout", service.UserLogout)

	// 帐号注销
	router.POST("/passport/cancel/check/", nil)
	router.POST("/passport/cancel/post/", nil)
	router.POST("/passport/cancel/do/", nil)

	return router
}
