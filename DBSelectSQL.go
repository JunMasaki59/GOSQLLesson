package main

import (
	"database/sql"
	//"errors"
	"fmt"
	"log"
	//"time"

	_ "github.com/go-sql-driver/mysql"
)

type Players71 struct {
	Id int
	Position string
	Name string 
	Height int
	Weight int
}

type Players70 struct {
	BirthYear int
	CountId int
}

func main() {
	db, err := sql.Open("mysql", "root:root_password@tcp(127.0.0.1:13306)/worldcup_db?parseTime=true&loc=Asia%2FTokyo")
	if err != nil {
		log.Fatalf("main sql.Open error err:%v", err)
	}
	defer db.Close()
	
	fmt.Println("----------------------------------------------------")
	fmt.Println("SQL練習問題-問71")
	fmt.Println("----------------------------------------------------")
	getRowsQ71(db)
	fmt.Println("----------------------------------------------------")
	fmt.Println("SQL練習問題-問70")
	fmt.Println("----------------------------------------------------")
	getRowsQ70(db)
	fmt.Println("----------------------------------------------------")
}

func getRowsQ71(db *sql.DB) {
	rows, err := db.Query("(SELECT id, position, name, height, weight FROM players WHERE height > 195) UNION ALL (SELECT id, position, name, height, weight FROM players WHERE weight > 95) ORDER BY id")
	if err != nil {
		log.Fatalf("getRows db.Query error err\n%v", err)
	}
	defer rows.Close()
	
	for rows.Next() {//rows.Next()で各レコードに対して操作が可能
		u := &Players71{}
//rows.Scan()で引数に渡したポインタにレコードの内容を読み込ますことができる
		err := rows.Scan(&u.Id, &u.Position, &u.Name, &u.Height, &u.Weight);
		if err != nil {
			log.Fatalf("getRows rows.Scan error err:%v", err)
		}
		fmt.Println(u)
	}

	err = rows.Err()
	if err != nil {
		log.Fatalf("getRows rows.Err error err:%v", err)
	}
}

func getRowsQ70(db *sql.DB) {
	rows, err := db.Query("(SELECT '1980', COUNT(id) FROM players WHERE birth BETWEEN '1980-01-01' AND '1980-12-31') UNION (SELECT '1981' AS '誕生年', COUNT(id) FROM players WHERE birth BETWEEN '1981-01-01' AND '1981-12-31')")
	if err != nil {
		log.Fatalf("getRows db.Query error err\n%v", err)
	}
	defer rows.Close()
	
	fmt.Println("&{誕生年 COUNT(Id)}")
	for rows.Next() {//rows.Next()で各レコードに対して操作が可能
		u := &Players70{}
//rows.Scan()で引数に渡したポインタにレコードの内容を読み込ますことができる
		err := rows.Scan(&u.BirthYear, &u.CountId);
		if err != nil {
			log.Fatalf("getRows rows.Scan error err:%v", err)
		}
		fmt.Println(u)
	}

	err = rows.Err()
	if err != nil {
		log.Fatalf("getRows rows.Err error err:%v", err)
	}
}