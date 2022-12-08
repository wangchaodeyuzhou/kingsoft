package mysql

import (
	"fmt"
	"testing"
)

func TestMySqlDB_Get(t *testing.T) {
	dsn := "root:root@tcp(127.0.0.1:3306)/user?charset=utf8mb4&parseTime=True&loc=Local"
	db := NewMySqlDB(dsn)
	db.Insert("mysqlTest", "K", map[string]int{"ff": 123})
	get, err := db.Get("mysqlTest", "K")
	if err != nil {
		fmt.Println("err", err)
		return
	}

	fmt.Println(get)

	fmt.Println("=========================")

	// p := &Person{UserName: "adskka", Sex: "男", Email: "100435@qq.com"}
	// db.CreateTable("person")
	// db.InsertTable(p)
	db.SelectTableWithQuery()
	db.SelectTableQueryRaws()
	db.SelectTable()
}
