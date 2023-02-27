package model

import "meteor/global"

// 自增主键表 利用生成
type Incr struct {
	global.GVA_MODEL
}

func (Incr) TableName() string {
	//实现TableName接口，以达到结构体和表对应，如果不实现该接口，gorm会自动扩展表名为articles（结构体+s）
	return "incr"
}
