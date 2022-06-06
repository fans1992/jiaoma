package v1

import (
	"github.com/fans1992/jiaoma/app/models/contact"
	"github.com/fans1992/jiaoma/app/policies"
	"github.com/fans1992/jiaoma/app/requests"
	"github.com/fans1992/jiaoma/pkg/auth"
	"github.com/fans1992/jiaoma/pkg/response"

	"github.com/gin-gonic/gin"
)

type ContactsController struct {
	BaseAPIController
}

func (ctrl *ContactsController) Index(c *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	data, pager := contact.Paginate(c, 10)
	response.JSON(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}

func (ctrl *ContactsController) Show(c *gin.Context) {
	contactModel := contact.Get(c.Param("id"))
	if contactModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, contactModel)
}

func (ctrl *ContactsController) Store(c *gin.Context) {

	request := requests.ContactRequest{}
	if ok := requests.Validate(c, &request, requests.ContactSave); !ok {
		return
	}

	contactModel := contact.Contact{
		UserID:       auth.CurrentUID(c),
		AcceptName:   request.AcceptName,
		Mobile:       request.Mobile,
		ContactEmail: request.ContactEmail,
		Address:      request.Address,
	}
	contactModel.Create()
	if contactModel.ID > 0 {
		response.Created(c, contactModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *ContactsController) Update(c *gin.Context) {

	contactModel := contact.Get(c.Param("id"))
	if contactModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyContact(c, contactModel); !ok {
		response.Abort403(c)
		return
	}

	request := requests.ContactRequest{}
	if ok := requests.Validate(c, &request, requests.ContactSave); !ok {
		return
	}

	contactModel.AcceptName = request.AcceptName
	contactModel.Mobile = request.Mobile
	contactModel.ContactEmail = request.ContactEmail
	contactModel.Address = request.Address
	contactModel.IsDefault = request.IsDefault

	rowsAffected := contactModel.Save()
	if rowsAffected > 0 {
		response.Data(c, contactModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *ContactsController) Delete(c *gin.Context) {

	contactModel := contact.Get(c.Param("id"))
	if contactModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyContact(c, contactModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := contactModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}
