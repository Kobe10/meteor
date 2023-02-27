package model

import "meteor/global"

// 自增主键表 利用生成
type App struct {
	AppId           string `json:"app_id" form:"app_id" gorm:"column:app_id;comment:应用id"`
	Uid             string `json:"uid" form:"uid" gorm:"column:uid;comment:用户uid"`
	AccessKeyId     string `json:"access_key_id" form:"access_key_id" gorm:"column:access_key_id;comment:应用用户名"`
	AccessKeySecret string `json:"access_key_secret" form:"access_key_secret" gorm:"column:access_key_secret;comment:应用密码"`
	Status          int    `json:"status" form:"status" gorm:"column:status;comment:应用状态  "`
	global.GVA_MODEL
}

func (App) TableName() string {
	//实现TableName接口，以达到结构体和表对应，如果不实现该接口，gorm会自动扩展表名为articles（结构体+s）
	return "app"
}
