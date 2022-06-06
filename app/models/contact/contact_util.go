package contact

import (
	"github.com/fans1992/jiaoma/pkg/app"
	"github.com/fans1992/jiaoma/pkg/database"
	"github.com/fans1992/jiaoma/pkg/paginator"
	"gorm.io/gorm/clause"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (contact Contact) {
	database.DB.Preload(clause.Associations).Where("id", idstr).First(&contact)
	return
}

func GetBy(field, value string) (contact Contact) {
	database.DB.Where("? = ?", field, value).First(&contact)
	return
}

func All() (contacts []Contact) {
	database.DB.Find(&contacts)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Contact{}).Where(" = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (contacts []Contact, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(Contact{}),
		&contacts,
		app.V1URL(database.TableName(&Contact{})),
		perPage,
	)
	return
}
