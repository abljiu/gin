package main

import (
	// "fmt"
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	// "golang.org/x/tools/go/analysis/passes/nilfunc"
)

func main() {
	//创建server
	r := gin.Default()
	db, err := sql.Open("mysql", "username:password@tcp(10.30.0.208/:3306)/dbname")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	r.LoadHTMLGlob("templates/*")

	r.GET("/main.go", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", nil)
	})
	
	// 定义一个路由，当请求 POST /login 时，接收 username 和 password，并返回接收到的数据
	r.POST("/login", func(c *gin.Context) {
		// 使用 c.PostForm() 方法获取表单提交的数据
		username := c.PostForm("username")
		password := c.PostForm("password")

		
		// 返回接收到的数据
		c.HTML(http.StatusOK,"log_res.html", gin.H{
			"username": username,
			"password": password,
		})
	})

	r.Run("10.30.0.208:8000")
}
