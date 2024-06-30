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
	var total int64

	// 先计算总数
	err := DB.Model(&Category{}).Count(&total).Error
	if err != nil {
		log.Printf("Error counting categories: %v", err)
		return nil, 0
	}

	// 分页查询
	err = DB.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&cate).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("Error finding categories: %v", err)
		return nil, int(total)
	}
	return cate, int(total)
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
// GetCateArt 查询分类下的所有文章
func GetCateArt(id int, pageSize int, pageNum int) ([]Article, int, int) {
	var cateArts []Article
	var total int64

	// 先计算总数
	err := DB.Model(&Article{}).Where("cid = ?", id).Count(&total).Error
	if err != nil {
		log.Printf("Error counting articles in category %d: %v", id, err)
		return nil, errmsg.ERROR_CATE_NOT_EXIST, 0
	}

	// 分页查询
	err = DB.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid = ?", id).Find(&cateArts).Error
	if err != nil {
		log.Printf("Error finding articles in category %d: %v", id, err)
		return nil, errmsg.ERROR_CATE_NOT_EXIST, int(total)
	}
	return cateArts, errmsg.SUCCESS, int(total)
}
