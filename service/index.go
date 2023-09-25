package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetIndex(ctx *gin.Context) {
	
	ctx.JSON(http.StatusOK, gin.H{
		"message": "welcome",
	})
}
