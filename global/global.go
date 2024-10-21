package global

import (
	"github.com/VanThen60hz/GoShop/pkg/logger"
	"github.com/VanThen60hz/GoShop/pkg/setting"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Rdb    *redis.Client
	Mdb    *gorm.DB
)
