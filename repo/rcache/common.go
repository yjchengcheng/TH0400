package rcache

import (
	"TH0400/logger"
	"TH0400/utils"
	"log"
	"sync"

	"github.com/go-redis/redis/v8"
)

var localCache *sync.Map
var rdb *redis.Client // 全局redis缓存客户端

// 初始化redis链接
func init() {
	if localCache == nil {
		localCache = new(sync.Map)
	}

	// ------redis---------
	if rdb == nil {
		opt, err := redis.ParseURL(utils.GetRedisURI())
		if err != nil {
			panic(err)
		}

		rdb = redis.NewClient(opt)

		ctx, f := utils.GetRTimeoutCtx()
		defer f()

		if rdb.Ping(ctx).Err() != nil {
			log.Fatalf("ping redis failed: %s", err.Error())
		}
	}

}

// MAP marshal and put 序列化后放入缓存
func MAP(key string, value interface{}) error {
	data, err := utils.MRE(value)
	if err != nil {
		logger.Errorf("redis; key: [%s], value: [xxx]; marshal error; %s", key, err.Error())
		return err
	}

	ctx, r := utils.GetRTimeoutCtx()
	defer r()

	if _, err := rdb.Set(ctx, key, data, utils.GetRedisExpiration()).Result(); err != nil {
		logger.Errorf("redis; key: [%s], value: [xxx]; set error; %s", key, err.Error())
		return err
	}

	return nil
}

// GAU get and unmarshal 从缓存中获取并反序列化
func GAU(key string, value interface{}) error {

	ctx, r := utils.GetRTimeoutCtx()
	defer r()

	data, err := rdb.Get(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			logger.Infof("redis; key: [%s] ,value: [xxx]; not exist", key)
			return err
		}
		logger.Errorf("redis; key: [%s], value: [xxx]; get error; %s", key, err.Error())
		return err
	}

	if err := utils.URE([]byte(data), value); err != nil {
		logger.Errorf("redis; key: [%s], value: [xxx]; unmarshal error; %s", key, err.Error())
		return err
	}

	return nil
}

// PutToken 缓存token
func PutToken(token string, userid int) {
	localCache.Store(token, userid)
}

// GetToken 获取token缓存
func GetToken(token string) (i int, ok bool) {
	value, ok := localCache.Load(token)
	if ok {
		return value.(int), ok
	}
	return
}
