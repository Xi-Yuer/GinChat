package service

import (
	models "GinChat/models/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserList(ctx *gin.Context) {

	data := models.GetUserList()

	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": data,
	})
}
