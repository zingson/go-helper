package hmongo

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB
// deprecated
type DB struct {
	Option
	*mongo.Database
}

// NewDB
// deprecated
func NewDB(option Option, opts ...*options.ClientOptions) (db *DB) {
	db = &DB{Option: option}
	db.SetOption(option, opts...)
	return
}

// SetOption deprecated
func (c *DB) SetOption(option Option, opts ...*options.ClientOptions) *DB {
	c.Option = option
	c.Database = newDatabase(option.Dsn, opts...)
	return c
}

// ModelCollection deprecated
type ModelCollection interface {
	TableName() string
}

// Model deprecated
func (c *DB) Model(m ModelCollection, opts ...*options.CollectionOptions) *mongo.Collection {
	return c.Collection(m.TableName(), opts...)
}
