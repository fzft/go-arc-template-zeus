package mongo

import (
	"context"
	"fmt"
	zeus "github.com/fzft/go-arc-template-zeus"
	"github.com/fzft/go-arc/template/db"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

type Mongo struct {
	Client *mongo.Client
}

var instance *Mongo
var once sync.Once

func GetStore() db.DB {

	once.Do(func() {
		clientOptions := options.Client().ApplyURI(viper.GetString("mongo.uri")).SetServerAPIOptions(options.ServerAPI(options.ServerAPIVersion1))
		client, err := mongo.Connect(context.Background(), clientOptions)

		if err != nil {
			panic(fmt.Sprintf("connect to mongo error: %v ", err))
		}

		err = client.Ping(context.Background(), nil)
		if err != nil {
			panic(fmt.Sprintf("ping mongo error: %v ", err))
		}

		instance = &Mongo{
			Client: client,
		}

		zeus.Logger.Info(fmt.Sprintf("mongo connected: %s ", viper.GetString("mongo.uri")))
	})

	return instance
}

func (m *Mongo) Close() {
	if m.Client == nil {
		return
	}

	m.Client.Disconnect(context.Background())
}
