package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5 MD5加密算法 v 是要加密的值
func MD5(v string) string {
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}
