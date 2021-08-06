package controller

import (
	"github.com/gin-gonic/gin"
	"go-web-edu/model"
	"go-web-edu/response"
	"go-web-edu/service"
	"strconv"
)

// 获取文章列表
func GetArticleList(ctx *gin.Context) {
	title := ctx.Request.FormValue("title")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))        // 页数
	per_num, _ := strconv.Atoi(ctx.DefaultQuery("per_num", "10")) // 每页个数

	category_id, _ := strconv.Atoi(ctx.Query("category_id")) // 分类ID
	order := ctx.DefaultQuery("order", "id")
	orderType := ctx.DefaultQuery("order_type", "DESC")

	where := make(map[string]interface{})
	if title != "" {
		where["title"] = title
	}
	if category_id != 0 {
		where["category_id"] = category_id
	}
	where["offset"] = (page - 1) * per_num
	where["limit"] = per_num
	where["order"] = order
	where["order_type"] = orderType


	list, err := service.GetArticleList(where)
	if err != nil {
		response.Error(ctx, nil, "获取失败")
	}
	response.Success(ctx, gin.H{"data": list}, "success")

}

// 根据ID获取文章
func GetArticle(ctx *gin.Context) {
	Id, _ := strconv.Atoi(ctx.Request.FormValue("id"))
	article, err := service.GetArticleById(Id)

	if err != nil {
		response.Error(ctx, gin.H{"articleId": Id}, "没有该文章 ")
		return
	}

	response.Success(ctx, gin.H{"data": article}, "获取成功")
}

// 创建文章
func CreateArticle(ctx *gin.Context) {
	title := ctx.DefaultPostForm("title", "default")
	content := ctx.DefaultPostForm("content", "2233")

	article := model.Article{Title: title, Content: content}
	err := service.CreateArticle(&article)
	if err != nil {
		response.Error(ctx, nil, "创建失败")
	} else {
		response.Success(ctx, gin.H{"article": article}, "创建成功")
	}

}

// 更新文章
func UpdateArticle(ctx *gin.Context) {
	Id, _ := strconv.Atoi(ctx.Request.FormValue("id"))
	article, err := service.GetArticleById(Id)

	if err != nil {
		response.Error(ctx, gin.H{"article": article}, "没有该文章 ")
		return
	}
	title := ctx.Request.FormValue("title")

	if title == "" {
		response.Error(ctx, nil, "title 不能为空")
		return
	}
	content := ctx.DefaultPostForm("content", "2233")
	categoryId := ctx.DefaultPostForm("category_id", "1")

	cate_id, _ := strconv.Atoi(categoryId)

	article.Title = title
	article.Content = content
	article.CategoryId = uint(cate_id)

	err = service.UpdateArticle(article)

	if err != nil {
		response.Error(ctx, nil, "更新失败")
	} else {
		response.Success(ctx, gin.H{"category": article}, "更新成功")
	}
}

// 删除文章
func DeleteArticle(ctx *gin.Context) {
	Id, _ := strconv.Atoi(ctx.Request.FormValue("id"))
	article, err := service.GetArticleById(Id)
	if err != nil {
		response.Error(ctx, gin.H{"article": article}, "没有该文章 ")
		return
	}
	re := service.DeleteArticle(article, Id)
	if re {
		response.Success(ctx, nil, "删除成功")
	} else {
		response.Error(ctx, nil, "删除失败")
	}
}
