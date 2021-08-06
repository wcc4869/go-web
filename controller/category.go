package controller

import (
	"github.com/gin-gonic/gin"
	"go-web-edu/model"
	"go-web-edu/response"
	"go-web-edu/service"
	"strconv"
)

// 分类列表
func GetCategoryList(ctx *gin.Context) {
	name := ctx.Request.FormValue("name")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))        // 页数
	per_num, _ := strconv.Atoi(ctx.DefaultQuery("per_num", "10")) // 每页个数

	order := ctx.DefaultQuery("order", "id")
	orderType := ctx.DefaultQuery("order_type", "DESC")

	where := make(map[string]interface{})
	if name != "" {
		where["name"] = name
	}
	where["offset"] = (page - 1) * per_num
	where["limit"] = per_num
	where["order"] = order
	where["order_type"] = orderType

	list, err := service.GetCategoryList(where)
	if err != nil {
		response.Error(ctx, nil, "获取失败")
	}
	response.Success(ctx, gin.H{"data": list}, "success")

}
//  根据ID获取分类
func GetCategory(ctx *gin.Context) {
	categoryId, _ := strconv.Atoi(ctx.Request.FormValue("id"))
	category, err := service.GetCategoryById(categoryId)

	if err != nil {
		response.Error(ctx, gin.H{"category_id": categoryId}, "没有该分类 ")
		return
	}

	response.Success(ctx, gin.H{"data": category}, "获取成功")
}

// 创建分类
func CreateCategory(ctx *gin.Context) {
	name := ctx.DefaultPostForm("name", "default")

	sort, _ := strconv.Atoi(ctx.DefaultPostForm("sort", "1"))
	category := model.Category{Name: name, Sort: sort}
	err := service.CreateCategory(&category)
	if err != nil {
		response.Error(ctx, nil, "创建失败")
	} else {
		response.Success(ctx, gin.H{"category": category}, "创建成功")
	}

}

// 更新分类
func UpdateCategory(ctx *gin.Context) {
	categoryId, _ := strconv.Atoi(ctx.Request.FormValue("id"))
	category, err := service.GetCategoryById(categoryId)

	if err != nil {
		response.Error(ctx, gin.H{"category_id": categoryId}, "没有该分类 ")
		return
	}
	name := ctx.Request.FormValue("name")

	if name == "" {
		response.Error(ctx, nil, "name 不能为空")
		return
	}

	category.Name = name

	//ctx.ShouldBind(&category)
	err = service.UpdateCategory(category)

	if err != nil {
		response.Error(ctx, nil, "更新失败")
	} else {
		response.Success(ctx, gin.H{"category": category}, "更新成功")
	}
}

// 删除分类
func DeleteCategory(ctx *gin.Context) {
	categoryId, _ := strconv.Atoi(ctx.Request.FormValue("id"))
	category, err := service.GetCategoryById(categoryId)

	if err != nil {
		response.Error(ctx, gin.H{"category_id": categoryId}, "没有该分类 ")
		return
	}
	re := service.DeleteCategory(category, categoryId)
	if re {
		response.Success(ctx, nil, "删除成功")
	} else {
		response.Error(ctx, nil, "删除失败")
	}
}
