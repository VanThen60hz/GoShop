package initialize

import (
	"fmt"

	"github.com/VanThen60hz/GoShop/global"
	"go.uber.org/zap"
)

func Run() {
	LoadConfig()

	m := global.Config.Mysql
	fmt.Println("Loading configuration mysql: ", m.Username, m.Password)
	InitLogger()
	global.Logger.Info("Config log ok!", zap.String("ok", "success"))
	InitMysql()
	InitRedis()
	InitRouter()

	r := InitRouter()

	r.Run(":8888")
}
