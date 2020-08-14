package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func sqlQuote(x interface{}) string {
	if x == nil {
		return "NULL"
	} else if _, ok := x.(int); ok {
		return fmt.Sprintf("%d", x)
	} else if _, ok := x.(uint); ok {
		return fmt.Sprintf("%d", x)
	} else if b, ok := x.(bool); ok {
		if b {
			return "TRUE"
		}
		return "FALSE"
	} else if s, ok := x.(string); ok {
		return s
		//return sqlQuoteString(s) // (not shown)
	} else {
		panic(fmt.Sprintf("unexpected type %T: %v", x, x))
	}
}

func sqlQuote2(x interface{}) string {
	switch x := x.(type) {
	case nil :
		return "NULL"
	case int,uint:
		return fmt.Sprintf("%d",x)
	case bool :
		if x { return "TRUE"}
		return "FALSE"
	case string :
		return x
	default:
		panic(fmt.Sprintf("unexpected type %T: %v", x, x))
	}

	//return ""
}

type Users struct {
	Username 	string `json="username"`
	Uid 		string
	Email 		string
}

func sqlQuery() {
	db, err := sql.Open("mysql","bq_ddl_user:cHkl8eWyndA1@tcp(172.16.76.194:13306)/hzkj_zh?charset=utf8")
	if err != nil {
		log.Fatalf("connect mysql failed ! [ %s ]", err )
	}
	fmt.Println("connect mysql success !")

	defer db.Close()

	rows , err := db.Query("select username,uid,email from boqii_users limit 1")
	var u Users
	if err != nil {
		log.Fatalf("select faild [%s]", err )
	}
	for rows.Next() {
		e := rows.Scan(u.Username,u.Uid,u.Email)
		fmt.Println(e)
	}


}

func main() {
	sqlQuery()
}