package utils

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 记载配置
func InitConfig() (interface{}, error) {
	// 加载环境配置，这里的 app 需要和 yaml 文件名保持一致
	viper.SetConfigName("app")
	// 文件夹路径
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	// 获取配置信息
	config := viper.Get("mysql")

	return config, err
}

// 加载 MySQL
func InitMySQL() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
