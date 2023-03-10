package hmongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Find 列表查询
// deprecated
func Find(ctx context.Context, c *mongo.Collection, filter interface{}, result interface{}, opts ...*options.FindOptions) (err error) {
	cursor, err := c.Find(ctx, filter, opts...)
	if err != nil {
		return err
	}
	if cursor.Err() != nil {
		return cursor.Err()
	}
	err = cursor.All(ctx, result)
	if err != nil {
		return err
	}
	return
}

// FindPage 分页查询
// deprecated
func FindPage(ctx context.Context, c *mongo.Collection, filter interface{}, sort bson.D, pageSize, pageNum int64, result interface{}, fo ...*options.FindOptions) (total int64, err error) {
	if sort == nil {
		sort = bson.D{bson.E{"_id", -1}} // 没有排序字段时，根据_id 倒序
	}
	if pageNum <= 0 {
		pageNum = 1
	}
	if pageSize <= 0 {
		pageSize = 20
	}
	op := options.Find().SetSort(sort).SetSkip(pageSize * (pageNum - 1)).SetLimit(pageSize)
	fo = append(fo, op)
	err = Find(ctx, c, filter, result, fo...)
	if err != nil {
		return 0, err
	}
	return c.CountDocuments(ctx, filter)
}

// FindList 列表查询
// deprecated
func FindList[R any](ctx context.Context, c *mongo.Collection, filter interface{}, opts ...*options.FindOptions) (list R, err error) {
	err = Find(ctx, c, filter, &list, opts...)
	return
}

// FindListPage 分页查询
// deprecated
func FindListPage[R any](ctx context.Context, c *mongo.Collection, filter interface{}, sort bson.D, pageSize, pageNum int64, fo ...*options.FindOptions) (list R, total int64, err error) {
	total, err = FindPage(ctx, c, filter, sort, pageSize, pageNum, &list, fo...)
	return
}
