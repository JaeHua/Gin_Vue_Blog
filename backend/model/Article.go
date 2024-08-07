package model

import (
	"backend/utils/errmsg"
	"errors"
	"github.com/jinzhu/gorm"
	"log"
)

type Article struct {
	Category Category `gorm:"foreignKey:Cid"`
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

// GetArtInfo 查询单个文章
func GetArtInfo(id int) (Article, int) {
	var art Article
	err := DB.Preload("Category").Where("id=?", id).First(&art).Error
	if err != nil {
		return art, errmsg.ERROR_ART_NOT_EXIST
	}
	return art, errmsg.SUCCESS
}

//todo 查询分类下的所有文章

// GetArt 分页查询文章列表
func GetArt(title string, pageSize int, pageNum int) ([]Article, int, int) {
	var arts []Article
	var total int64

	if title == "" {
		err = DB.Order("Updated_At DESC").Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&arts).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("Error finding articles: %v", err)
			return nil, errmsg.ERROR, int(total)
		}
	} else {
		err = DB.Order("Updated_At DESC").Preload("Category").Where("title LIKE ?", title+"%").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&arts).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("Error finding articles: %v", err)
			return nil, errmsg.ERROR, int(total)
		}
	}
	// 先计算总数
	if title == "" {
		err = DB.Model(&Article{}).Count(&total).Error

	} else {
		err = DB.Model(&Article{}).Where("title LIKE ?", title+"%").Count(&total).Error

	}
	if err != nil {
		log.Printf("Error counting articles: %v", err)
		return nil, errmsg.ERROR, 0
	}
	return arts, errmsg.SUCCESS, int(total)
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
