package main

import (
	"github.com/gin-gonic/gin"
	"go-web-edu/controller"
)

func CollectRoute(r *gin.Engine) *gin.Engine {

	// 分类路由
	categoryRoutes := r.Group("/category")
	categoryRoutes.GET("/list", controller.GetCategoryList)
	categoryRoutes.POST("/create", controller.CreateCategory)
	categoryRoutes.PUT("/update", controller.UpdateCategory)
	categoryRoutes.DELETE("/delete", controller.DeleteCategory)
	categoryRoutes.GET("/get", controller.GetCategory)

	return r
}
