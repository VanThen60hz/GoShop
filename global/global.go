package global

import (
	"github.com/VanThen60hz/GoShop/pkg/logger"
	"github.com/VanThen60hz/GoShop/pkg/setting"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
)
