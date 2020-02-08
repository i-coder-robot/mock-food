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
			Id:    "26cd5373-7709-468e-aa89-34ddcac20b3d",
			Src:   domain + "/image?imageName=meishidatu",
			Title: "美食",
		}, model.IndexNavItem{
			Id:    "ab36ffb0-4c6c-4d6f-8ed5-20840896d496",
			Src:   domain + "/image?imageName=waimaidatu",
			Title: "美团外卖",
		}, model.IndexNavItem{
			Id:    "547b2627-b107-45ef-b144-257a1f8bbc8d",
			Src:   domain + "/image?imageName=xiuxianyuledatu",
			Title: "休闲娱乐",
		}, model.IndexNavItem{
			Id:    "b26c6db6-3f6c-4389-b2ba-866792ffe205",
			Src:   domain + "/image?imageName=xiuxianyuledatu",
			Title: "酒店",
		}, model.IndexNavItem{
			Id:    "b5e2c25d-da0d-41e8-a42f-1c8ad1f95430",
			Src:   domain + "/image?imageName=hotel",
			Title: "演出",
		}, model.IndexNavItem{
			Id:    "c5fbc033-a779-402a-a06c-cbb4c9a2e590",
			Src:   domain + "/image?imageName=lirenmeifadatu",
			Title: "丽人",
		}, model.IndexNavItem{
			Id:    "e83533d7-de2d-4c52-bc9a-769a5ee96030",
			Src:   domain + "/image?imageName=jingdianmenpiaodatu",
			Title: "周边",
		}, model.IndexNavItem{
			Id:    "b465c601-bc83-4f5e-bc91-031ea5b86f0b",
			Src:   domain + "/image?imageName=travel",
			Title: "景点/门票",
		}, model.IndexNavItem{
			Id:    "de43b517-dbe2-4329-8eb8-1f28d3cd9f6e",
			Src:   domain + "/image?imageName=KTVdatu",
			Title: "KTV",
		}, model.IndexNavItem{
			Id:    "3b8f0fe4-09e8-4c0d-b2d5-2093ea568dda",
			Src:   domain + "/image?imageName=qinzi",
			Title: "亲子",
		},
	};
	c.JSON(http.StatusOK, gin.H{
		"items": items,
	})
}

func SubNavHandler(c *gin.Context) {
	items := []model.IndexNavItem{
		model.IndexNavItem{
			Id:    "b9a936d5-8cfd-4b9f-8057-abc14a8b99d7",
			Src:   domain + "/image?imageName=miaosha",
			Title: "限时秒杀",
		}, model.IndexNavItem{
			Id:    "49a8d7be-bcb2-45ad-b333-5489823af088",
			Src:   domain + "/image?imageName=huoguo",
			Title: "火锅",
		}, model.IndexNavItem{
			Id:    "703909c0-07df-4d63-9b16-cdad3d5feb24",
			Src:   domain + "/image?imageName=minsu",
			Title: "民宿",
		},
		model.IndexNavItem{
			Id:    "fab30bcf-842a-419a-b6be-ae2ea74ee49b",
			Src:   domain + "/image?imageName=hair",
			Title: "美发",
		}, model.IndexNavItem{
			Id:    "6f78b854-7473-4656-aaf7-9a92ef6aa95f",
			Src:   domain + "/image?imageName=health",
			Title: "按摩/足疗",
		}, model.IndexNavItem{
			Id:    "8a5c1877-deac-4ff1-bab1-8e7138ddd3ed",
			Src:   domain + "/image?imageName=bath",
			Title: "洗浴汗蒸",
		}, model.IndexNavItem{
			Id:    "30286152-1c1f-49af-96ff-33c8fe361fdf",
			Src:   domain + "/image?imageName=sports",
			Title: "运动健身",
		}, model.IndexNavItem{
			Id:    "5efbad00-ebcd-4c41-aaea-65c023b5f45e",
			Src:   domain + "/image?imageName=edu",
			Title: "教育培训",
		}, model.IndexNavItem{
			Id:    "5de3967a-0fbe-4d68-8d54-d2e49ec2aec8",
			Src:   domain + "/image?imageName=vote",
			Title: "聚餐投票",
		}, model.IndexNavItem{
			Id:    "c7464890-c961-4af7-954d-b6845e0b4364",
			Src:   domain + "/image?imageName=all-small",
			Title: "全部",
		},
	};
	c.JSON(http.StatusOK, gin.H{
		"items": items,
	})
}

func TeamHandler(c *gin.Context) {

	item := model.TeamModel{
		TeamHead:      "超",
		TeamHead1:     "值",
		TeamHead2:     "拼",
		TeamHead3:     "团",
		TeamImg:       "http://localhost:8080/image?imageName=niuroumian",
		FoodName:      "牛肉面",
		TeamIcon:      "http://localhost:8080/image?imageName=icon",
		TeamHotelName: "丰瑞餐厅",
		TeamPersons:   "2人团",
		TeamPrice:     "5.4",
		Price:         "9",
	}
	c.JSON(http.StatusOK, gin.H{"item": item})
}

func RushHandler(c *gin.Context) {
	item := model.TeamModel{
		TeamHead:      "限",
		TeamHead1:     "时",
		TeamHead2:     "秒",
		TeamHead3:     "杀",
		TeamImg:       "http://localhost:8080/image?imageName=rush",
		FoodName:      "肥牛丼套餐1份",
		TeamIcon:      "http://localhost:8080/image?imageName=icon",
		TeamHotelName: "熊吞大碗丼",
		TeamPersons:   "2人团",
		TeamPrice:     "32.8",
		Price:         "68",
		Distance:      "<100m",
	}
	c.JSON(http.StatusOK, gin.H{"item": item})
}

func GuessHandler(c *gin.Context) {
	items := []model.GuessItem{
		{
			Id:        "7a7b309c-a925-424b-9c91-6827ca73cf35",
			Src:       domain + "/image?imageName=001",
			Title:     "清真金园涮肉",
			Desc:      "代金券1张，除酒水外全场通用。代金券1张，除酒水外全场通用。代金券1张，除酒水外全场通用。",
			GoodPrice: "69",
			Price:     "100",
			SoldNum:   1071,
		}, {
			Id:        "36affccc-71e6-410e-99db-a624178ff04d",
			Src:       domain + "/image?imageName=002",
			Title:     "奶茶嫁给粉",
			Desc:      "[7店通用]低至8.8折代金券1张",
			GoodPrice: "8.8",
			Price:     "10",
			SoldNum:   24595,
		}, {
			Id:        "2fce074e-a70e-4456-8158-f10a618766e0",
			Src:       domain + "/image?imageName=003",
			Title:     "品匠老北京铜锅涮肉",
			Desc:      "招牌双人套餐。",
			GoodPrice: "138",
			Price:     "209",
			SoldNum:   125,
		},
	}
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
	file, ok := ioutil.ReadFile(dir + "/static/IndexNav/" + imageName + ".png")
	if ok != nil {
		fmt.Println(ok.Error())
	}
	c.Writer.WriteString(string(file))
}
