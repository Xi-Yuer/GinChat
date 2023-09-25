package models

import (
	"GinChat/utils"
	"fmt"
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
func GetUserList() []*UserBasic {
	db, err := utils.InitMySQL()
	db.AutoMigrate(&UserBasic{})
	if err != nil {
		return nil
	}
	data := make([]*UserBasic, 10)
	db.Find(&data)
	for _, v := range data {
		fmt.Println(v)
	}
	return data
}
