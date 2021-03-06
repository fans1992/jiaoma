package requests

import (
	"github.com/fans1992/jiaoma/app/requests/validators"
	"github.com/fans1992/jiaoma/pkg/auth"
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type UserUpdateProfileRequest struct {
	NickName string `valid:"nick_name" json:"nick_name"`
	Sex      string `valid:"sex" json:"sex"`
	Address  string `valid:"address" json:"address"`
}

func UserUpdateProfile(data interface{}, c *gin.Context) map[string][]string {

	// 查询用户名重复时，过滤掉当前用户 ID
	rules := govalidator.MapData{
		"nick_name": []string{"required", "between:3,20"},
		"sex":       []string{"min_cn:1", "max_cn:4", "in:男,女"},
		"address":   []string{"min_cn:2", "max_cn:40"},
	}

	messages := govalidator.MapData{
		"nick_name": []string{
			"required:用户名为必填项",
			"alpha_num:用户名格式错误，只允许数字和英文",
			"between:用户名长度需在 3~20 之间",
			"not_exists:用户名已被占用",
		},
		"sex": []string{
			"min_cn:描述长度需至少 1 个字",
			"max_cn:描述长度不能超过 4 个字",
			"in:性别为 男 或者 女",
		},
		"address": []string{
			"min_cn:城市需至少 2 个字",
			"max_cn:城市不能超过 40 个字",
		},
	}
	return validate(data, rules, messages)
}

type UserUpdateEmailRequest struct {
	NewEmail         string `json:"new_email,omitempty" valid:"new_email"`
	VerificationCode string `json:"verification_code,omitempty" valid:"verification_code"`
}

func UserUpdateEmail(data interface{}, c *gin.Context) map[string][]string {

	currentUser := auth.CurrentUser(c)
	rules := govalidator.MapData{
		"new_email": []string{
			"required",
			"min:4",
			"max:30",
			"email",
			"not_exists:users,email," + currentUser.GetStringID(),
			"not_in:" + *currentUser.Email,
		},
		"verification_code": []string{"required", "digits:6"},
	}
	messages := govalidator.MapData{
		"new_email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
			//"not_exists:Email 已被占用",
			//"not_in:新的 Email 与老 Email 一致",
		},
		"verification_code": []string{
			"required:验证码答案必填",
			"digits:验证码长度必须为 6 位的数字",
		},
	}

	errs := validate(data, rules, messages)
	_data := data.(*UserUpdateEmailRequest)
	errs = validators.ValidateVerifyCode(_data.NewEmail, _data.VerificationCode, errs)

	return errs
}

type UserUpdatePhoneRequest struct {
	Mobile    string `json:"mobile,omitempty" valid:"mobile"`
	Code      string `json:"code,omitempty" valid:"code"`
	NewMobile string `json:"new_mobile,omitempty" valid:"new_mobile"`
}

func UserUpdatePhone(data interface{}, c *gin.Context) map[string][]string {

	currentUser := auth.CurrentUser(c)

	rules := govalidator.MapData{
		"mobile": []string{
			"required",
			"digits:11",
			"in:" + currentUser.Mobile,
		},
		"code": []string{"required", "digits:6"},
		"new_mobile": []string{
			"required",
			"digits:11",
			"not_exists:users,mobile," + currentUser.GetStringID(),
			"not_in:" + currentUser.Mobile,
		},
	}
	messages := govalidator.MapData{
		"mobile": []string{
			"required:手机号为必填项，参数名称 phone",
			"digits:手机号长度必须为 11 位的数字",
			"in:旧手机号不正确",
		},
		"code": []string{
			"required:验证码答案必填",
			"digits:验证码长度必须为 6 位的数字",
		},
		"new_mobile": []string{
			"required:手机号为必填项，参数名称 phone",
			"digits:手机号长度必须为 11 位的数字",
			"not_exists:手机号已被占用",
			"not_in:新的手机与老手机号一致",
		},
	}

	errs := validate(data, rules, messages)
	_data := data.(*UserUpdatePhoneRequest)
	errs = validators.ValidateVerifyCode(_data.NewMobile, _data.Code, errs)

	return errs
}

type UserUpdatePasswordRequest struct {
	Mobile      string `json:"mobile,omitempty" valid:"mobile"`
	Code        string `json:"code,omitempty" valid:"code"`
	NewPassword string `valid:"new_password" json:"new_password,omitempty"`
}

func UserUpdatePassword(data interface{}, c *gin.Context) map[string][]string {
	currentUser := auth.CurrentUser(c)
	rules := govalidator.MapData{
		"mobile": []string{
			"required",
			"digits:11",
			"in:" + currentUser.Mobile,
		},
		"code":         []string{"required", "digits:6"},
		"new_password": []string{"required", "min:6"},
	}
	messages := govalidator.MapData{
		"mobile": []string{
			"required:手机号为必填项，参数名称 phone",
			"digits:手机号长度必须为 11 位的数字",
			"in:旧手机号不正确",
		},
		"code": []string{
			"required:验证码答案必填",
			"digits:验证码长度必须为 6 位的数字",
		},
		"new_password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
	}

	// 确保 comfirm 密码正确
	errs := validate(data, rules, messages)
	_data := data.(*UserUpdatePasswordRequest)
	errs = validators.ValidateVerifyCode(_data.Mobile, _data.Code, errs)

	return errs
}

type UserUpdateAvatarRequest struct {
	Avatar *multipart.FileHeader `valid:"avatar" form:"avatar"`
}

func UserUpdateAvatar(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		// size 的单位为 bytes
		// - 1024 bytes 为 1kb
		// - 1048576 bytes 为 1mb
		// - 20971520 bytes 为 20mb
		"file:avatar": []string{"ext:png,jpg,jpeg", "size:20971520", "required"},
	}
	messages := govalidator.MapData{
		"file:avatar": []string{
			"ext:ext头像只能上传 png, jpg, jpeg 任意一种的图片",
			"size:头像文件最大不能超过 20MB",
			"required:必须上传图片",
		},
	}

	return validateFile(c, data, rules, messages)
}
