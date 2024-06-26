package model

import (
	"backend/utils/errmsg"
	"errors"
	"github.com/jinzhu/gorm"
)

type Article struct {
	Category Category
	gorm.Model
	Title   string `gorm:"type:varchar(100);not null" json:"title"`
	Cid     int    `gorm:"type:int;not null" json:"cid"`
	Desc    string `gorm:"type:varchar(200)" json:"desc"`
	Content string `gorm:"type:longtext" json:"content"`
	Img     string `gorm:"type:varchar(100)" json:"img"`
}

// CreateArt 创建文章
func CreateArt(data *Article) int {

	err := DB.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// todo 查询单个文章

//todo 查询分类下的所有文章

// GetArt 分页查询文章列表
func GetArt(pageSize int, pageNum int) []Article {
	var art []Article
	//分页查询核心逻辑
	err = DB.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&art).Error
	if err != nil && !errors.Is(gorm.ErrRecordNotFound, err) {
		return nil
	}
	return art
}

// DeleteArt 删除文章信息
func DeleteArt(id int) int {
	var art Article
	err := DB.Where("id=?", id).Delete(&art).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// EditArt 编辑文章信息
func EditArt(id int, data *Article) int {
	var art Article
	var maps = make(map[string]interface{})
	maps["title"] = data.Title
	maps["cid"] = data.Cid
	maps["desc"] = data.Desc
	maps["content"] = data.Content
	maps["img"] = data.Img
	err := DB.Model(&art).Where("id=?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
