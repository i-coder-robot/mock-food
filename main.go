package main

import (
	"github.com/gin-gonic/gin"
	"github.com/i-coder-robot/mock-food/handler"
)

func main() {
	r := gin.Default()
	r.GET("/index",handler.IndexHandler)
	r.GET("/nav",handler.NavHandler)
	r.GET("subNav",handler.SubNavHandler)
	r.GET("/food",handler.GoodFoodHandler)
	r.GET("/restaurant",handler.GoodRestaurantHandler)
	r.GET("/me",handler.MeHandler)
	r.GET("/image",handler.ImageHandler)
	r.Run()
}
