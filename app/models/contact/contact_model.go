//Package contact 模型
package contact

import (
	"github.com/fans1992/jiaoma/app/models"
	"github.com/fans1992/jiaoma/app/models/user"
	"github.com/fans1992/jiaoma/pkg/database"
)

type Contact struct {
	models.BaseModel

	UserID       string `json:"user_id,omitempty"`
	AcceptName   string `json:"accept_name,omitempty"`
	Mobile       string `json:"mobile,omitempty"`
	ContactEmail string `json:"contact_email,omitempty"`
	Province     *int16 `json:"province,omitempty"`
	City         *int16 `json:"city,omitempty"`
	Area         *int16 `json:"area,omitempty"`
	AddressName  string `json:"address_name,omitempty"`
	Address      string `json:"address,omitempty"`
	IsDefault    int8   `json:"is_default,omitempty"`

	// 通过 user_id 关联用户
	User user.User `json:"user"`

	models.CommonTimestampsField
}

func (Contact) TableName() string {
	return "ibrand_addresses"
}

func (contact *Contact) Create() {
	database.DB.Create(&contact)
}

func (contact *Contact) Save() (rowsAffected int64) {
	result := database.DB.Save(&contact)
	return result.RowsAffected
}

func (contact *Contact) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&contact)
	return result.RowsAffected
}
