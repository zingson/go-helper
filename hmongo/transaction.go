package hmongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

// Transaction 数据库事务
func Transaction(client *mongo.Client, fn func(sessionContext mongo.SessionContext) (interface{}, error)) (interface{}, error) {
	ctx := context.Background()
	session, err := client.StartSession()
	if err != nil {
		return nil, err
	}
	defer session.EndSession(ctx)
	return session.WithTransaction(ctx, func(sessCtx mongo.SessionContext) (interface{}, error) {
		return fn(sessCtx)
	})
}
