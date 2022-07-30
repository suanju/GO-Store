package Redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"gopkg.in/ini.v1"
	"os"
)

var Db *redis.Client
var cfg *ini.File
var KeyName keyNameStruct

type RConfig struct {
	IP       string `ini:"ip"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
}

func init() {
	var err error
	var redisConfig = &RConfig{}

	//初始化redis 键名
	initRedisKey()
	//获取数据库配置

	cfg, err = ini.Load("Config/config.ini")

	err = cfg.Section("redis").MapTo(redisConfig)

	//fmt.Printf("INN 配置文件内容为--%v \n", redisConfig)
	if err != nil {
		fmt.Printf("读取配置文件错误 \n: %v", err)
		os.Exit(1)
	}

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
