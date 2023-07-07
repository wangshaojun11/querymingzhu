package controller

import (
	"github.com/gin-gonic/gin"
	"uisee.com/queryDataBase-api/common"
	"uisee.com/queryDataBase-api/models"
)

func Query(c *gin.Context) {
	DB := common.GetDB()

	namespace := c.PostForm("namespace")
	startdate := c.PostForm("starddate")
	enddate := c.PostForm("enddate")

	// 通过输入的内容查询数据库
	var results []models.AlertLabel
	DB.Raw(`SELECT DISTINCT al.Value, ag.groupKey
	FROM AlertLabel al
	JOIN Alert a ON a.ID = al.AlertID
	JOIN AlertGroup ag ON ag.ID = a.alertGroupID
	WHERE al.Label = 'container'
	  AND a.ID IN (
		SELECT a.ID
		FROM AlertLabel al
		JOIN Alert a ON a.ID = al.AlertID
		JOIN AlertGroup ag ON ag.ID = a.alertGroupID
		WHERE al.Label = 'namespace' AND al.Value = ?
		  AND ag.time >= ? AND ag.time < ?
	  );`, namespace, startdate, enddate).Scan(&results)

	c.HTML(200, "queryinfo.html", gin.H{
		"results": results,
	})
	// 跳转到/
	// c.Redirect(301, "/query/info")

	// 返回数据
	// fmt.Println(results)
	// c.JSON(200, gin.H{
	// 	"code": 200,
	// 	"data": results,
	// 	"msg":  "查询成功",
	// })
}
