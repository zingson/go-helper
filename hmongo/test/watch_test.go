package test

import (
	"context"
	"github.com/zingson/go-helper/hmongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
	"time"
)

type Test struct {
	Name  string    `bson:"name"`
	Age   int64     `bson:"age"`
	Ctime time.Time `bson:"ctime"`
}

func (Test) TableName() string {
	return "test"
}

type TestDao struct {
	*hmongo.DaoImpl[Test]
}

func NewTestDao() *TestDao {
	return &TestDao{hmongo.NewDaoImpl[Test](hmongo.NowDatabase(hmongo.Dsn()))}
}

func TestWatch(t *testing.T) {

	// bson.M{"$match": bson.D{{"fullDocument.name", "dui"}}}
	changeStream, err := NewTestDao().Collection().Watch(nil, bson.A{}, options.ChangeStream().SetBatchSize(100))
	if err != nil {
		t.Error(err.Error())
		return
	}
	defer changeStream.Close(context.Background())
	for changeStream.Next(context.Background()) {
		var val map[string]any
		err = changeStream.Decode(&val)
		if err != nil {
			t.Error(err.Error())
			continue
		}
		t.Logf("%v", val)
		// map[_id:map[_data:82640B0865000000012B022C0100296E5A10044F1F74C7C8134A158132431FBABEA1EA46645F696400645EE84144E6892F0A80577E770004] clusterTime:{1678444645 1} documentKey:map[_id:ObjectID("5ee84144e6892f0a80577e77")] ns:map[coll:test db:himkt] operationType:update updateDescription:map[removedFields:[] updatedFields:map[t:t121]]]
	}

}
