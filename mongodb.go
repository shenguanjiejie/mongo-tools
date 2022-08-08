package mongotools

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/shenguanjiejie/go-tools"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongoPipeline 解析str pipline
func MongoPipeline(str string) mongo.Pipeline {
	var pipeline = make(mongo.Pipeline, 0)
	str = strings.TrimSpace(str)
	if strings.Index(str, "[") != 0 {
		var cell bson.D
		bson.UnmarshalExtJSON([]byte(str), false, &cell)
		pipeline = append(pipeline, cell)
	} else {
		bson.UnmarshalExtJSON([]byte(str), false, &pipeline)
	}
	return pipeline
}

// LogCursor 打印cursor信息
func LogCursor(cursor *mongo.Cursor) error {
	var results []interface{}
	err := cursor.All(context.Background(), &results)
	if err != nil {
		tools.Slogln(err)
		return err
	}
	jsonPipe, err := json.MarshalIndent(results, "", "    ")
	if err != nil {
		tools.Slogln(err)
		return err
	}

	tools.Slogln(string(jsonPipe))
	return nil
}

// MongoClient 创建client并连接
func MongoClient(URI string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(URI)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		tools.Slogln(err)
		return nil, err
	}
	if err = client.Ping(context.Background(), readpref.Primary()); err != nil {
		tools.Slogln("mongo ping err:%s", err.Error())
		return nil, err
	}

	return client, nil
}
