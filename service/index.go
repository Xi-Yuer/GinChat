package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// GetIndex godoc
// @Summary ping GetIndex
// @Schemes
// @Description do ping
// @Tags 首页
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /index [get]
func GetIndex(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"message": "welcome",
	})
}


