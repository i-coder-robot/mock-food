package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHandler(c *gin.Context){}
func GoodFoodHandler(c *gin.Context){}
func GoodRestaurantHandler(c *gin.Context){}
func MeHandler(c *gin.Context){}
func NavHandler(c *gin.Context){
	items []model.IndexNavItem;
	c.JSON(http.StatusOK,gin.H{

	})
}
