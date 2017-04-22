package penguin

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	id      int
	mac     string
	version string
)

func postgresqltest() {
	db, err := sql.Open("postgres", "user=rails password=rails dbname=rails sslmode=disable")
	if err != nil {
		fmt.Println("连接数据库报错" + err.Error())
		return
	}
	sql := `select id,mac,version from public.devices`
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println("执行sql报错:" + err.Error())
		return
	}
	for rows.Next() {
		err := rows.Scan(&id, &mac, &version)
		if err != nil {
			println("查询出错", err.Error())
		}
		println(" id:", id, " mac:", mac, "version:", version, "\r\n")
	}
	fmt.Println(rows)

}
