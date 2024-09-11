package model

import (
	"backend/utils/errmsg"
	"fmt"
	"github.com/jordan-wright/email"
	"math/rand"
	"net/smtp"
	"time"
)

// MailCode 验证码+号码结构体
type MailCode struct {
	Mail  string `json:"mail"`
	VCode string `json:"vcode"`
}

// SendEmailValidate 发送邮件
func SendEmailValidate(em []string) (string, int, error) {
	e := email.NewEmail()
	e.From = fmt.Sprintf("Blog <2624857134@qq.com>")
	e.To = em
	e.Subject = "神秘验证码"
	// 生成6位随机验证码
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vCode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	t := time.Now().Format("2006-01-02 15:04:05")

	//设置文件发送的内容
	content := fmt.Sprintf(`
        <div style="font-family: Arial, sans-serif; background-color: #f5f5f5; padding: 20px; border-radius: 5px;">
        <div style="background-color: #ffffff; padding: 20px; border-radius: 5px; box-shadow: 0 0 10px rgba(0,0,0,0.1);">
        <div style="font-size: 18px; font-weight: bold; color: #333333; margin-bottom: 10px;">
        尊敬的 %s，您好！
        </div>
        <div style="padding: 8px 40px 8px 50px; background-color: #f2f2f2; border-radius: 5px;">
        <p style="color: #555555;">您于 %s 提交的邮箱验证，本次验证码为<u><strong style="font-size: 20px; color: #ff6600;">%s</strong></u>，为了保证账号安全，验证码有效期为5分钟。请确认为本人操作，切勿向他人泄露，感谢您的理解与使用。</p>
        </div>
        <div style="margin-top: 10px;">
        <p style="color: #777777; font-size: 14px;">此邮箱为系统邮箱，请勿回复。</p>
        </div>
        </div>
        </div>
        `, em[0], t, vCode)
	e.Text = []byte(content)

	//设置服务器相关的配置
	_ = e.Send("smtp.qq.com:25", smtp.PlainAuth("", "2624857134@qq.com", "ggtwjomswvvheaci", "smtp.qq.com"))

	return vCode, errmsg.SUCCESS, err
}
