package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/i-coder-robot/mock-food"
)

func main() {
	fmt.Println("aaa")
	r := gin.Default()
	r.GET("/index",handler.IndexHandler)
}
