package requests

import (
	"github.com/fans1992/jiaoma/app/requests/validators"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type ResetByPhoneRequest struct {
	Mobile          string `json:"mobile,omitempty" valid:"mobile"`
	Code            string `json:"code,omitempty" valid:"code"`
	NewPassword     string `valid:"new_password" json:"new_password,omitempty"`
	ConfirmPassword string `valid:"confirm_password" json:"confirm_password,omitempty"`
}

// ResetByPhone 验证表单，返回长度等于零即通过
func ResetByPhone(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"mobile":           []string{"required", "digits:11"},
		"code":             []string{"required", "digits:6"},
		"new_password":     []string{"required", "min:6"},
		"confirm_password": []string{"required", "min:6"},
	}
	messages := govalidator.MapData{
		"mobile": []string{
			"required:手机号为必填项，参数名称 phone",
			"digits:手机号长度必须为 11 位的数字",
		},
		"verify_code": []string{
			"required:验证码答案必填",
			"digits:验证码长度必须为 6 位的数字",
		},
		"new_password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
		"confirm_password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
	}

	errs := validate(data, rules, messages)

	_data := data.(*ResetByPhoneRequest)
	errs = validators.ValidatePasswordConfirm(_data.NewPassword, _data.ConfirmPassword, errs)
	errs = validators.ValidateVerifyCode(_data.Mobile, _data.Code, errs)

	return errs
}
