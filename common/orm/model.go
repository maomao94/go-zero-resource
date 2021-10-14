package orm

import (
	"time"

	"gorm.io/gorm"
)

type GVA_MODEL struct {
	ID        uint           `gormx:"primarykey"` // 主键ID
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gormx:"index" json:"-"` // 删除时间
}
