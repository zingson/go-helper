package hmongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
使用说明：
type JobDao struct {
	hmongo.DaoImpl[*Job]
}

func NewJobDao() *JobDao {
	return &JobDao{*hmongo.NewDaoImpl[*Job](db)}
}

*/

type Table interface {
	TableName() string
}

func NewDaoImpl[T Table](db *mongo.Database) *DaoImpl[T] {
	return new(DaoImpl[T]).SetDatabase(db)
}

// DaoImpl Mongodb 常用操作实现
type DaoImpl[T Table] struct {
	model   T
	db      *mongo.Database
	colOpts []*options.CollectionOptions
}

func (o *DaoImpl[T]) SetDatabase(db *mongo.Database) *DaoImpl[T] {
	o.db = db
	return o
}

func (o *DaoImpl[T]) Database() *mongo.Database {
	if o.db == nil {
		panic("core.DaoImpl.db is nil")
	}
	return o.db
}

func (o *DaoImpl[T]) SetCollectionOptions(opts ...*options.CollectionOptions) *DaoImpl[T] {
	o.colOpts = opts
	return o
}

func (o *DaoImpl[T]) Collection(opts ...*options.CollectionOptions) (c *mongo.Collection) {
	return o.Database().Collection(o.model.TableName(), append(o.colOpts, opts...)...)
}

func (o *DaoImpl[T]) InsertOne(ctx context.Context, document any) (id string, err error) {
	r, err := o.Collection().InsertOne(ctx, document)
	if err != nil {
		return
	}
	id = fmt.Sprintf("%v", r.InsertedID)
	return
}

func (o *DaoImpl[T]) InsertMany(ctx context.Context, documents []any) (r *mongo.InsertManyResult, err error) {
	return o.Collection().InsertMany(ctx, documents)
}

func (o *DaoImpl[T]) UpdateOne(ctx context.Context, filter bson.D, update bson.M, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return o.Collection().UpdateOne(ctx, filter, update, opts...)
}

func (o *DaoImpl[T]) UpdateMany(ctx context.Context, filter bson.D, update bson.M, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return o.Collection().UpdateMany(ctx, filter, update, opts...)
}

func (o *DaoImpl[T]) DeleteOne(ctx context.Context, filter bson.D) (*mongo.DeleteResult, error) {
	return o.Collection().DeleteOne(ctx, filter)
}

func (o *DaoImpl[T]) Count(ctx context.Context, filter bson.D, opts ...*options.CountOptions) (total int64, err error) {
	return o.Collection().CountDocuments(ctx, filter, opts...)
}

func (o *DaoImpl[T]) FindAll(ctx context.Context) (values []*T, err error) {
	return o.Find(ctx, bson.D{}, options.Find().SetSort(bson.M{"_id": -1}))
}

func (o *DaoImpl[T]) FindOne(ctx context.Context, filter bson.D, opts ...*options.FindOneOptions) (value *T, err error) {
	err = o.Collection().FindOne(ctx, filter, opts...).Decode(&value)
	return
}

func (o *DaoImpl[T]) FindOneAndUpdate(ctx context.Context, filter bson.D, update bson.M, opts ...*options.FindOneAndUpdateOptions) (value *T, err error) {
	err = o.Collection().FindOneAndUpdate(ctx, filter, update, opts...).Decode(&value)
	return
}

func (o *DaoImpl[T]) Find(ctx context.Context, filter bson.D, opts ...*options.FindOptions) (list []*T, err error) {
	return DaoFind[T](o.Collection(), ctx, filter, opts...)
}

func DaoFind[R any](col *mongo.Collection, ctx context.Context, filter bson.D, opts ...*options.FindOptions) (list []*R, err error) {
	return DaoCursor[R](col.Find(ctx, filter, opts...))
}

func (o *DaoImpl[T]) FindPage(ctx context.Context, filter bson.D, sort bson.D, pageNum, pageSize int64) (list []*T, total int64, err error) {
	if sort == nil || len(sort) == 0 {
		sort = append(sort, bson.E{"_id", -1})
	}
	if pageNum < 1 {
		pageNum = 1
	}
	if pageSize < 1 {
		pageSize = 20
	}

	list, err = o.Find(ctx, filter, options.Find().SetSort(sort).SetSkip((pageNum-1)*pageSize).SetLimit(pageSize))
	if err != nil {
		return
	}

	total, err = o.Count(ctx, filter)
	if err != nil {
		return
	}
	return
}

func (o *DaoImpl[T]) Aggregate(ctx context.Context, pipeline bson.A, opts ...*options.AggregateOptions) (list []*T, err error) {
	return o.Cursor(o.Collection().Aggregate(ctx, pipeline, opts...))
}

// DaoAggregate 聚合查询返回自定义类型
func DaoAggregate[R any](c *mongo.Collection, ctx context.Context, pipeline bson.A, opts ...*options.AggregateOptions) (list []*R, err error) {
	return DaoCursor[R](c.Aggregate(ctx, pipeline, opts...))
}

func DaoCursor[R any](cursor *mongo.Cursor, e error) (values []*R, err error) {
	if err = e; err != nil {
		return
	}
	if err = cursor.Err(); err != nil {
		return
	}
	if err = cursor.All(context.TODO(), &values); err != nil {
		return
	}
	return
}

func (o *DaoImpl[T]) Cursor(cursor *mongo.Cursor, e error) (values []*T, err error) {
	return DaoCursor[T](cursor, e)
}

func (o *DaoImpl[T]) Transaction(fn func(sessionContext mongo.SessionContext) (any, error), opts ...*options.SessionOptions) (v any, err error) {
	ctx := context.Background()
	session, err := o.db.Client().StartSession(opts...)
	if err != nil {
		return
	}
	defer session.EndSession(ctx)
	return session.WithTransaction(ctx, fn)
}
