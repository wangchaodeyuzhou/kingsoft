package mysql

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gogf/gf/v2/util/gconv"
)

var DataNotExistErr = errors.New("data not exist")

type MySqlDB struct {
	addr   string
	client *sql.DB
}

func NewMySqlDB(addr string) *MySqlDB {
	db, err := sql.Open("mysql", addr)
	if err != nil {
		panic(err)
	}

	result := &MySqlDB{addr: addr, client: db}

	result.createTable()
	return result
}

func (m *MySqlDB) createTable() {
	sqlStr := "CREATE TABLE IF NOT EXISTS " + "mysqlTest" + " ( `k` varchar(127)  NOT NULL, `data` MediumBlob," +
		"PRIMARY KEY (`k`) )  ENGINE=MyIsam DEFAULT CHARSET=utf8mb4"
	fmt.Println("sqlStr : ", sqlStr)
	_, err := m.client.Exec(sqlStr)
	if err != nil {
		m.Close()
		panic("mysql create table failed")
	}
}

func (m *MySqlDB) Get(tableName string, key any) (any, error) {
	indexKey := gconv.String(key)
	sqlStr := "select data from " + tableName + " where k = ?"
	fmt.Println("get sqlStr : ", sqlStr)
	var result []byte
	err := m.client.QueryRow(sqlStr, indexKey).Scan(&result)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, DataNotExistErr
		}

		fmt.Println("mysql get data error", "key", indexKey, "err", err)
	}

	originData := make(map[string]any)
	err = json.Unmarshal(result, &originData)
	if err != nil {
		fmt.Println("json unmarshal err", err)
		return nil, err
	}

	return originData, nil
}

func (m *MySqlDB) Insert(tableName string, key any, data any) {
	indexKey := gconv.String(key)
	binData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("mysql marshal data err", err)
		return
	}

	sqlStr := "REPLACE INTO " + tableName + "(k,data) VALUES(?,?)"
	fmt.Println("insert sqlStr: ", sqlStr)
	_, err = m.client.Exec(sqlStr, indexKey, binData)
	if err != nil {
		fmt.Println("mysql insert data err", err)
	}
}

func (m *MySqlDB) Delete(tableName string, key any) {
	indexKey := gconv.String(key)
	sqlStr := "delete from " + tableName + "where k = ?"
	fmt.Println("delete sqlStr : ", sqlStr)

	_, err := m.client.Exec(sqlStr, indexKey)
	if err != nil {
		fmt.Println("delete data err", err)
		return
	}
}

func (m *MySqlDB) Close() {
	err := m.client.Close()
	if err != nil {
		fmt.Println("close mysql fail", "err", err.Error())
		return
	}
}
