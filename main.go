package main

import (
	"github.com/gin-gonic/gin"
	"github.com/i-coder-robot/mock-food/handler"
)

func main() {
	r := gin.Default()
	r.GET("/index",handler.IndexHandler)
	r.GET("/discountLeft",handler.DiscountLeftHandler)
	r.GET("/discountRight",handler.DiscountRightHandler)

	r.GET("/image",handler.ImageHandler)


	r.GET("/nav",handler.NavHandler)
	r.GET("/subNav",handler.SubNavHandler)
	r.GET("/team",handler.TeamHandler)
	r.GET("/rush",handler.RushHandler)

	r.GET("/guess",handler.GuessHandler)

	r.GET("/restaurantNav",handler.GoodRestaurantNavHandler)
	r.GET("/restaurantBillBoard",handler.GoodRestaurantBillBoardHandler)
	r.GET("/restaurantTabItem",handler.GoodRestaurantTabItemHandler)

	r.GET("/me",handler.MeHandler)

	r.GET("/dySameCity",handler.SameCityHandler)

	r.Run()
}
