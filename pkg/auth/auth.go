// Package auth 授权相关逻辑
package auth

import (
	"errors"
	"github.com/fans1992/jiaoma/app/models/user"
	"github.com/fans1992/jiaoma/pkg/logger"

	"github.com/gin-gonic/gin"
)

// Attempt 尝试登录
func Attempt(email string, password string) (user.User, error) {
	userModel := user.GetByMulti(email)
	if userModel.ID == 0 {
		return user.User{}, errors.New("账号不存在")
	}

	if !userModel.ComparePassword(password) {
		return user.User{}, errors.New("密码错误")
	}

	return userModel, nil
}

// LoginByMobile  登录指定用户
func LoginByMobile(mobile string) (user.User, bool, error) {
	userModel := user.GetByMobile(mobile)
	if userModel.ID == 0 {
		// 创建用户
		userModel = user.User{
			Mobile: mobile,
		}
		userModel.Create()
		if userModel.ID > 0 {
			return userModel, true, nil
		} else {
			return user.User{}, true, errors.New("创建用户失败，请稍后尝试~")
		}
	}

	return userModel, false, nil
}

// CurrentUser 从 gin.context 中获取当前登录用户
func CurrentUser(c *gin.Context) user.User {
	userModel, ok := c.MustGet("current_user").(user.User)
	if !ok {
		logger.LogIf(errors.New("无法获取用户"))
		return user.User{}
	}
	// db is now a *DB value
	return userModel
}

// CurrentUID 从 gin.context 中获取当前登录用户 ID
func CurrentUID(c *gin.Context) string {
	return c.GetString("current_user_id")
}
