/*
@Time : 2025/11/02 10:00
@Author : bibinocode
@File : user_model.go
@Software: GoLand
*/

package models

type UserModel struct {
	Model
	Nickname       string `gorm:"size:32 comment:'昵称'" json:"nickname"`
	Username       string `gorm:"size:32 comment:'用户名'" json:"username"`
	Password       string `gorm:"size:64 comment:'密码'" json:"password"`
	Openid         string `gorm:"size:64 comment:'openid'" json:"openId"`
	Tel            string `gorm:"size:12 comment:'手机号'" json:"tel"`
	Email          string `gorm:"size:128 comment:'邮箱'" json:"email"`
	RegisterSource int8   `gorm:"comment:'注册来源'" json:"registerSource"` // 0 手机号 1 邮箱 2 微信
	Avatar         string `gorm:"size:255 comment:'头像'" json:"avatar"`
	Scope          int    `gorm:"comment:'积分'" json:"scope"`
}
