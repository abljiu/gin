package Database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

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
