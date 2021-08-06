package routers

import (
	"github.com/gin-gonic/gin"
	"go-web-edu/controller"
)

func LoadCategory(r *gin.Engine) {
	categoryRoutes := r.Group("/category")
	categoryRoutes.GET("/list", controller.GetCategoryList)
	categoryRoutes.POST("/create", controller.CreateCategory)
	categoryRoutes.PUT("/update", controller.UpdateCategory)
	categoryRoutes.DELETE("/delete", controller.DeleteCategory)
	categoryRoutes.GET("/get", controller.GetCategory)

	r.GET("/hello", func(c *gin.Context) {
		// c.JSON：返回JSON格式的数据
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})

}
