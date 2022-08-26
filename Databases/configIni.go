package Databases

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

var cfg *ini.File
var err error
var SqlConfig *SqlConfigStruct
var RConfig *RConfigStruct
var ProjectConfig *ProjectConfigStruct
var ProjectUrl string

type SqlConfigStruct struct {
	IP       string `ini:"ip"`
	Port     int    `ini:"port"`
	User     string `ini:"user"`
	Host     int    `ini:"host"`
	Password string `ini:"password"`
	Database string `ini:"database"`
}

type RConfigStruct struct {
	IP       string `ini:"ip"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
}

type ProjectConfigStruct struct {
	UrlStates bool   `ini:"url_states"`
	Url       string `ini:"url"`
	UrlTest   string `ini:"url_test"`
}

func init() {
	SqlConfig = &SqlConfigStruct{}
	cfg, err = ini.Load("Config/config.ini")
	if err != nil {
		fmt.Printf("配置文件不存在,请检查环境: %v \n", err)
		os.Exit(1)
	}

	err = cfg.Section("mysql").MapTo(SqlConfig)
	if err != nil {
		fmt.Printf("Mysql读取配置文件错误: %v \n", err)
		os.Exit(1)
	}
	RConfig = &RConfigStruct{}
	err = cfg.Section("redis").MapTo(RConfig)
	if err != nil {
		fmt.Printf("Redis读取配置文件错误: %v \n", err)
		os.Exit(1)
	}
	ProjectConfig = &ProjectConfigStruct{}
	err = cfg.Section("project").MapTo(ProjectConfig)
	if err != nil {
		fmt.Printf("Project读取配置文件错误: %v \n", err)
		os.Exit(1)
	}
	//判断是否为正式环境
	if ProjectConfig.UrlStates {
		ProjectUrl = ProjectConfig.Url
	} else {
		ProjectUrl = ProjectConfig.UrlTest
	}
}
