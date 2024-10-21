package initialize

import (
	"context"
	"fmt"

	"github.com/VanThen60hz/GoShop/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	r := global.Config.Redis

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", r.Host, r.Port),
		Password: r.Password,
		DB:       r.Database,
		PoolSize: 10,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("initialize Redis error: %v", zap.Error(err))
	}

	fmt.Println("initialize Redis successfully")
	global.Rdb = rdb
	redisExample()
}

func redisExample() {
	err := global.Rdb.Set(ctx, "score", "100", 0).Err()
	if err != nil {
		fmt.Println("set key error in redis ", zap.Error(err))
		return
	}

	value, err := global.Rdb.Get(ctx, "score").Result()
	if err != nil {
		fmt.Println("get key error in redis ", zap.Error(err))
		return
	}

	global.Logger.Info("value of score is: ", zap.String("score", value))
}
