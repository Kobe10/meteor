package model

import "meteor/global"

// 自增主键表 利用生成
type ClickUrlRecord struct {
	UrlCode string `json:"url_code" form:"url_code" gorm:"column:url_code;comment:url编码"`
	ClickIp string `json:"click_ip" form:"click_ip" gorm:"column:click_ip;comment:点击用户的ip"`
	Browser string `json:"browser" form:"browser" gorm:"column:browser;comment:用户浏览器"`
	Phone   string `json:"phone" form:"phone" gorm:"column:phone;comment:点击用户手机号"`
	Area    string `json:"area" form:"area" gorm:"column:area;comment:用户地区"`
	global.GVA_MODEL
}

func (ClickUrlRecord) TableName() string {
	//实现TableName接口，以达到结构体和表对应，如果不实现该接口，gorm会自动扩展表名为articles（结构体+s）
	return "click_url_record"
}
