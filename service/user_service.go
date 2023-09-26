package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": make([]int, 1),
	})
}
