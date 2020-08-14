package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)


type Users struct {
	Host 	string
	User 	string
}

func sqlQuery() {
	db, err := sql.Open("mysql","bq_ddl_user:cHkl8eWyndA1@tcp(172.16.76.194:13306)/mysql?charset=utf8")
	if err != nil {
		log.Fatalf("connect mysql failed ! [ %s ]", err )
	}
	defer db.Close()
	fmt.Println("connect mysql success !")

	rows , err := db.Query("select Host,User from user limit 1")
	var u Users
	if err != nil {
		log.Fatalf("select faild [%s]", err )
	}
	fmt.Println(rows)
	for rows.Next() {
		e := rows.Scan(&u.Host,&u.User)
		fmt.Println(e)
	}


}

func main() {
	sqlQuery()
}