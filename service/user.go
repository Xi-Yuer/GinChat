package service

import (
	models "GinChat/models/user"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
	data := models.GetUserList(params.Limit, params.Page)
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": data,
	})
}

func CreateUser(ctx *gin.Context) {
	user := &models.UserBasic{}
	err := ctx.ShouldBind(user)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"code": http.StatusBadRequest,
				"msg":  err.Error(),
			})
		return
	}
	error := models.CreateUser(user)
	if error != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"code": http.StatusBadRequest,
				"msg":  error.Error(),
			})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"msg":  "创建成功",
	})
}

func DeleteUser(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, _ := strconv.Atoi(idStr)
	models.DeleteUser(id)
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "删除成功",
	})
}

func UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	user := &models.UserBasic{}
	err := ctx.ShouldBind(user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"msg": err.Error(),
		})
	}
	fmt.Println(id)
	models.UpdateUser(id, user)
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "更新成功",
	})
}
