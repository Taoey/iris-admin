package test

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"testing"
)

var gdb *sql.DB

func init() {
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.44.69:3306)/test")
	gdb = db
	if err != nil {
		fmt.Println(err)
	}
}

func TestConnect(t *testing.T) {
	db, err := sql.Open("mysql", "root:123456@tcp(192.168.44.69:3306)/test")
	gdb = db
	fmt.Println(db, err)
}

type User struct {
	Id   int
	Name string
	Age  int
}

func TestQuery(t *testing.T) {
	sql := "SELECT * FROM users WHERE age = ?"

	rows, err := gdb.Query(sql, 12)

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var user User
		strings, _ := rows.Columns()
		fmt.Println(strings)
		if err := rows.Scan(&user.Id, &user.Name, &user.Age); err != nil {
			log.Fatal(err)
		}
		fmt.Println(user)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
