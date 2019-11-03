package init

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" //加载mysql
	"github.com/jinzhu/gorm"
)

var Eloquent *gorm.DB

// mysql 数据库初始化
func init() {
	var err error
	// root:root ==>> 用户名:密码
	Eloquent, err = gorm.Open("mysql", "root:root@tcp(119.23.105.191:3306)/go_leaning?charset=utf8")

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	if Eloquent.Error != nil {
		fmt.Printf("database error %v", Eloquent.Error)
	}
}

// Postgres 数据库初始化
//func initPostgres() {
//	var err error
//	Eloquent, err = gorm.Open("postgres", "host:www.gochanghai.com dbname=test port=5432  user=root password=root  charset=utf8 parseTime=True loc=Local timeout=10ms")
//
//	if err != nil {
//		fmt.Printf("postgres connect error %v", err)
//	}
//
//	if Eloquent.Error != nil {
//		fmt.Printf("database error %v", Eloquent.Error)
//	}
//}
