package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// 小写
func Md5Encode(data string) string {
	md5 := md5.New()
	md5.Write([]byte(data))
	temStr := md5.Sum(nil)
	return hex.EncodeToString(temStr)
}

// 大写
func MD5Encode(data string) string {
	return strings.ToUpper(Md5Encode(data))
}
