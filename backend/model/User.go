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
	Username string `gorm:"type:varchar(20);not null" json:"username"`
	Password string `gorm:"type:varchar(20);not null" json:"password"`
	Role     int    `gorm:"type:int" json:"role"`
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

// GetUsers 分页查询用户列表
func GetUsers(pageSize int, pageNum int) []User {
	var users []User
	//分页查询核心逻辑
	err = DB.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&users).Error
	if err != nil && !errors.Is(gorm.ErrRecordNotFound, err) {
		return nil
	}
	return users
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
	if user.Role != 0 {
		return errmsg.ERROR_USER_NO_RIGHT
	}
	return errmsg.SUCCESS
}
