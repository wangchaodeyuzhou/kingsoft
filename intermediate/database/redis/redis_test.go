package redis

import (
	"fmt"
	"testing"
)

func TestRedis(t *testing.T) {
	db := NewRedisDB("localhost:6379")
	db.Insert("wcdyz", "ww", map[string]any{"kkg": "vv"})
	get, err := db.Get("wcdyz", "ww")
	if err != nil {
		fmt.Println("err", err)
		return
	}

	defer db.Close()
	fmt.Println("get data : ", get)
	db.Delete("wcdyz", "www")
	fmt.Println("=======")
	db.SetKey()
}
