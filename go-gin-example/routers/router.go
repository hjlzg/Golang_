package routers

import (
	"github.com/gin-gonic/gin"
	"go-gin-example/pkg/setting"
	v12 "go-gin-example/routers/api/v1"
)

func InitRouter() *gin.Engine{
	r:=gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	apiV1:=r.Group("/api/v1")
	{
		//获取标签列表
		apiV1.GET("/tags", v12.GetTags)

		//新建标签
		apiV1.POST("/tags", v12.AddTag)

		//更新指定标签
		apiV1.PUT("/tags/:id", v12.EditTag)

		//删除指定标签
		apiV1.DELETE("/tags/:id", v12.DeleteTag)
	}

	return r
}
