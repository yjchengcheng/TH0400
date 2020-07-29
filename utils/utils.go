package utils

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"time"
)

// GetCacheTimeout 获取缓存过期时间
func GetCacheTimeout() time.Duration {
	// todo
	// 防止缓存雪崩 缓存过期时间应该有一定的随机性

	return time.Hour
}

// GetTimeoutCtx 读取配置文件的超时配置并返回context
func GetTimeoutCtx() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Second*10) // todo 读取配置文件
}

// MRE marshal or error 序列化或者返回错误
func MRE(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// EncodePassword 密码加密
func EncodePassword(password string) string {
	key := []byte("secret")
	h := hmac.New(sha256.New, key)
	h.Write([]byte(password))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))

}
