package data_struct

import (
	"crypto/md5"
	"encoding/hex"
)

func Hash(src string) string {
	res := md5.Sum([]byte(src))
	result := hex.EncodeToString(res[:]) //对应的参数为切片，需要将数组转换为切片。
	return result
}
