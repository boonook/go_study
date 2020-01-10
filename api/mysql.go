package api

import (
	"database/sql"
	"fmt"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id       string
	Username string
	Password string
}

func Mysql() map[string]*User {
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/ZHDJ?charset=utf8")
	if err != nil {
		beego.Error("连接数据库出错", err)
		return nil
	} else {
		beego.Info("连接数据库成功")
	}
	rows, err := db.Query("select id,userName from user")
	type UserInfo struct {
		id       int
		userName string
	}
	var u UserInfo
	fmt.Println("rows--------", rows)
	for rows.Next() {
		err = rows.Scan(&u.id, &u.userName)
		fmt.Println(u)
	}
	return nil
}
