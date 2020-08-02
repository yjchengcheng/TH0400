package utils

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

var (
	mysqlConnTimeout time.Duration
	redisConnTimeout time.Duration
	redisExpiration  time.Duration
)

func init() {

	viper.AutomaticEnv()
	viper.SetEnvPrefix("TH0400")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/home/zhanlan/projects/goProjects/TH0400")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("config file not found")
		} else {
			log.Fatal("failed read config file: ", err.Error())
		}
	}

	mysqlConnTimeout = time.Second * time.Duration(mustGetInt("db.conntimeout"))
	redisConnTimeout = time.Second * time.Duration(mustGetInt("redis.conntimeout"))
	redisExpiration = time.Second * time.Duration(mustGetInt("redis.expiration"))
}

// GetSQLURI 获取数据库链接地址
func GetSQLURI() (uri string) {
	return mustGetString("db.uri")
}

// GetRedisURI 获取redis链接地址
func GetRedisURI() (uri string) {
	return mustGetString("redis.uri")
}

// GetRedisTimeout 获取redis链接的超时时间
func GetRedisTimeout() (timeout time.Duration) {
	return redisConnTimeout
}

// GetRedisExpiration 获取redis缓存过期时间
func GetRedisExpiration() (timeout time.Duration) {
	return redisExpiration
}

// GetMysqlTimeout 获取Mysql链接的超时时间
func GetMysqlTimeout() (timeout time.Duration) {
	return mysqlConnTimeout
}

func mustGetString(key string) (value string) {
	value = viper.GetString(key)
	if value == "" {
		log.Fatalf("config key:[%s]  --> value:[xxx] must set", key)
	}
	return
}

func mustGetInt(key string) (value int) {
	value = viper.GetInt(key)
	if value == 0 {
		log.Fatalf("config key:[%s]  --> value:[xxx] must set", key)
	}
	return
}
