package model

import (
	"meteor/global"
	"time"
)

// 自增主键表 利用生成
type ShortUrlConfig struct {
	UrlCode       string    `json:"url_code" form:"url_code" gorm:"column:url_code;comment:url编码"`
	ValidEnable   int8      `json:"valid_enable" form:"valid_enable" gorm:"column:valid_enable;comment:是否有效"`
	InvalidTime   time.Time `json:"invalid_time" form:"invalid_time" gorm:"column:invalid_time;comment:无效时间"`
	EncryptEnable int8      `json:"encrypt_enable" form:"encrypt_enable" gorm:"column:encrypt_enable;comment:是否加密"`
	EncryptPasswd string    `json:"encrypt_passwd" form:"encrypt_passwd" gorm:"column:encrypt_passwd;comment:密码"`
	SourceUrl     string    `json:"source_url" form:"source_url" gorm:"column:source_url;comment:长链地址"`
	global.GVA_MODEL
}

func (ShortUrlConfig) TableName() string {
	//实现TableName接口，以达到结构体和表对应，如果不实现该接口，gorm会自动扩展表名为articles（结构体+s）
	return "short_url_config"
}
