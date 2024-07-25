package initialize

import (
	"github.com/VanThen60hz/GoShop/global"
	"github.com/VanThen60hz/GoShop/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
