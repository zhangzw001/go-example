package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

//| user  | CREATE TABLE `user` (
//`id` bigint(20) NOT NULL AUTO_INCREMENT,
//`name` varchar(45) DEFAULT '',
//`age` int(11) NOT NULL DEFAULT '0',
//`sex` tinyint(3) NOT NULL DEFAULT '0',
//`phone` varchar(45) NOT NULL DEFAULT '',
//PRIMARY KEY (`id`)
//) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8

type Users struct {
	id 	int64
	name string
	age int8
	sex int8
	phone string
}


const (
	driver = "mysql"
	dbUser = "root"
	dbPass = "boqii.123"
	dbHost = "172.16.76.194"
	dbPort = 3316
	dbName = "beetest"

)
func sqlQuery() {
	string := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",dbUser,dbPass,dbHost,dbPort,dbName)
	db, err := sql.Open(driver,string)
	if err != nil {
		log.Fatalf("connect mysql failed ! [ %s ]", err )
	}
	defer db.Close()
	fmt.Println("connect mysql success !")

	// 查询结果
	sql :="select * from user where id in (1,2,3,4)"
	rows , err := db.Query(sql)
	if err != nil {
		log.Fatalf("select faild [%s]", err )
	}
	var u Users
	for rows.Next() {
		rows.Scan(&u.id,&u.name,&u.age,&u.sex,&u.phone)
	}
	fmt.Println(u.phone)

}

func main() {
	sqlQuery()
}