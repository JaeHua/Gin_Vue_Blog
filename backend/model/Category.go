package model

import (
	"backend/utils/errmsg"
	"errors"
	"github.com/jinzhu/gorm"
	"log"
)

type Category struct {
	ID   int    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// CheckCate 查询用户
func CheckCate(name string) (code int) {
	var cate Category
	if DB == nil {
		log.Println("Database is not initialized")
		return errmsg.ERROR
	}
	DB.Select("id").Where("name=?", name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ERROR_CATEGORY_USED
	}
	return errmsg.SUCCESS

}

// CreateCate 创建分类
func CreateCate(data *Category) int {

	err := DB.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetCate 分页查询分类列表
func GetCate(pageSize int, pageNum int) ([]Category, int) {
	var cate []Category
	var total int
	//分页查询核心逻辑
	err := DB.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cate).Count(&total).Error
	if err != nil && !errors.Is(gorm.ErrRecordNotFound, err) {
		return nil, 0
	}
	return cate, total
}

// DeleteCate 删除分类信息
func DeleteCate(id int) int {
	var cate Category
	err := DB.Where("id=?", id).Delete(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// EditCate 编辑分类信息
func EditCate(id int, data *Category) int {
	var cate Category
	var maps = make(map[string]interface{})
	maps["name"] = data.Name

	err := DB.Model(&cate).Where("id=?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetCateArt 查询分类下的所有文章
func GetCateArt(id int, pageSize int, pageNum int) ([]Article, int, int) {
	var cateArts []Article
	var total int
	err := DB.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid=?", id).Find(&cateArts).Count(&total).Error
	if err != nil {
		return nil, errmsg.ERROR_CATE_NOT_EXIST, 0

	}
	return cateArts, errmsg.SUCCESS, total
}
