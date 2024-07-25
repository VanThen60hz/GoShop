package initialize

import (
	"fmt"

	"github.com/VanThen60hz/GoShop/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./config/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	// read config
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("failed to read configuration %w", err))
	}

	// read server config
	// fmt.Println("Server Port: ", viper.GetInt("server.port"))
	// fmt.Println("Server Key: ", viper.GetString("security.jwt.key"))

	// config structure
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable to decode configuration %v", err)
	}
}
