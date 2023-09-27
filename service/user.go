package service

import (
	models "GinChat/models/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

var user = models.UserBasic{}

type UserQueryParams struct {
	Limit int `form:"limit" binding:"required"`
	Page  int `form:"page"  binding:"required"`
}

func GetUserList(ctx *gin.Context) {
	params := &UserQueryParams{}
	err := ctx.ShouldBind(params)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"code": http.StatusBadRequest,
				"msg":  "参数错误",
			})
		return
	}
	data := user.GetUserList(params.Limit, params.Page)
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": data,
	})
}

func CreateUser(ctx *gin.Context) {
	user := models.UserBasic{}
	err := ctx.ShouldBind(&user)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"code": http.StatusBadRequest,
				"msg":  err.Error(),
			})
		return
	}
	data := user.CreateUser(&user)
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": data,
	})
}
