package service

import (
	"go-web-edu/model"
)

func GetArticleById(id int) (article *model.Article, err error) {
	article = &model.Article{}
	err = GetDB().First(&article, id).Error

	return article, err
}

func CreateArticle(article *model.Article) (err error) {
	err = GetDB().Create(&article).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateArticle(article *model.Article) (err error) {
	err = GetDB().Save(&article).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteArticle(article *model.Article, id int) bool {
	err := GetDB().Where("id=?", id).Delete(&article).Error
	if err != nil {
		return false
	} else {
		return true
	}
}

func GetArticleList(option map[string]interface{}) (res []model.Article, err error) {

	con := GetDB().Debug().Model(&model.Article{}).Offset(option["offset"]).Limit(option["limit"])
	title, ok := option["title"].(string)
	ln := "%" + title + "%" // interface to string
	if ok {
		con = con.Where("title LIKE ?", ln)
	}
	cate_id, ok := option["category_id"]

	if ok {
		con = con.Where("category_id = ?", cate_id)
	}
	order := option["order"].(string)
	orderType := option["order_type"].(string)
	con = con.Order(order + " " + orderType)

	err = con.Find(&res).Error

	if err != nil {
		return res, err
	}
	return res, nil
}
