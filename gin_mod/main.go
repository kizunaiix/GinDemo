package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"this": "hello world",
		})
	})
	err := r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}
