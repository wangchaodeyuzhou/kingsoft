package mogodb

import (
	"fmt"
	"testing"
)

func TestMongoDB(t *testing.T) {
	dsn := "mongodb://localhost:27017"
	db := NewMongoDB(dsn)

	db.Insert("wrc", "LL", map[string]any{"ccdsd": "cdsb"})
	get, err := db.Get("wrc", "LL")
	if err != nil {
		fmt.Println("err := ", err)
		return
	}

	fmt.Println("get := ", get)

	db.SelectManyMongoDB("wrc")
	// db.Delete("wrc", "kk")
}
