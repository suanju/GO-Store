package Redis

import (
	"GO-Store/Databases"
	"fmt"
	"github.com/go-redis/redis"
)

var Db *redis.Client
var KeyName keyNameStruct

func init() {
	var err error
	var redisConfig = Databases.RConfig

	//初始化redis 键名
	initRedisKey()
	// 创建链接
	Db = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConfig.IP, redisConfig.Port),
		Password: redisConfig.Password, // no password set
		DB:       0,                    // use default DB
	})
	_, err = Db.Ping().Result()
	if err != nil {
		fmt.Printf("redis连接失败:%v \n", err)
	}

}
