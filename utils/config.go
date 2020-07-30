package utils

import (
	"log"

	"github.com/spf13/viper"
)

// 此文件用于读取和解析配置文件

func init() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("TH0400")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("config file not found")
		} else {
			log.Fatal("failed read config file: ", err.Error())
		}
	}
}

func mustGetString(key string) (value string) {
	value = viper.GetString(key)
	if value == "" {
		log.Fatalf("config key:[%s]  --> value:[xxx] must set", key)
	}
	return
}

// GetSQLURI 获取数据库链接地址
func GetSQLURI() (uri string) {
	return mustGetString("db")
}

// GetRedisURI 获取数据库链接地址
func GetRedisURI() (uri string) {
	return mustGetString("redis")
}
