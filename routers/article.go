package routers

import (
	"github.com/gin-gonic/gin"
	"go-web-edu/controller"
)

func LoadArticle(r *gin.Engine) {
	articleRoute := r.Group("/article")
	articleRoute.GET("/list", controller.GetArticleList)
	articleRoute.GET("/detail", controller.GetArticle)
	articleRoute.POST("/create", controller.CreateArticle)
	articleRoute.PUT("/update", controller.UpdateArticle)
	articleRoute.DELETE("/article", controller.DeleteArticle)
}
