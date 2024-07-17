package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// 😀🆒🐉👌
func main() {
	r := gin.Default()

	//CORS策略
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:80", "http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
	}))

	r.LoadHTMLGlob("../resources/templates/*.html")
	r.Static("static", "../resources/static")

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "{hello.WLAN}")
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

	//post： /api/login
	r.POST("/api/login", func(ctx *gin.Context) {
		var request struct {
			Content string `json:"content"`
		}

		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if request.Content == "hello api" {
			ctx.JSON(http.StatusOK, gin.H{"content": "hello client!"})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"content": "who r u?"})
		}
	})

	err := r.Run("0.0.0.0:8080")
	if err != nil {
		log.Fatalln(err)
	}
}
