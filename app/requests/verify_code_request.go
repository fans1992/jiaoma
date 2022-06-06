package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type VerifyCodePhoneRequest struct {
	Mobile string `json:"mobile,omitempty" valid:"mobile"`
}

// VerifyCodePhone 验证表单，返回长度等于零即通过
func VerifyCodePhone(data interface{}, c *gin.Context) map[string][]string {
	// 1. 定制认证规则
	rules := govalidator.MapData{
		"mobile": []string{"required", "digits:11"},
	}

	// 2. 定制错误消息
	messages := govalidator.MapData{
		"mobile": []string{
			"required:手机号为必填项，参数名称 mobile",
			"digits:手机号长度必须为 11 位的数字",
		},
	}

	errs := validate(data, rules, messages)

	return errs
}

type VerifyCodeEmailRequest struct {
	NewEmail string `json:"new_email,omitempty" valid:"new_email"`
}

// VerifyCodeEmail 验证表单，返回长度等于零即通过
func VerifyCodeEmail(data interface{}, c *gin.Context) map[string][]string {

	// 1. 定制认证规则
	rules := govalidator.MapData{
		"new_email": []string{"required", "min:4", "max:30", "email"},
	}

	// 2. 定制错误消息
	messages := govalidator.MapData{
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
		},
	}

	errs := validate(data, rules, messages)

	return errs
}
