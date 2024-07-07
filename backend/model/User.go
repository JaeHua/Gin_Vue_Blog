package model

import (
	"backend/utils/errmsg"
	"encoding/base64"
	"errors"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(20);not null" json:"password" validate:"required,min=6,max=20" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色码"`
}

// CheckUser 查询用户
func CheckUser(name string) (code int) {
	var users User
	if DB == nil {
		log.Println("Database is not initialized")
		return errmsg.ERROR
	}
	DB.Select("id").Where("username=?", name).First(&users)
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS

}

// CreateUser 创建用户
func CreateUser(data *User) int {
	//加密操作
	data.Password = ScryptPw(data.Password)
	err := DB.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// GetUserInfo 查询单个用户信息
func GetUserInfo(id int) (User, int) {
	var user User
	err := DB.Where("id=?", id).First(&user).Error
	if err != nil {
		return user, errmsg.ERROR
	}
	return user, errmsg.SUCCESS
}

// GetUsers 分页查询用户列表
func GetUsers(username string, pageSize int, pageNum int) ([]User, int) {
	var user User
	var users []User
	var total int64

	// 分页查询
	if username == "" {
		// 查询所有用户
		err = DB.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("Error finding users: %v", err)
			return nil, 0
		}
	} else {
		// 模糊查询
		err = DB.Where("username LIKE ?", username+"%").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("Error finding users: %v", err)
			return nil, 0
		}
	}

	// 计算总数
	if username == "" {
		err = DB.Model(&user).Count(&total).Error
	} else {
		err = DB.Model(&user).Where("username LIKE ?", username+"%").Count(&total).Error
	}

	if err != nil {
		log.Printf("Error counting users: %v", err)
		return nil, 0
	}
	return users, int(total)
}

// ScryptPw 密码加密
func ScryptPw(password string) string {
	const keyLen = 10
	salt := make([]byte, 10)
	salt = []byte{3, 5, 45, 24, 64, 32, 12, 222}
	HashPw, err2 := scrypt.Key([]byte(password), salt, 16384, 8, 1, keyLen)
	if err2 != nil {
		log.Fatal(err2)
	}
	fpw := base64.StdEncoding.EncodeToString(HashPw)
	return fpw
}

// DeleteUser 删除用户信息
func DeleteUser(id int) int {
	var user User
	err := DB.Where("id=?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// EditUser 编辑用户信息
func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	//密码之外的信息更新，要更新密码应该是新的单独接口
	maps["username"] = data.Username
	maps["role"] = data.Role
	err := DB.Model(&user).Where("id=?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func CheckLogin(username string, password string) int {
	var user User
	DB.Where("username=?", username).First(&user)
	if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIST
	}
	if user.Password != ScryptPw(password) {
		return errmsg.ERROR_PASSWORD_WORNG
	}
	if user.Role != 1 {
		return errmsg.ERROR_USER_NO_RIGHT
	}
	return errmsg.SUCCESS
}
