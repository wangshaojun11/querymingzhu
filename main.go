package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"uisee.com/queryDataBase-api/common"
)

func main() {
	InitConfig()    // 程序一开始就读取配置
	common.InitDB() // 程序一开始就初始化数据库

	r := gin.Default()
	r = CollectRoute(r)

	port := viper.GetString("server.port")
	if port != "" { // 设置了端口则监听指定端口
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}

func InitConfig() {
	workDir, _ := os.Getwd()                 // 获取当前工作目录
	viper.SetConfigName("application")       // 设置读取的配置文件名
	viper.SetConfigType("yml")               // 设置读取配置文件类型
	viper.AddConfigPath(workDir + "/config") // 设置配置文件路径
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
