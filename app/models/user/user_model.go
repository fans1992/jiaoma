// Package user 存放用户 Model 相关逻辑
package user

import (
	"github.com/fans1992/jiaoma/app/models"
	"github.com/fans1992/jiaoma/pkg/database"
	"github.com/fans1992/jiaoma/pkg/hash"
)

// User 用户模型
type User struct {
	models.BaseModel

	Name            string `json:"name,omitempty"`
	NickName        string `json:"nick_name,omitempty"`
	Email           string `json:"Email,omitempty"`
	Mobile          string `json:"mobile,omitempty"`
	Password        string `json:"-"`
	Status          string `json:"status,omitempty"`
	Sex             string `json:"sex,omitempty"`
	Avatar          string `json:"avatar,omitempty"`
	City            string `json:"city,omitempty"`
	Address         string `json:"address,omitempty"`
	Company         string `json:"company,omitempty"`
	Birthday        string `json:"birthday,omitempty"`
	RememberToken   string `json:"remember_token,omitempty"`
	QQ              string `json:"qq,omitempty"`
	IsWechatManager string `json:"is_wechat_manager,omitempty"`

	models.CommonTimestampsField
}

// Create 创建用户，通过 User.ID 来判断是否创建成功
func (userModel *User) Create() {
	database.DB.Create(&userModel)
}

// ComparePassword 密码是否正确
func (userModel *User) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, userModel.Password)
}

func (userModel *User) Save() (rowsAffected int64) {
	result := database.DB.Save(&userModel)
	return result.RowsAffected
}
