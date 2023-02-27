package model

import "meteor/global"

// 自增主键表 利用生成
type User struct {
	UserCode string `json:"user_code" form:"user_code" gorm:"column:user_code;comment:用户编码"`
	Uid      string `json:"uid" form:"uid" gorm:"column:uid;comment:用户uid"`
	Phone    string `json:"phone" form:"phone" gorm:"column:phone;comment:用户手机号"`
	Status   string `json:"status" form:"status" gorm:"column:status;comment:用户状态"`
	global.GVA_MODEL
}

func (User) TableName() string {
	//实现TableName接口，以达到结构体和表对应，如果不实现该接口，gorm会自动扩展表名为articles（结构体+s）
	return "user"
}
