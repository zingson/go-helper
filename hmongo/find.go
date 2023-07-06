package hmongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Find 列表查询
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

// Aggregate 聚合查询
func Aggregate(ctx context.Context, c *mongo.Collection, pipeline interface{}, result interface{}, opts ...*options.AggregateOptions) (err error) {
	cursor, err := c.Aggregate(ctx, pipeline, opts...)
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

// AggregatePage 待测试
// deprecated
func AggregatePage(ctx context.Context, c *mongo.Collection, pipeline bson.A, pageNum, pageSize int64, result interface{}, opts ...*options.AggregateOptions) (total int64, err error) {
	if pageNum <= 0 {
		pageNum = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}
	err = Aggregate(ctx, c, append(pipeline, bson.M{"$skip": pageSize * (pageNum - 1)}, bson.M{"$limit": pageSize}), result, opts...)
	if err != nil {
		return
	}

	var countMap []map[string]int64
	err = Aggregate(ctx, c, append(pipeline, bson.M{"$count": "count"}), &countMap, opts...)
	if err != nil {
		return
	}
	if countMap != nil {
		for _, m := range countMap {
			total = m["count"]
		}
	}
	return
}

// AggregateSearch 集合多字段关键字搜索
// deprecated
func AggregateSearch(ctx context.Context, c *mongo.Collection, filter bson.M, concat bson.A, keywords string, sort bson.D, pageNum, pageSize int64, result interface{}, opts ...*options.AggregateOptions) (total int64, err error) {
	if filter == nil {
		filter = bson.M{}
	}
	// 排序必须有一个字段
	if sort == nil {
		sort = bson.D{bson.E{"_id", -1}}
	}
	if pageNum == 0 {
		pageNum = 1
	}
	if pageSize == 0 {
		pageSize = 10
	}

	// 拼接字段，关键字匹配搜索
	searchText := bson.M{"$match": bson.M{}}
	match := bson.M{"$match": bson.M{}}
	if keywords != "" && concat != nil {
		concatA := bson.A{}
		for _, v := range concat {
			concatA = append(concatA, bson.M{"$ifNull": bson.A{bson.M{"$toString": v}, ""}}, ":") // 关键字查询，注意：concat不能拼接空字段
		}
		searchText = bson.M{"$set": bson.M{"search_text": bson.M{"$concat": concatA}}}
		match = bson.M{"$match": bson.M{"search_text": bson.M{"$regex": keywords}}}
	}
	skip := bson.M{"$skip": pageSize * (pageNum - 1)}
	limit := bson.M{"$limit": pageSize}
	err = Aggregate(ctx, c, bson.A{bson.M{"$match": filter}, searchText, match, bson.M{"$sort": sort}, skip, limit}, result, opts...)
	if err != nil {
		return
	}
	var countmap []map[string]int64
	err = Aggregate(ctx, c, bson.A{bson.M{"$match": filter}, searchText, match, bson.M{"$count": "count"}}, &countmap, opts...)
	if err != nil {
		return
	}
	if countmap != nil {
		for _, m := range countmap {
			total = m["count"]
		}
	}
	return
}

// FindList 列表查询
func FindList[R any](ctx context.Context, c *mongo.Collection, filter interface{}, opts ...*options.FindOptions) (list R, err error) {
	err = Find(ctx, c, filter, &list, opts...)
	return
}

// FindListPage 分页查询
func FindListPage[R any](ctx context.Context, c *mongo.Collection, filter interface{}, sort bson.D, pageSize, pageNum int64, fo ...*options.FindOptions) (list R, total int64, err error) {
	total, err = FindPage(ctx, c, filter, sort, pageSize, pageNum, &list, fo...)
	return
}
