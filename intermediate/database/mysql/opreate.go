package mysql

import (
	"fmt"
)

type Person struct {
	UserID   int    `db:"user_id"`
	UserName string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type Place struct {
	Country string `db:"country"`
	City    string `db:"city"`
	TelCode int    `db:"telcode"`
}

func (m *MySqlDB) CreateTable(tableName string) {
	sqlS := "DROP TABLE IF EXISTS " + tableName + ";"
	res, err := m.client.Exec(sqlS)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println("res : ", res)
	sqlStr := "CREATE TABLE " + tableName + " ( `user_id` int(12) NOT NULL AUTO_INCREMENT, `username` varchar(127)  NOT NULL, `sex` varchar(11), `email` varchar(127), PRIMARY KEY (`user_id`))  ENGINE=MyIsam DEFAULT CHARSET=utf8mb4"
	fmt.Println("sql Str: ", sqlStr)

	result, err := m.client.Exec(sqlStr)
	if err != nil {
		fmt.Println("err", err)
		m.Close()
		return
	}

	fmt.Println("result: ", result)
}

func (m *MySqlDB) InsertTable(p *Person) {
	sqlStr := "INSERT INTO person(username,sex, email) VALUES(?, ?, ?)"
	result, err := m.client.Exec(sqlStr, p.UserName, p.Sex, p.Email)
	if err != nil {
		fmt.Println("err", err)
		return
	}

	affected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("err", err)
		return
	}

	fmt.Println("affected: ", affected)

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("err", err)
		return
	}

	fmt.Println("insert success: id", id)

}

func (m *MySqlDB) SelectTableWithQuery() {
	sqlStr := "SELECT user_id, username, sex, email FROM person"
	rows, err := m.client.Query(sqlStr)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		temp := new(Person)
		err = rows.Scan(&temp.UserID, &temp.UserName, &temp.Sex, &temp.Email)
		if err != nil {
			fmt.Println("err", err)
			return
		}
		fmt.Println("person ", temp)
	}
}

func (m *MySqlDB) SelectTableQueryRaws() {
	sqlStr := "SELECT user_id, username, sex, email FROM person"
	fmt.Println("sqlStr: ", sqlStr)

	p := Person{}
	err := m.client.QueryRow(sqlStr).Scan(&p.UserID, &p.UserName, &p.Sex, &p.Email)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println(p)
}

func (m *MySqlDB) SelectTable() {
	sqlStr := "SELECT user_id, username, sex, email FROM person"
	fmt.Println("sqlStr: ", sqlStr)

	result, err := m.client.Exec(sqlStr)
	if err != nil {
		fmt.Println("err", err)
		return
	}

	affected, err := result.LastInsertId()
	if err != nil {
		fmt.Println("err", err)
		return
	}

	fmt.Println(affected)
}
