package model

import (
	"backend/utils/errmsg"
)

type Profile struct {
	ID     int    `gorm:"primaryKey" json:"id"`
	Name   string `gorm:"type:varchar(20)" json:"name" `
	QQ     string `gorm:"type:varchar(20)" json:"qq"`
	Desc   string `gorm:"type:varchar(200)" json:"desc"`
	Email  string `gorm:"type:varchar(200)" json:"email"`
	Site   string `gorm:"type:varchar(200)" json:"site"`
	Img    string `gorm:"type:varchar(200)" json:"img"`
	Github string `gorm:"type:varchar(200)" json:"github"`
	Avatar string `gorm:"type:varchar(200)" json:"avatar"`
}

// CreateProfile 创建用户信息
func CreateProfile(data *Profile) int {
	//加密操作
	err := DB.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetProfile 获取个人信息
func GetProfile(email string) (Profile, int) {
	var profile Profile
	err = DB.Where("email=?", email).First(&profile).Error
	if err != nil {
		return profile, errmsg.ERROR
	}
	return profile, errmsg.SUCCESS
}

// UpdateProfile 更新个人信息
func UpdateProfile(email string, data *Profile) int {
	var profile1 Profile
	var updates = make(map[string]interface{})

	if data.Name != "" {
		updates["name"] = data.Name
	}
	if data.QQ != "" {
		updates["qq"] = data.QQ
	}
	if data.Desc != "" {
		updates["desc"] = data.Desc
	}
	//if data.Email != "" {
	//	updates["email"] = data.Email
	//}
	if data.Site != "" {
		updates["site"] = data.Site
	}
	if data.Img != "" {
		updates["img"] = data.Img
	}
	if data.Github != "" {
		updates["github"] = data.Github
	}
	if data.Avatar != "" {
		updates["avatar"] = data.Avatar
	}
	err = DB.Model(&profile1).Where("email=?", email).Updates(updates).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}
