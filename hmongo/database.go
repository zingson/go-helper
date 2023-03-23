package hmongo

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
	"sync"
)

var dbCache sync.Map

// NowDatabase 当前数据库连接
func NowDatabase(dsn string, opts ...*options.ClientOptions) (database *mongo.Database) {
	if v, ok := dbCache.Load(dsn); ok {
		database = v.(*mongo.Database)
		return
	}
	database = newDatabase(dsn, opts...)
	dbCache.Store(dsn, database)
	return
}

// newDatabase 新建数据库连接
func newDatabase(dsn string, opts ...*options.ClientOptions) *mongo.Database {
	if dsn == "" {
		panic(errors.New("ERROR:MongoDB dsn is nil"))
	}

	var dbName = (strings.Split((strings.Split(dsn, "/"))[3], "?"))[0]
	opts = append(opts, options.Client().ApplyURI(dsn))
	client, err := mongo.Connect(context.Background(), opts...)
	if err != nil {
		panic(err)
	}
	return client.Database(dbName)
}
