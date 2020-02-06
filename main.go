package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/i-coder-robot/mock-food/handler"
)

func main() {
	fmt.Println("aaa")
	r := gin.Default()
	r.GET("/index",handler.IndexHandler)
	r.GET("/index",handler.IndexHandler)
	r.GET("/index",handler.IndexHandler)
	r.GET("/me",handler.IndexHandler)
}
