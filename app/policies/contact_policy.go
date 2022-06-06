package policies

import (
    "github.com/fans1992/jiaoma/app/models/contact"
    "github.com/fans1992/jiaoma/pkg/auth"

    "github.com/gin-gonic/gin"
)

func CanModifyContact(c *gin.Context, contactModel contact.Contact) bool {
    return auth.CurrentUID(c) == contactModel.UserID
}

// func CanViewContact(c *gin.Context, contactModel contact.Contact) bool {}
// func CanCreateContact(c *gin.Context, contactModel contact.Contact) bool {}
// func CanUpdateContact(c *gin.Context, contactModel contact.Contact) bool {}
// func CanDeleteContact(c *gin.Context, contactModel contact.Contact) bool {}
