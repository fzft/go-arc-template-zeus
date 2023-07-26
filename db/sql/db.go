package sql

import (
	"database/sql"
	"fmt"
	"github.com/fzft/go-arc/template/db"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"sync"
)

type Mysql struct {
	Client *sql.DB
}

var instance *Mysql
var once sync.Once

func GetStore() db.DB {

	once.Do(func() {
		dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
			viper.GetString("mysql.host"),
			viper.GetString("mysql.port"),
			viper.GetString("mysql.user"),
			viper.GetString("mysql.password"),
			viper.GetString("mysql.db"))
		client, err := sql.Open("mysql", dataSourceName)
		if err == nil {

			// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
			client.SetConnMaxLifetime(viper.GetDuration("mysql.max_lifetime"))
			client.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns"))
			client.SetMaxOpenConns(viper.GetInt("mysql.max_open_conns"))
			instance = &Mysql{
				Client: client,
			}
		} else {
			panic(fmt.Sprintf("connect to mysql error: %v ", err))
		}
	})

	return instance
}

func (m *Mysql) Close() {
	if m.Client == nil {
		return
	}

	m.Client.Close()
}
