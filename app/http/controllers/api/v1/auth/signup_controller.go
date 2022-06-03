// Package auth 处理用户身份认证相关逻辑
package auth

import (
	v1 "github.com/fans1992/jiaoma/app/http/controllers/api/v1"
	"github.com/fans1992/jiaoma/app/models/user"
	"github.com/fans1992/jiaoma/app/requests"
	"github.com/fans1992/jiaoma/pkg/jwt"
	"github.com/fans1992/jiaoma/pkg/response"

	"github.com/gin-gonic/gin"
)

// SignupController 注册控制器
type SignupController struct {
	v1.BaseAPIController
}

// IsPhoneExist 检测手机号是否被注册
func (sc *SignupController) IsPhoneExist(c *gin.Context) {
	// 获取请求参数，并做表单验证
	request := requests.SignupPhoneExistRequest{}
	if ok := requests.Validate(c, &request, requests.SignupPhoneExist); !ok {
		return
	}

	//  检查数据库并返回响应
	response.Data(c, gin.H{
		"is_new_user": !user.IsPhoneExist(request.Mobile),
	})
}

// SignupUsingPhone 使用手机和验证码进行注册
func (sc *SignupController) SignupUsingPhone(c *gin.Context) {

	// 1. 验证表单
	request := requests.SignupUsingPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.SignupUsingPhone); !ok {
		return
	}

	// 2. 验证成功，创建数据
	userModel := user.User{
		Mobile:   request.Mobile,
		Password: request.Password,
	}
	userModel.Create()

	if userModel.ID > 0 {
		token := jwt.NewJWT().IssueToken(userModel.GetStringID())
		response.Created(c, gin.H{
			"token_type":   "Bearer",
			"access_token": token,
			"wechat_user":  false,
		})
		return
	}

	response.Abort500(c, "创建用户失败，请稍后尝试~")
}

