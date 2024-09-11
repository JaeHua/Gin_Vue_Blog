package utils

import (
	"fmt"
	"github.com/go-ini/ini"
)

// 全局变量
var (
	AppMode     string
	HttpPort    string
	JwtKey      string
	Db          string
	DbHost      string
	DbPort      string
	DbUser      string
	DbPassword  string
	DbName      string
	AccessKey   string
	SecreteKey  string
	Bucket      string
	QiniuServer string
	RdHost      string
	RdPort      string
)

// 初始化
func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Printf("读取配置文件出错，请检查路径：", err)

	}
	LoadServe(file)
	LoadDatabase(file)
	LoadQiNiu(file)
	LoadRedis(file)
}

func LoadServe(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("ncsvkjailvhiaubvca")

}
func LoadDatabase(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassword = file.Section("database").Key("DbPassword").MustString("jjh529529")
	DbName = file.Section("database").Key("DbName").MustString("GinBlog")

}
func LoadQiNiu(file *ini.File) {
	AccessKey = file.Section("qiniu").Key("AccessKey").String()
	SecreteKey = file.Section("qiniu").Key("SecreteKey").String()
	Bucket = file.Section("qiniu").Key("Bucket").String()
	QiniuServer = file.Section("qiniu").Key("QiniuServer").String()

}

func LoadRedis(file *ini.File) {
	RdHost = file.Section("redis").Key("RdHost").String()
	RdPort = file.Section("redis").Key("RdPort").String()
}
