package Mysql

import (
	"fmt"
	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"math/rand"
	"os"
	"time"
)

var Db *gorm.DB
var cfg *ini.File

type SqlConfig struct {
	IP       string `ini:"ip"`
	Port     int    `ini:"port"`
	User     string `ini:"user"`
	Host     int    `ini:"host"`
	Password string `ini:"password"`
	Database string `ini:"database"`
}

type RedisConfig struct {
	IP   string `ini:"ip"`
	Port int    `ini:"port"`
}

func init() {
	var err error

	var mysqlConfig = &SqlConfig{}
	//获取数据库配置

	//时间戳设置随机数
	rand.Seed(time.Now().UnixNano())

	cfg, err = ini.Load("Config/config.ini")

	err = cfg.Section("mysql").MapTo(mysqlConfig)

	//fmt.Printf("INN 配置文件内容为--%v \n", mysqlConfig)
	if err != nil {
		fmt.Printf("读取配置文件错误: %v \n", err)
		os.Exit(1)
	}

	// 创建链接
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local", mysqlConfig.User, mysqlConfig.Password, mysqlConfig.IP, mysqlConfig.Port, mysqlConfig.Database)
	//dsn := "root:root@tcp(127.0.0.1:3306)/go_study?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("数据库链接错误- %v \n", err)
	}
	if Db.Error != nil {
		fmt.Printf("数据库错误- %v \n", Db.Error)
	}
}
