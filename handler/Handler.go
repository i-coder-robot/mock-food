package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/i-coder-robot/mock-food/model"
	"io/ioutil"
	"net/http"
	"os"
)
const domain = "http://localhost:8080"
func IndexHandler(c *gin.Context)          {}
func GoodFoodHandler(c *gin.Context)       {}
func GoodRestaurantHandler(c *gin.Context) {}
func MeHandler(c *gin.Context)             {}
func NavHandler(c *gin.Context) {
	items := []model.IndexNavItem{
		model.IndexNavItem{
			Id:        "26cd5373-7709-468e-aa89-34ddcac20b3d",
			Src:       domain+"/image?imageName=meishidatu",
			Title:     "美食",
		}, model.IndexNavItem{
			Id:        "ab36ffb0-4c6c-4d6f-8ed5-20840896d496",
			Src:        domain+"/image?imageName=waimaidatu",
			Title:     "美团外卖",
		}, model.IndexNavItem{
			Id:        "547b2627-b107-45ef-b144-257a1f8bbc8d",
			Src:        domain+"/image?imageName=xiuxianyuledatu",
			Title:     "休闲娱乐",
		}, model.IndexNavItem{
			Id:        "b26c6db6-3f6c-4389-b2ba-866792ffe205",
			Src:        domain+"/image?imageName=xiuxianyuledatu",
			Title:     "酒店",
		}, model.IndexNavItem{
			Id:        "b5e2c25d-da0d-41e8-a42f-1c8ad1f95430",
			Src:        domain+"/image?imageName=hotel",
			Title:     "演出",
		}, model.IndexNavItem{
			Id:        "c5fbc033-a779-402a-a06c-cbb4c9a2e590",
			Src:       domain+"/image?imageName=lirenmeifadatu",
			Title:     "丽人",
		}, model.IndexNavItem{
			Id:        "e83533d7-de2d-4c52-bc9a-769a5ee96030",
			Src:       domain+"/image?imageName=jingdianmenpiaodatu",
			Title:     "周边",
		}, model.IndexNavItem{
			Id:        "b465c601-bc83-4f5e-bc91-031ea5b86f0b",
			Src:        domain+"/image?imageName=travel",
			Title:     "景点/门票",
		}, model.IndexNavItem{
			Id:        "de43b517-dbe2-4329-8eb8-1f28d3cd9f6e",
			Src:        domain+"/image?imageName=KTVdatu",
			Title:     "KTV",
		}, model.IndexNavItem{
			Id:        "3b8f0fe4-09e8-4c0d-b2d5-2093ea568dda",
			Src:        domain+"/image?imageName=qinzi",
			Title:     "亲子",
		},
	};
	c.JSON(http.StatusOK, gin.H{
		"items": items,
	})
}
func ImageHandler(c *gin.Context) {
	imageName := c.Query("imageName")
	dir, _ := os.Getwd()
	println(dir)
	//exPath := filepath.Dir(dir)
	//println(exPath)
	file, ok := ioutil.ReadFile(dir+"/static/IndexNav/"+imageName+".png")
	if ok!=nil {
		fmt.Println(ok.Error())
	}
	c.Writer.WriteString(string(file))
}
