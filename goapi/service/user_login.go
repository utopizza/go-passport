package service

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	accountService "github.com/utopizza/go-passport/account/service"
	"github.com/utopizza/go-passport/consts"
	sessionService "github.com/utopizza/go-passport/session/service"
)

func UserLogin(ctx *gin.Context) {
	loginName := ctx.PostForm("login_name")
	password := ctx.PostForm("password")

	// 参数检查
	if loginName == "" || password == "" {
		log.Println("login_name or password is empty")
		ctx.JSON(http.StatusOK, gin.H{
			"code": consts.CodeInvalidParams,
			"data": "login_name and password can not be empty",
		})
		return
	}

	// 根据登录名，查帐号信息
	accountInfo, err := accountService.ReadAccountInfoByName(ctx, loginName)
	if err != nil {
		log.Printf("failed to read account, err:%+v", err)
		ctx.JSON(http.StatusOK, gin.H{
			"code": consts.CodeServerInnerError,
			"data": "login_name and password can not be empty",
		})
		return
	}

	// 帐号不存在
	if accountInfo == nil {
		log.Printf("account not found for login_name:%+v", loginName)
		ctx.JSON(http.StatusOK, gin.H{
			"code": consts.CodeAccountOrPasswordIncorrect,
			"data": "account or password incorrect",
		})
		return
	}

	// 密码不对
	if accountInfo.Password != password {
		log.Printf("password not correct for login_name:%+v", loginName)
		ctx.JSON(http.StatusOK, gin.H{
			"code": consts.CodeAccountOrPasswordIncorrect,
			"data": "account or password incorrect",
		})
		return
	}

	// 创建session数据（是否需要清除旧的session？让其自然过期）
	sessionData, err := sessionService.CreateSession(ctx, accountInfo.UserId)
	if err != nil {
		log.Printf("failed to create session, err:%+v", err)
		ctx.JSON(http.StatusOK, gin.H{
			"code": consts.CodeServerInnerError,
			"data": "failed to create session",
		})
		return
	}

	//种session_id到cookie
	ctx.SetCookie("session_id", sessionData.Key, 3600*24*30, "", ctx.Request.Host, false, false)

	// 返回登录成功
	ctx.JSON(http.StatusOK, gin.H{
		"code": consts.CodeSuccess,
		"data": accountInfo,
	})
}

func UserLogout(ctx *gin.Context) {
	// 从cookie中获取登录态
	sessionKey, err := ctx.Cookie("session_id")
	if err != nil {
		log.Printf("failed to read session_id from cookie, err:%+v", err)
		ctx.JSON(http.StatusOK, gin.H{
			"code": consts.CodeServerInnerError,
			"data": "failed to read session_id",
		})
		return
	}

	// 查session
	sessionData, err := sessionService.ReadSession(ctx, sessionKey)
	if err != nil {
		log.Printf("failed to read session from redis, err:%+v", err)
		ctx.JSON(http.StatusOK, gin.H{
			"code": consts.CodeServerInnerError,
			"data": "failed to read session data",
		})
		return
	}

	if sessionData == nil {
		log.Printf("no session founded, sessionKey:%+v", sessionKey)
		ctx.JSON(http.StatusOK, gin.H{
			"code": consts.CodeServerInnerError,
			"data": "no session founded",
		})
		return
	}

	// 注销登录态
	if err := sessionService.DeleteSession(ctx, sessionKey); err != nil {
		log.Printf("failed to delete session, error:%+v", err)
		ctx.JSON(http.StatusOK, gin.H{
			"code": consts.CodeServerInnerError,
			"data": "failed to delete session",
		})
		return
	}

	// 返回登出成功
	ctx.JSON(http.StatusOK, gin.H{
		"code": consts.CodeSuccess,
		"data": "logout success",
	})
}
