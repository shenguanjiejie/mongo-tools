/*
 * @Author: shenruijie shenruijie@sensetime.com
 * @Date: 2022-10-24 17:46:49
 * @LastEditors: shenruijie shenruijie@sensetime.com
 * @LastEditTime: 2022-10-24 18:35:08
 * @FilePath: /mongo-tools/mongodb_test.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package mongotools

import (
	"context"
	"testing"

	"github.com/shenguanjiejie/go-tools"
	"go.mongodb.org/mongo-driver/bson"
)

type Person struct {
	Name string `json:"name" bson:"name"`
	Age  int    `json:"age" bson:"age"`
}

func TestLog(t *testing.T) {
	client, _ := MongoClient("mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000")
	col := client.Database("test").Collection("testC")
	// col.InsertOne(context.Background(), &Person{"Another", 20})
	cursor, err := col.Find(context.Background(), bson.M{})
	if err != nil {
		tools.Logln(err)
		t.Error(err)
		return
	}

	LogCursor(cursor)
}

func TestSingleLog(t *testing.T) {
	client, _ := MongoClient("mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000")
	col := client.Database("test").Collection("testC")
	// col.InsertOne(context.Background(), &Person{"Another", 20})
	result := col.FindOne(context.Background(), bson.M{})

	LogSingleResult(result)
}
