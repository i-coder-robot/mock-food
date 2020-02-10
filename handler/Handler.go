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

func IndexHandler(c *gin.Context) {}

func DiscountLeftHandler(c *gin.Context) {
	items := []model.DiscountModel{
		{
			Id:                     "ae0181a5-a87e-4c5e-bdea-0a078fa13b6e",
			DiscountItemIcon:       domain + "/image?imageName=icon",
			DiscountItemSrc:        domain + "/image?imageName=goods003",
			DiscountItemTitle:      "仅79元！价值168元的午时女生半价自助，提供免提供免提供免提供免提供免提供免提供免提供免",
			DiscountItemHotel:      "熹炭炉端烧酒",
			DiscountItemGoodsPrice: 104,
			DiscountItemPrice:      208,
			Discount:               "5折",
		}, {
			Id:                     "bb9f6780-fbdf-4f71-8636-a8f10d3df1b6",
			DiscountItemIcon:       domain + "/image?imageName=icon",
			DiscountItemSrc:        domain + "/image?imageName=goods004",
			DiscountItemTitle:      "仅79元！价值168元的午时女生半价自助，提供免提供免提供免提供免提供免提供免提供免提供免",
			DiscountItemHotel:      "伊藤野菜村",
			DiscountItemGoodsPrice: 79,
			DiscountItemPrice:      168,
			Discount:               "5折",
		}, {
			Id:                     "bb1d56d8-f391-42cf-8266-e733e77d524c",
			DiscountItemSrc:        domain + "/image?imageName=goods005",
			DiscountItemTitle:      "仅79元！价值168元的午时女生半价自助，提供免提供免提供免提供免提供免提供免提供免提供免",
			DiscountItemIcon:       domain + "/image?imageName=icon",
			DiscountItemHotel:      "伊藤野菜村",
			DiscountItemGoodsPrice: 79,
			DiscountItemPrice:      168,
			Discount:               "5折",
		},
	}

	c.JSON(http.StatusOK, gin.H{
		"items": items,
	})
}

func DiscountRightHandler(c *gin.Context) {
	items := []model.DiscountModel{
		model.DiscountModel{
			Id:                     "0ed02e4e-f612-4bff-ab54-e8201b97379a",
			DiscountItemSrc:        domain + "/image?imageName=goods001",
			DiscountItemTitle:      "仅79元！价值168元的午时女生半价自助，提供免提供免提供免提供免提供免提供免提供免提供免",
			DiscountItemIcon:       domain + "/image?imageName=icon",
			DiscountItemHotel:      "伊藤野菜村",
			DiscountItemGoodsPrice: 79,
			DiscountItemPrice:      168,
			Discount:               "5折",
		},
		model.DiscountModel{
			Id:                     "ae0181a5-a87e-4c5e-bdea-0a078fa13b6e",
			DiscountItemSrc:        domain + "/image?imageName=goods002",
			DiscountItemTitle:      "仅79元！价值168元的午时女生半价自助，提供免提供免提供免提供免提供免提供免提供免提供免",
			DiscountItemIcon:       domain + "/image?imageName=icon",
			DiscountItemHotel:      "熹炭炉端烧酒",
			DiscountItemGoodsPrice: 79,
			DiscountItemPrice:      168,
			Discount:               "5折",
		},
	}
	c.JSON(http.StatusOK, gin.H{
		"items": items,
	})
}

func GoodRestaurantNavHandler(c *gin.Context) {
	items := []model.GoodRestaurantBillBoard{
		{
			Id:    "adafef98-570b-425e-8581-435496d05ec6",
			Title: "热门榜",
			Desc:  "全城火热",
			Src:   domain + "/image?imageName=good_restrant001",
		},
		{
			Id:    "0872d59b-3975-47a4-a939-ea609f770b19",
			Title: "评价榜",
			Desc:  "吃货都爱",
			Src:   domain + "/image?imageName=good_restrant_huoguo",
		},
		{
			Id:    "99ddd79c-d649-4ef7-adb2-43f689029d13",
			Title: "口味榜",
			Desc:  "舌尖盛宴",
			Src:   domain + "/image?imageName=tiandian",
		},
		{
			Id:    "504ec4b1-2680-4274-a792-9bb4433d8aad",
			Title: "小吃快餐榜",
			Desc:  "博大精深",
			Src:   domain + "/image?imageName=jiangzhe",
		},
		{
			Id:    "507ea105-6f50-41a5-9f93-7b7111e7c9d1",
			Title: "面包甜点榜",
			Desc:  "色味具佳",
			Src:   domain + "/image?imageName=xiaochi",
		},
		{
			Id:    "e596f743-4de7-42fe-97c8-875345477424",
			Title: "火锅榜",
			Desc:  "酣畅淋漓",
			Src:   domain + "/image?imageName=hanguo",
		},
	}
	c.JSON(http.StatusOK, gin.H{
		"items": items,
	})
}

func GoodRestaurantBillBoardHandler(c *gin.Context) {
	items := []model.GoodRestaurantNavModel{
		{
			Id:    "564541c5-f801-4f31-8a07-fe471996d582",
			Src:   domain + "/image?imageName=xiang",
			Title: "必吃榜",
			Desc:  "106家餐厅上榜",
		}, {
			Id:    "059b58eb-d6bf-4f7a-9ed2-8de6d899fee6",
			Src:   domain + "/image?imageName=ribencai",
			Title: "网红店榜",
			Desc:  "150家商户上榜",
		}, {
			Id:    "739d0505-cf56-4ee1-ac54-562890da1949",
			Src:   domain + "/image?imageName=dongbei",
			Title: "黑珍珠",
			Desc:  "20家餐厅上榜",
		},
	}
	c.JSON(http.StatusOK, gin.H{
		"items": items,
	})
}

func GoodRestaurantTabItemHandler(c *gin.Context) {
	items := []model.GoodRestaurantTabItem{
		{
			Id:    "7cd08b80-172b-4859-a5b5-4108ef616387",
			Src:   domain + "/image?imageName=002",
			Title: "小吊梨汤(德隆店)",
			Star:  5,
			Price: "￥89/人",
			Area:  "天通苑",
			Type:  "京菜",
			Desc:  "昌平区北京菜热门榜第2名",
			Team:  "小吊梨汤必吃套餐，建议2人使用形式形式形式",
			Quan:  "100元工作日代金券1张,100元工作日代金券1张",
		},
		{
			Id:    "2b13711f-c97e-4465-86b0-245aa46e3e15",
			Src:   domain + "/image?imageName=good_restrant002",
			Title: "开杠钵钵鸡 冒鸭血(右安门店)",
			Star:  4.5,
			Price: "￥54/人",
			Area:  "开阳里",
			Type:  "小吃快餐",
			Desc:  "开阳里小吃快餐热门榜第2名",
			Team:  "",
			Quan:  "",
		},
	}
	c.JSON(http.StatusOK, gin.H{
		"items": items,
	})
}

func MeHandler(c *gin.Context) {
	items1 := []model.Me{
		{
			Id:    "1ff2d71a-b523-434a-9d98-6f39f6eaa3d5",
			Src:   domain + "/image?imageName=me001",
			Title: "待付款",
		}, {
			Id:    "16172c9d-b859-488d-8557-7226d3686095",
			Src:   domain + "/image?imageName=me002",
			Title: "可使用",
		}, {
			Id:    "aa586b1d-8113-4392-abf0-2d7c01ec2132",
			Src:   domain + "/image?imageName=me003",
			Title: "退款/售后",
		}, {
			Id:    "9c281e08-9851-45a7-9384-10b0e1b26a99",
			Src:   domain + "/image?imageName=me004",
			Title: "全部订单",
		},
	}

	items2 := []model.Me{
		{
			Id:    "71a003c2-1dd5-4929-aa62-7358e69445e3",
			Src:   domain + "/image?imageName=me005",
			Title: "领券中心",
		}, {
			Id:    "cd5ea1ed-3543-423b-a473-14edf615f719",
			Src:   domain + "/image?imageName=me006",
			Title: "今天吃啥",
		}, {
			Id:    "a24c9321-a855-4eae-82e9-1473c24a46c6",
			Src:   domain + "/image?imageName=me007",
			Title: "聚餐投票",
		}, {
			Id:    "1b8425a7-1904-458c-83de-b48c622f3d12",
			Src:   domain + "/image?imageName=me008",
			Title: "贡献信息",
		},
	}
	items := model.Items{
		Items1: items1,
		Items2: items2,
		HeadPic:domain + "/image?imageName=headPic",
		Profile:domain + "/image?imageName=profile",
		MyCode:domain + "/image?imageName=myCode",
		Arrow:domain + "/image?imageName=arrow",
	}
	c.JSON(http.StatusOK, gin.H{
		"items": items,
	})
}

func NavHandler(c *gin.Context) {
	items := []model.IndexNavItem{
		model.IndexNavItem{
			Id:    "26cd5373-7709-468e-aa89-34ddcac20b3d",
			Src:   domain + "/image?imageName=meishIdatu",
			Title: "美食",
		}, model.IndexNavItem{
			Id:    "ab36ffb0-4c6c-4d6f-8ed5-20840896d496",
			Src:   domain + "/image?imageName=waimaIdatu",
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
	file, ok := ioutil.ReadFile(dir + "/static/images/" + imageName + ".png")
	if ok != nil {
		fmt.Println(ok.Error())
	}
	c.Writer.WriteString(string(file))
}
