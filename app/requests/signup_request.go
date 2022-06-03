// Package requests 处理请求数据和表单验证
package requests

import (
	"github.com/fans1992/jiaoma/app/requests/validators"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type SignupPhoneExistRequest struct {
	Phone string `json:"phone,omitempty" valid:"phone"`
}

func SignupPhoneExist(data interface{}, c *gin.Context) map[string][]string {
	// 自定义验证规则
	rules := govalidator.MapData{
		"phone": []string{"required", "digits:11"},
	}
	// 自定义验证出错时的提示
	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项，参数名称 phone",
			"digits:手机号长度必须为 11 位的数字",
		},
	}
	return validate(data, rules, messages)
}

// SignupUsingPhoneRequest 通过手机注册的请求信息
type SignupUsingPhoneRequest struct {
	Mobile   string `json:"mobile,omitempty" valid:"mobile"`
	Code     string `json:"code,omitempty" valid:"code"`
	Password string `valid:"password" json:"password,omitempty"`
}

func SignupUsingPhone(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"mobile":   []string{"required", "digits:11", "not_exists:ibrand_user,mobile"},
		"password": []string{"required", "min:6"},
		"code":     []string{"required", "digits:6"},
	}

	messages := govalidator.MapData{
		"mobile": []string{
			"required:手机号为必填项，参数名称 mobile",
			"digits:手机号长度必须为 11 位的数字",
		},
		"password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
		"code": []string{
			"required:验证码答案必填",
			"digits:验证码长度必须为 6 位的数字",
		},
	}

	errs := validate(data, rules, messages)

	_data := data.(*SignupUsingPhoneRequest)
	errs = validators.ValidateVerifyCode(_data.Mobile, _data.Code, errs)

	return errs
}
