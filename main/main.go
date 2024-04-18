package main

import (
	"gin/database"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//创建server
	r := gin.Default()
	r.LoadHTMLGlob("./templates/*")

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "start.html", nil)
	})
	//初始化数据库
	err := database.InitDB()
	if err != nil {
		log.Fatal("初始化数据库失败:", err)
	}

	//请求 POST /register 时，接受数据
	r.POST("/start", func(c *gin.Context) {

		//
	})

	//请求 POST /login 时，接收数据
	r.POST("/login", func(c *gin.Context) {
		// 使用 c.PostForm() 方法获取表单提交的数据
		// username := c.PostForm("username")
		// password := c.PostForm("password")

		//校验用户信息

		//验证成功 进入主界面

		//验证失败 继续留在登陆界面
	})
	r.Run(":8000")
}
