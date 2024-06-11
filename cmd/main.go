package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/static", "../resources/static")
	r.LoadHTMLGlob("../resources/templates/*.html") //加载templates目录下所有文件

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	r.GET("/hello", func(c *gin.Context) {
		c.HTML(http.StatusOK, "hello.html", gin.H{})
	})
	r.GET("/jsTest", func(c *gin.Context) {
		c.HTML(http.StatusOK, "jsTest.html", gin.H{})
	})
	r.GET("/test1", func(c *gin.Context) {
		c.HTML(http.StatusOK, "test1.html", gin.H{})
	})

	for range 3 {
		fmt.Print("n!!!!\n")
	}

	err := r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}
