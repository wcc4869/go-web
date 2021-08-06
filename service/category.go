package service

import (
	"go-web-edu/model"
)

func CreateCategory(category *model.Category) (err error) {
	err = GetDB().Create(&category).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateCategory(category *model.Category) (err error) {
	err = GetDB().Save(&category).Error
	if err != nil {
		return err
	}
	return nil
}

func GetCategoryList(option map[string]interface{}) (res []model.Category, err error) {

	con := GetDB().Model(&model.Category{}).Offset(option["offset"]).Limit(option["limit"])
	name, ok := option["name"].(string)
	ln := "%" + name + "%" // interface to string
	if ok {
		con = con.Where("name LIKE ?", ln)
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

func GetCategoryById(id int) (category *model.Category, err error) {
	category = &model.Category{}
	err = GetDB().First(&category, id).Error

	return category, err
}

func DeleteCategory(category *model.Category, id int) bool {
	err := GetDB().Where("id=?", id).Delete(&category).Error
	if err != nil {
		return false
	} else {
		return true
	}
}
