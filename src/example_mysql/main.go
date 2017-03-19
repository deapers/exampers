// A simaple example for mysql database.
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type FundInfo struct {
	FundCode string
	FundName string
}

func main() {
	db, err := sql.Open("mysql", "cta:cta@tcp(127.0.0.1:3306)/cta?charset=utf8")
	if err != nil {
		fmt.Println("连接数据库失败")
	}
	defer db.Close()
	var fundinfo []FundInfo = make([]FundInfo, 0)
	sqlStr := "select a.c_fundcode,a.c_fundname from tfundinfo a "
	fmt.Println(sqlStr)
	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Println(err)
	} else {
		for i := 0; rows.Next(); i++ {
			var u FundInfo
			rows.Scan(&u.FundCode, &u.FundName)
			fundinfo = append(fundinfo, u)
		}
		fmt.Println(fundinfo)
	}
}
