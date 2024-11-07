package table

import (
	"time"
)

// FileMeta 文件元数据表
type FileMeta struct {
	ID        int64      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	FileName  string     `gorm:"column:file_name;type:varchar(512);not null;default:''" json:"file_name" binding:"required"`
	FileSize  int64      `gorm:"column:file_size;type:bigint;not null;default:0" json:"file_size"`
	Owner     string     `gorm:"column:owner;type:varchar(32);not null;default:''" json:"owner"`
	CreatedAt time.Time  `gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at;type:datetime" json:"deleted_at,omitempty"`
}

// TableName 指定表名
func (FileMeta) TableName() string {
	return "t_file_meta"
}
