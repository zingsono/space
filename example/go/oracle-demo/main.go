package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "gopkg.in/goracle.v2"
)

type Tb struct {
	ID    string
	VALUE string
}

func QueryForMaps(db *sql.DB, query string, args ...interface{}) ([]map[string]interface{}, error) {
	rows, err := db.Query(query, args...)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	columns, _ := rows.Columns()
	// 临时存储每行数据
	cache := make([]interface{}, len(columns))
	// 为每一列初始化一个指针
	for index, _ := range cache {
		var a interface{}
		cache[index] = &a
	}
	// 返回的切片
	var list []map[string]interface{}
	for rows.Next() {
		_ = rows.Scan(cache...)
		item := make(map[string]interface{})
		for i, data := range cache {
			// //取实际类型
			item[columns[i]] = *data.(*interface{})
		}
		list = append(list, item)
	}
	return list, nil
}

// 读写oracle数据库表  go get gopkg.in/goracle.v2
func main() {
	drivers := sql.Drivers()
	fmt.Println(drivers)

	dataSourceName := "oracle://unionlive:unionlive@proxy.unionlive.com:1521/orcl"
	db, err := sql.Open("goracle", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	rs, err := QueryForMaps(db, "select * from tb")
	byte, err := json.Marshal(rs)
	fmt.Println(string(byte))

}
