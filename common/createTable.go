package common

import (
	models "GinChat/models/user"
	"GinChat/utils"
)

func CreateTables() {
	utils.DB.AutoMigrate(&models.UserBasic{})
}
