package pg

import (
	"database/sql"
	"fmt"
	"github.com/fzft/go-arc/template/db"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"sync"
)

type Psql struct {
	Client *sql.DB
}

var instance *Psql
var once sync.Once

func GetStore() db.DB {

	once.Do(func() {
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			viper.GetString("pq.host"),
			viper.GetInt("pq.port"),
			viper.GetString("pq.user"),
			viper.GetString("pq.password"),
			viper.GetString("pq.db"))
		client, err := sql.Open("postgres", psqlInfo)

		if err != nil {
			panic(fmt.Sprintf("connect to pq error: %v ", err))
		}

		// Open doesn't open a connection. Validate DSN data:
		err = client.Ping()
		if err != nil {
			panic(fmt.Sprintf("ping pq error: %v ", err))
		}

		instance = &Psql{
			Client: client,
		}
	})

	return instance
}

func (m *Psql) Close() {
	if m.Client == nil {
		return
	}

	m.Client.Close()
}
