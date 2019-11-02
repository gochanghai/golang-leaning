package init

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" //加载mysql
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

// mysql 数据库初始化
func init() {
	var err error
	// root:root ==>> 用户名:密码
	db, err = gorm.Open("mysql", "root:root@tcp(www.changhai.com:3306)/test?charset=utf8&parseTime=True&loc=Local&timeout=10ms")

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	if db.Error != nil {
		fmt.Printf("database error %v", db.Error)
	}
}

// mysql 数据库初始化
func initPostgres() {
	var err error
	db, err = gorm.Open("postgres", "host:www.gochanghai.com dbname=test port=5432  user=root password=root  charset=utf8 parseTime=True loc=Local timeout=10ms")

	if err != nil {
		fmt.Printf("postgres connect error %v", err)
	}

	if db.Error != nil {
		fmt.Printf("database error %v", db.Error)
	}
}
