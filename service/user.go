package service

import (
	models "GinChat/models/user"
	"GinChat/utils"
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
	// 参数校验
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
	// 注册逻辑校验
	hasExistName := models.FindUserByName(user.Name)
	if hasExistName.ID != 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "用户已注册",
		})
		return
	}
	// 创建用户
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

func UserLogin(ctx *gin.Context) {
	phone := ctx.PostForm("phone")
	password := ctx.PostForm("password")

	ok := models.UserLogin(phone, password)
	if !ok {
		utils.Error(ctx, "手机号或密码错误")
		return
	}

	utils.Success(ctx, "登陆成功")
}
