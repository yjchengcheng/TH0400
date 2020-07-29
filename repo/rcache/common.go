package rcache

import (
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
		opt, err := redis.ParseURL("redis://localhost:6479/1") // todo 配置文件
		if err != nil {
			panic(err)
		}

		rdb = redis.NewClient(opt)

		ctx, f := utils.GetTimeoutCtx()
		defer f()
		if rdb.Ping(ctx).Err() != nil {
			log.Fatalf("ping redis failed: %s", err.Error())
		}
	}

}

// MAC marshal and cache 序列化后放入缓存
func MAC(key string, value interface{}) error {
	data, err := utils.MRE(value)
	if err != nil {
		return err
	}

	ctx, r := utils.GetTimeoutCtx()
	defer r()

	if _, err := rdb.Set(ctx, key, data, utils.GetCacheTimeout()).Result(); err != nil {
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
