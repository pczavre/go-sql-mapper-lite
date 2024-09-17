package main

import (
	"database/sql"
	"fmt"
	"log"

	mapper "github.com/pczavre/go-sql-mapper-lite/pkg/sqlMapper"
)

type UserRow struct {
	Id    sql.NullInt16
	Name  sql.NullString
	Email sql.NullString
	DOB   sql.NullString
}

func main() {

	db, err := sql.Open("mysql", "connectionString")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	result, err := db.Query("select * from users")

	if err != nil {
		log.Fatal(err)
	}
	defer result.Close()

	var users []*UserRow

	for result.Next() {
		var row *UserRow = &UserRow{}

		if err := result.Scan(mapper.GetPointerArray(row)...); err != nil {
			log.Fatal(err)
		}

		users = append(users, row)
	}

	fmt.Println(users)
}
