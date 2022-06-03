package requests

import (
	"github.com/fans1992/jiaoma/app/requests/validators"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type LoginByPhoneRequest struct {
	Mobile string `json:"mobile,omitempty" valid:"mobile"`
	Code   string `json:"code,omitempty" valid:"code"`
}

// LoginByPhone 验证表单，返回长度等于零即通过
func LoginByPhone(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"mobile": []string{"required", "digits:11"},
		"code":   []string{"required", "digits:6"},
	}
	messages := govalidator.MapData{
		"mobile": []string{
			"required:手机号为必填项，参数名称 phone",
			"digits:手机号长度必须为 11 位的数字",
		},
		"code": []string{
			"required:验证码答案必填",
			"digits:验证码长度必须为 6 位的数字",
		},
	}

	errs := validate(data, rules, messages)

	// 手机验证码
	_data := data.(*LoginByPhoneRequest)
	errs = validators.ValidateVerifyCode(_data.Mobile, _data.Code, errs)

	return errs
}

type LoginByPasswordRequest struct {
	Mobile   string `valid:"mobile" json:"mobile"`
	Password string `valid:"password" json:"password,omitempty"`
}

// LoginByPassword 验证表单，返回长度等于零即通过
func LoginByPassword(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"mobile":   []string{"required", "min:3"},
		"password": []string{"required", "min:6"},
	}
	messages := govalidator.MapData{
		"mobile": []string{
			"required:登录 ID 为必填项，支持手机号、邮箱和用户名",
			"min:登录 ID 长度需大于 3",
		},
		"password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
	}

	errs := validate(data, rules, messages)

	return errs
}
