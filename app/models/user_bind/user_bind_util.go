package user_bind

import (
	"github.com/fans1992/jiaoma/pkg/app"
	"github.com/fans1992/jiaoma/pkg/database"
	"github.com/fans1992/jiaoma/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (userBind UserBind) {
	database.DB.Where("id", idstr).First(&userBind)
	return
}

func GetBy(field, value string) (userBind UserBind) {
	database.DB.Where("? = ?", field, value).First(&userBind)
	return
}

func All() (userBinds []UserBind) {
	database.DB.Find(&userBinds)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(UserBind{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

// IsWechatUser 判断是否是公众号用户
func IsWechatUser(userID uint64) bool {
	var count int64
	database.DB.Model(UserBind{}).Where("user_id = ?", userID).Where("type = ?", "official_account").Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (userBinds []UserBind, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(UserBind{}),
		&userBinds,
		app.V1URL(database.TableName(&UserBind{})),
		perPage,
	)
	return
}
