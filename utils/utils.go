package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

// MRE marshal or error 序列化或者返回错误
func MRE(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// URE unmarshal or error 反序列化或返回错误
//
// 注意 v 必须是指针类型
func URE(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// EncodePassword 密码加密
func EncodePassword(password string) string {
	key := []byte("secret")
	h := hmac.New(sha256.New, key)
	h.Write([]byte(password))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))

}
