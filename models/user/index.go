package models

import (
	"GinChat/utils"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string `json:"name" form:"name"`
	Password      string `json:"password" form:"password"`
	Phone         string `json:"phone" form:"phone"`
	Email         string `json:"email" form:"email"`
	ClientIP      string `json:"clientIP" form:"clientIP"`
	Identity      string `json:"identity" form:"identity"`
	ClientPort    string `json:"clientPort" form:"clientPort"`
	LoginTime     string `json:"loginTime" form:"loginTime"`
	LoginOutTime  string `json:"loginOutTime" form:"loginOutTime"`
	HeartbeatTime string `json:"heartbeatTime" form:"heartbeatTime"`
	IsLogout      bool   `json:"isLogout" form:"isLogout"`
	DeviceInfo    string `json:"deviceInfo" form:"deviceInfo"`
}

// 表名
func (tabel *UserBasic) TableName() string {
	return "user_basic"
}

// CRUD
func CreateUser(u *UserBasic) error {
	u.Password = utils.Md5Encode(u.Password)
	return utils.DB.Create(u).Error
}

func GetUserList(limit, page int) []*UserBasic {
	userList := make([]*UserBasic, limit)
	utils.DB.Limit(limit).Offset((page - 1) * limit).Find(&userList)
	return userList
}

func DeleteUser(id int) {
	utils.DB.Delete(&UserBasic{}, id)
}

func UpdateUser(id string, u *UserBasic) *gorm.DB {
	return utils.DB.Model(u).Where("id = ?", id).Updates(u)
}

// 通过条件查询用户
func FindUserByName(name string) *UserBasic {
	user := &UserBasic{}
	utils.DB.Where("name = ?", name).First(user)
	return user
}
func FindUserByPhone(phone string) *UserBasic {
	user := &UserBasic{}
	utils.DB.Where("phone = ?", phone).First(user)
	return user
}
func FindUserByEmail(email string) *UserBasic {
	user := &UserBasic{}
	utils.DB.Where("email = ?", email).First(user)
	return user
}

func UserLogin(phone, password string) bool {
	hasExist := FindUserByPhone(phone)
	if hasExist.Phone == "" {
		return false
	}

	if hasExist.Password != password {
		return false
	}

	return true
}
