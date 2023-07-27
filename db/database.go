package db

import (
	"fmt"
	zeus "github.com/fzft/go-arc-template-zeus"
	"github.com/fzft/go-arc-template-zeus/db/mongo"
	"github.com/fzft/go-arc-template-zeus/db/pg"
	"github.com/fzft/go-arc-template-zeus/db/redis"
	"github.com/fzft/go-arc-template-zeus/db/sql"
	"github.com/spf13/viper"
)

type DBRegistry map[string]DB

func (r DBRegistry) keys() []string {
	keys := make([]string, 0, len(r))
	for k := range r {
		keys = append(keys, k)
	}

	return keys
}

func (r DBRegistry) Close() {
	for _, db := range r {
		db.Close()
	}
}

var registry = make(DBRegistry)

const (
	MySQL string = "mysql"
	Mongo        = "mongo"
	PgSQL        = "pgsql"
	Redis        = "redis"
)

// DB is the interface that wraps the basic methods of a database.
type DB interface {
	Close()
}

// DBInit initializes the database. through the configuration file.
func DBInit() {

	dbConfig := viper.GetStringMap("db")
	for db, _ := range dbConfig {
		switch db {
		case MySQL:
			registry[MySQL] = sql.GetStore()
		case Mongo:
			registry[Mongo] = mongo.GetStore()
		case PgSQL:
			registry[PgSQL] = pg.GetStore()
		case Redis:
			registry[Redis] = redis.GetStore()
		}

	}

	zeus.Logger.Info(fmt.Sprintf("db registry: %v", registry.keys()))

}

func DBClose() {
	registry.Close()
}
