package redis

import (
	"context"
	"fmt"
	"github.com/fzft/go-arc/template/db"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"sync"
)

type Redis struct {
	Client *redis.Client
}

var instance *Redis
var once sync.Once

func GetStore() db.DB {

	once.Do(func() {

		ctx := context.Background()
		connAddr := fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port"))
		client := redis.NewClient(&redis.Options{
			Addr:     connAddr,
			Password: viper.GetString("redis.password"),
			DB:       viper.GetInt("redis.db"),
		})

		_, err := client.Ping(ctx).Result()
		if err != nil {
			panic(fmt.Sprintf("connect to redis error: %v ", err))
		}
		instance = &Redis{
			Client: client,
		}

	})

	return instance
}

func (m *Redis) Close() {
	if m.Client == nil {
		return
	}

	m.Client.Close()
}
