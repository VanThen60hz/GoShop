package po

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       int64     `gorm:"primaryKey; autoIncrement; not null"`
	UUID     uuid.UUID `gorm:"column:uuid; type:varchar(36); not null; index:idx_uuid; unique"`
	UserName string    `gorm:"column:user_name; type:varchar(255)"`
	IsActive bool      `gorm:"column:is_active; type:boolean"`
	Roles    []Role    `gorm:"many2many:go_user_roles;"`
}

func (u *User) TableName() string {
	return "go_db_user"
}
