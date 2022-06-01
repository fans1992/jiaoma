package v1

import (
	"github.com/fans1992/jiaoma/app/models/link"
	"github.com/fans1992/jiaoma/pkg/response"

	"github.com/gin-gonic/gin"
)

type LinksController struct {
	BaseAPIController
}

func (ctrl *LinksController) Index(c *gin.Context) {
	response.Data(c, link.AllCached())
}
