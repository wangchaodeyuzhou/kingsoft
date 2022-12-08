package mogodb

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type MongoDB struct {
	addr     string
	client   *mongo.Client
	dataBase *mongo.Database
}

func NewMongoDB(addr string) *MongoDB {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(addr))
	if err != nil {
		panic("connect mongo fail")
	}

	return &MongoDB{
		addr:     addr,
		client:   client,
		dataBase: client.Database("game"),
	}
}

func (m *MongoDB) Get(tableName string, key any) (any, error) {
	indexKey := gconv.String(key)
	result := m.dataBase.Collection(tableName).FindOne(context.Background(), bson.D{{"_id", indexKey}})
	data := make(map[string]any)
	err := result.Decode(data)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		fmt.Println("mongo get data error", err)
		return nil, err
	}
	return data, nil
}

// Insert 不存在就插入， 存在就更新
func (m *MongoDB) Insert(tableName string, key, data any) {
	indexKey := gconv.String(key)
	tempData := data.(map[string]any)
	_, err := m.dataBase.Collection(tableName).UpdateOne(context.Background(), bson.D{{"_id", indexKey}}, bson.D{{"$set", tempData}}, options.Update().SetUpsert(true))
	if err != nil {
		fmt.Println("err := ", err)
		return
	}

	fmt.Println("insert success")
}

func (m *MongoDB) Delete(tableName string, key any) {
	indexKey := gconv.String(key)
	res, err := m.dataBase.Collection(tableName).DeleteOne(context.Background(), bson.D{{"_id", indexKey}})
	if err != nil {
		fmt.Println("mongo delete err", err)
		return
	}

	fmt.Println("res := ", res.DeletedCount)
	fmt.Println("delete data success")
}

func (m *MongoDB) Close() {
	err := m.client.Disconnect(context.Background())
	if err != nil {
		fmt.Println("client err", err)
		return
	}
}

func (m *MongoDB) SelectManyMongoDB(tableName string) (any, error) {
	cur, err := m.dataBase.Collection(tableName).Find(context.Background(), bson.D{{}}, options.Find().SetLimit(5))
	if err != nil {
		fmt.Println("err", err)
		return nil, err
	}

	var res []map[string]any
	for cur.Next(context.TODO()) {
		var result map[string]any
		err := cur.Decode(&result)
		if err != nil {
			fmt.Println("decode err :=", err)
			return nil, err
		}
		res = append(res, result)
	}
	cur.Close(context.TODO())
	fmt.Println("res := ", res)
	return res, err
}
