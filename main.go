package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed resources/templates/*.html
var templatesFS embed.FS

//go:embed resources/static/*
var staticFS embed.FS

func main() {
	r := gin.Default()

	// 使用 gin 的嵌入式文件系统
	tmpl := template.Must(template.New("").ParseFS(templatesFS, "resources/templates/*.html"))
	r.SetHTMLTemplate(tmpl)

	// 使用嵌入的静态文件
	staticServer := http.FileServer(http.FS(staticFS))
	r.StaticFS("/static", http.FS(staticFS))
	r.GET("/static", func(c *gin.Context) {
		staticServer.ServeHTTP(c.Writer, c.Request)
	})

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
