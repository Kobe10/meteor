package global

import (
	"time"
)

type GVA_MODEL struct {
	ID             uint `gorm:"primarykey"`
	IsDel          int
	CreateId       string
	LastModifyId   string
	CreateTime     time.Time
	LastModifyTime time.Time
}
