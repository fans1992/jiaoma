//Package user_bind 模型
package user_bind

import (
	"github.com/fans1992/jiaoma/app/models"
	"github.com/fans1992/jiaoma/app/models/user"
	"github.com/fans1992/jiaoma/pkg/database"
)

type UserBind struct {
	models.BaseModel

	UserID      string `json:"user_id,omitempty"`
	Type        string `json:"type,omitempty"`
	AppID       string `json:"app_id,omitempty"`
	OpenID      string `json:"open_id,omitempty"`
	NickName    string `json:"nick_name,omitempty"`
	Sex         string `json:"sex,omitempty"`
	Email       string `json:"email,omitempty"`
	Avatar      string `json:"avatar,omitempty"`
	City        string `json:"city,omitempty"`
	Province    string `json:"province,omitempty"`
	Country     string `json:"country,omitempty"`
	Language    string `json:"language,omitempty"`
	Subscribe   string `json:"subscribe,omitempty"`
	SubscribeAt string `json:"subscribe_at,omitempty"`
	Unionid     string `json:"union,omitempty"`

	// 通过 user_id 关联用户
	User user.User `json:"user"`

	models.CommonTimestampsField
}

func (UserBind) TableName() string {
	// Put table name in here
	return "ibrand_user_bind"
}

func (userBind *UserBind) Create() {
	database.DB.Create(&userBind)
}

func (userBind *UserBind) Save() (rowsAffected int64) {
	result := database.DB.Save(&userBind)
	return result.RowsAffected
}

func (userBind *UserBind) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&userBind)
	return result.RowsAffected
}
