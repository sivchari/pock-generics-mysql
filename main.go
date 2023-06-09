package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	var err error
	db, err = sqlx.Open("mysql", "docker:docker@tcp(localhost:3306)/general")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	userQuery, todoQuery := "SELECT * FROM users", "SELECT * FROM todos"
	users, err := Select[User](userQuery)
	if err != nil {
		panic(err.Error())
	}
	for _, user := range users {
		fmt.Println("user is ", user)
	}
	todos, err := Select[Todo](todoQuery)
	if err != nil {
		panic(err.Error())
	}
	for _, todo := range todos {
		fmt.Println("todo is ", todo)
	}
}

type User struct {
	ID   int    `db:"id"`
	Name string `db:"username"`
}

type Todo struct {
	ID     int    `db:"id"`
	UserID int    `db:"user_id"`
	Title  string `db:"title"`
}

type queryconstraints interface {
	User | Todo
}

type query[T queryconstraints] interface {
	Select(query string) ([]*T, error)
}

var db *sqlx.DB

func Select[T queryconstraints](query string) ([]*T, error) {
	rows, err := db.Queryx(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []*T
	for rows.Next() {
		var row T
		if err := rows.StructScan(&row); err != nil {
			return nil, err
		}
		result = append(result, &row)
	}
	return result, nil
}
