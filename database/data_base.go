package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type User struct {
	id   int
	name string
	age  int
}

// 初始化数据库
func InitDB() (err error) {
	db, err = sql.Open("mysql", "abljiu:123456@tcp(localhost:3306)/my_data")
	if err != nil {
		return err
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		return err
	}

	return nil
}

func insertRow(newUser User) {
	//需要插入的sql语句，？表示占位参数
	sqlStr := "insert into user(name,age) values(?,?)"
	//把user结构体的name、age字段依次传给sqlStr的占位参数
	ret, err := db.Exec(sqlStr, newUser.name, newUser.age)
	if err != nil { //执行sql语句报错
		fmt.Println("插入失败,err", err)
		return
	}
	newID, err := ret.LastInsertId() //新插入数据的ID，默认为主键
	//rowsNumber, err:= ret.RowsAffected() //受影响的行数
	if err != nil {
		fmt.Println("获取id失败,err", err)
		return
	}
	fmt.Println("插入成功，id为：", newID)
}

// 添加新用户
func AddUser() bool {

	return false
}

// 删除用户
func DeleteUser() bool {
	return false
}

// 查询用户密码
func InquirePwd(Username string) string {
	pwd := ""
	return pwd
}

// 校验用户的密码
func CheckPwd(Username string) bool {
	return false
}

//
