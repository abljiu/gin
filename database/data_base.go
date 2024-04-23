package database

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *sql.DB

type User struct {
	Id   int
	Name string
	Age  int
}

// InitDB 初始化数据库
func InitDB() (err error) {

	dsn := "abljiu:123456@tcp(localhost:3306)/my_data?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		return err
	}
	return nil
}

func insertRow(newUser User) {
	//需要插入的sql语句，？表示占位参数
	sqlStr := "insert into user(name,age) values(?,?)"
	//把user结构体的name、age字段依次传给sqlStr的占位参数
	ret, err := db.Exec(sqlStr, newUser.Name, newUser.Age)
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

// AddUser 添加新用户
func AddUser() bool {

	return false
}

// DeleteUser 删除用户
func DeleteUser() bool {
	return false
}

// InquirePwd 查询用户密码
func InquirePwd(Username string) string {
	pwd := ""
	return pwd
}

// CheckPwd 校验用户的密码
func CheckPwd(Username string) bool {

	return false
}
