package models

import (
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string ``
	Password      string
	Phone         string
	Email         string
	ClientIP      string
	Identity      string
	ClientPort    string
	LoginTime     time.Time
	LoginOutTime  time.Time
	HeartbeatTime time.Time
	IsLogout      bool
	DeviceInfo    string
}

// 表名
func (tabel *UserBasic) TableName() string {
	return "user_basic"
}

// CRUD
func GetUserList() {
}
