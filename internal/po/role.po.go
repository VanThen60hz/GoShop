package po

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	ID       int64  `gorm:"primaryKey; autoIncrement; not null; comment:'primary key is ID'"`
	RoleName string `gorm:"column:role_name; type:varchar(255)"`
	RoleNote string `gorm:"column:role_note; type:text"`
}

func (r *Role) TableName() string {
	return "go_db_role"
}
