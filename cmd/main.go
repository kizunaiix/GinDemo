package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("../resources/templates/*.html")
	r.Static("static", "../resources/static")

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
		fmt.Print("go 1.22!!!\n")
	}

	err := r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}
