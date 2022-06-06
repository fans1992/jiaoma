package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type ContactRequest struct {
	AcceptName   string `valid:"accept_name" json:"accept_name,omitempty"`
	Mobile       string `valid:"mobile" json:"mobile,omitempty"`
	ContactEmail string `valid:"contact_email" json:"contact_email,omitempty"`
	Address      string `valid:"address" json:"address,omitempty"`
}

func ContactSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"accept_name":   []string{"required", "min_cn:2", "max_cn:8"},
		"mobile":        []string{"required", "digits:11"},
		"contact_email": []string{"required", "min:4", "max:30", "email"},
		"address":       []string{"required", "min_cn:3", "max_cn:40"},
	}
	messages := govalidator.MapData{
		"name": []string{
			"required:名称为必填项",
			"min_cn:名称长度需至少 2 个字",
			"max_cn:名称长度不能超过 8 个字",
		},
		"mobile": []string{
			"required:手机号为必填项，参数名称 phone",
			"digits:手机号长度必须为 11 位的数字",
		},
		"contact_email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
		},
		"address": []string{
			"required:帖子标题为必填项",
			"min_cn:标题长度需大于 3",
			"max_cn:标题长度需小于 40",
		},
	}
	return validate(data, rules, messages)
}
