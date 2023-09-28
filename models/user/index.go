package models

import (
	"GinChat/utils"
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string    `json:"name" form:"name" binding:"required"`
	Password      string    `json:"password" form:"password" binding:"required"`
	Phone         string    `json:"phone" form:"phone" binding:"required"`
	Email         string    `json:"email" form:"email" binding:"required"`
	ClientIP      string    `json:"clientIP" form:"clientIP" binding:"required"`
	Identity      string    `json:"identity" form:"identity" binding:"required"`
	ClientPort    string    `json:"clientPort" form:"clientPort" binding:"required"`
	LoginTime     time.Time `json:"loginTime" form:"loginTime"`
	LoginOutTime  time.Time `json:"loginOutTime" form:"loginOutTime"`
	HeartbeatTime time.Time `json:"heartbeatTime" form:"heartbeatTime"`
	IsLogout      bool      `json:"isLogout" form:"isLogout"`
	DeviceInfo    string    `json:"deviceInfo" form:"deviceInfo"`
}

// 表名
func (tabel *UserBasic) TableName() string {
	return "user_basic"
}

// CRUD
func (*UserBasic) CreateUser(u *UserBasic) error {
	return utils.DB.Create(&u).Error
}

func (*UserBasic) GetUserList(limit, page int) []*UserBasic {
	userList := make([]*UserBasic, limit)
	utils.DB.Limit(limit).Offset((page - 1) * limit).Find(&userList)
	return userList
}
