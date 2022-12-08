package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/gogf/gf/v2/util/gconv"
)

var DataNotExistErr = errors.New("data not exist")

type RedisDB struct {
	addr   string
	client *redis.Client
}

func NewRedisDB(addr string) *RedisDB {
	return &RedisDB{
		addr:   addr,
		client: redis.NewClient(&redis.Options{Addr: addr, DB: 1, Password: ""}),
	}
}

func (r *RedisDB) Insert(tableName string, key any, data any) {
	redisKey := gconv.String(key)
	redisData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("err", err)
		return
	}

	cmd := r.client.HSet(context.Background(), redisKey, tableName, redisData)
	if cmd.Err() != nil {
		fmt.Println("err", err)
		return
	}

	fmt.Println("insert success")
}

func (r *RedisDB) Get(tableName string, key any) (any, error) {
	redisKey := gconv.String(key)
	result, err := r.client.HGet(context.Background(), redisKey, tableName).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, DataNotExistErr
		}
		return nil, err
	}

	originData := make(map[string]any)
	err = json.Unmarshal(result, &originData)
	if err != nil {
		fmt.Println("unmarshal data error", err)
		return nil, err
	}

	return originData, nil
}

func (r *RedisDB) Delete(tableName string, key any) {
	redisKey := gconv.String(key)
	cmd := r.client.HDel(context.Background(), redisKey, tableName)
	if cmd.Err() != nil {
		fmt.Println("delete data error", cmd.Err())
	}
	fmt.Println("delete success")
}

func (r *RedisDB) Close() {
	err := r.client.Close()
	if err != nil {
		fmt.Println("err", err)
		return
	}
}

func (r *RedisDB) SetKey() {
	cmd := r.client.Do(context.Background(), "MSet", "abc", 100, "efg", 300)
	fmt.Println("cmd ", cmd)

	do := r.client.Do(context.Background(), "MGet", "abc", "efg")
	fmt.Println("do ", do)

	c := r.client.Do(context.Background(), "LPush", "book_list", "abc", "ceg", 300)
	fmt.Println("c :", c)

	c2 := r.client.Do(context.Background(), "LPop", "book_list")
	fmt.Println("c2 := ", c2)

	c3 := r.client.Do(context.Background(), "HSet", "book", "abc", 100)
	fmt.Println("c3 :=", c3)

	c4 := r.client.Do(context.Background(), "HGet", "book", "abc")
	fmt.Println("c4 := ", c4)
}
