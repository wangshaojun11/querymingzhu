package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"uisee.com/queryDataBase-api/controller"
)

func HtmlHome(e *gin.Context) {
	e.HTML(200, "index.html", nil)
}

func HtmlMingzhu(e *gin.Context) {
	markdownContent := viper.GetString("config.markdown") // 从配置文件中获取 markdown 字段的内容
	e.HTML(200, "mingzhu.html", gin.H{                    // 将 markdownContent 注入到 HTML 模板中
		"MarkdownContent": markdownContent,
	})
}

func HtmlQuery(e *gin.Context) {
	e.HTML(200, "query.html", nil)
}

func HtmlQueryInfo(e *gin.Context) {
	e.HTML(200, "queryinfo.html", nil)
}

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Static("/assets", "./assets") // 加载样式
	r.LoadHTMLGlob("templates/*")   // 加载html文件

	r.POST("/query/databases", controller.Query) // 查询接口
	r.GET("/mingzhu", HtmlMingzhu)               // 明珠工具页面
	r.GET("/query", HtmlQuery)                   // 查询页面
	r.GET("/query/info", HtmlQueryInfo)          // 查询结果页面
	r.GET("/", HtmlHome)                         // 主页
	return r
}
