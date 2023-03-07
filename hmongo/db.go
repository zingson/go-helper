package hmongo

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Option struct {
	Dsn string `json:"dsn"`
}

type DB struct {
	Option
	*mongo.Database
}

func NewDB(option Option, opts ...*options.ClientOptions) (db *DB) {
	db = &DB{Option: option}
	db.SetOption(option, opts...)
	return
}

func (c *DB) SetOption(option Option, opts ...*options.ClientOptions) *DB {
	c.Option = option
	c.Database = newDatabase(option.Dsn, opts...)
	return c
}

type ModelCollection interface {
	TableName() string
}

func (c *DB) Model(m ModelCollection, opts ...*options.CollectionOptions) *mongo.Collection {
	return c.Collection(m.TableName(), opts...)
}
