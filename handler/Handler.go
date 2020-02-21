package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/i-coder-robot/mock-food/model"
	"io/ioutil"
	"net/http"
	"os"
)

const domain = "http://192.168.0.104:8080"

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
		Items1:  items1,
		Items2:  items2,
		HeadPic: domain + "/image?imageName=headPic",
		Profile: domain + "/image?imageName=profile",
		MyCode:  domain + "/image?imageName=myCode",
		Arrow:   domain + "/image?imageName=arrow",
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
		TeamImg:       "http://192.168.0.104:8080/image?imageName=niuroumian",
		FoodName:      "牛肉面",
		TeamIcon:      "http://192.168.0.104:8080/image?imageName=icon",
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
		TeamImg:       "http://192.168.0.104:8080/image?imageName=rush",
		FoodName:      "肥牛丼套餐1份",
		TeamIcon:      "http://192.168.0.104:8080/image?imageName=icon",
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

func SameCityHandler(c *gin.Context) {
	items1 := []model.SameCity{
		model.SameCity{
			Id:        "643075b6-13eb-4ba0-a6e9-a46d61017943",
			Title:     "1917电影",
			PhotoAddr: "http://192.168.0.104:8080/image?imageName=douyin001",
			Distance:  "8.6km",
			AvatarUrl: "http://192.168.0.104:8080/image?imageName=60",
			Location:  "北京市",
			Random:    9,
			Desc:      "1917年，第一次世界大战进入最激烈之际，两个年仅16岁的英国士兵接到的命令，需立即赶往死亡前线，向那里的将军传达一个“立刻停止进攻”讯息。 ",
		},
		model.SameCity{
			Id:        "be4b27db-85ea-447a-882c-e777aaee0f7e",
			Title:     "寄生虫",
			PhotoAddr: "http://192.168.0.104:8080/image?imageName=douyin002",
			Distance:  ">10km",
			AvatarUrl: "http://192.168.0.104:8080/image?imageName=61",
			Location:  "北京市三里屯通盈中心洲际酒店",
			Random:    5,
			Desc:      "就这样，基宇来到了朴社长（李善均 饰）家中，并且见到了他的太太（赵汝贞 饰），没过多久，基宇的妹妹和父母也如同寄生虫一般的进入了朴社长家里工作。",
		},
		model.SameCity{
			Id:        "fbd35e30-3734-4e88-9a85-bc5a063653b8",
			Title:     "",
			PhotoAddr: "http://192.168.0.104:8080/image?imageName=douyin003",
			Distance:  "8.6km",
			AvatarUrl: "http://192.168.0.104:8080/image?imageName=21",
			Location:  "北京市",
			Random:    8,
			Desc:      "基宇（崔宇植 饰）出生在一个贫穷的家庭之中，和妹妹基婷（朴素丹 饰）以及父母在狭窄的地下室里过着相依为命的日子。",
		},
		model.SameCity{
			Id:        "a7caf770-2d7f-4615-98a1-612e3cc48c2b",
			Title:     "",
			PhotoAddr: "http://192.168.0.104:8080/image?imageName=douyin004",
			Distance:  "5.6km",
			AvatarUrl: "http://192.168.0.104:8080/image?imageName=80",
			Location:  "北京市",
			Random:    20,
			Desc:      "一天，基宇的同学上门拜访，他告诉基宇，自己在一个有钱人家里给他们的女儿做家教，太太是一个头脑简单出手又阔绰的女人，因为自己要出国留学，所以将家教的职位暂时转交给基宇。",
		},
	}
	items2 := []model.SameCity{
		model.SameCity{
			Id:        "96e3624b-1b98-46a2-9278-cc544bfcc3cc",
			Title:     "",
			PhotoAddr: "http://192.168.0.104:8080/image?imageName=douyin005",
			Distance:  "7.4km",
			AvatarUrl: "http://192.168.0.104:8080/image?imageName=42",
			Location:  "北京市",
			Random:    16,
			Desc:      "然而，他们的野心并没有止步于此，基宇更是和大小姐坠入了爱河。随着时间的推移，朴社长家里隐藏的秘密渐渐浮出了水面。",
		},
		model.SameCity{
			Id:        "4cf14f39-9a06-46f5-b5db-0eb154f38366",
			Title:     "",
			PhotoAddr: "http://192.168.0.104:8080/image?imageName=douyin006",
			Distance:  "6.6km",
			AvatarUrl: "http://192.168.0.104:8080/image?imageName=51",
			Location:  "北京市",
			Random:    4,
			Desc:      "电影改编自世界名著《小妇人》，由奥斯卡提名最佳导演格雷塔.葛伟格执导，直击当代女性困惑。",
		},
		model.SameCity{
			Id:        "7ebfa315-1e53-4fb0-ba16-d3992fc72135",
			Title:     "",
			PhotoAddr: "http://192.168.0.104:8080/image?imageName=douyin007",
			Distance:  "6.0km",
			AvatarUrl: "http://192.168.0.104:8080/image?imageName=50",
			Location:  "北京市",
			Random:    5,
			Desc:      "作为2020年颁奖季热门，影片聚齐两代神仙阵容“世纪大同框”——金球奖最佳女主“伯德小姐”西尔莎·罗南与屡获提名的超人气演员“甜茶”再度携手诠释错过的真爱，“赫敏”艾玛·沃森与“黑寡妇师妹”弗洛伦斯·皮尤联袂呈现手足情深，三获奥斯卡的老戏骨梅尔·斯特里普与四次金球奖得主劳拉·邓恩倾力加盟。导演葛伟格用细腻又颇具新意的叙事手法，为这部经典文学名作带来充满活力的现代风格。",
		},
		model.SameCity{
			Id:        "a782ccaa-43c8-420d-b0a2-5c089d450a6c",
			Title:     "",
			PhotoAddr: "http://192.168.0.104:8080/image?imageName=douyin008",
			Distance:  "9.6km",
			AvatarUrl: "http://192.168.0.104:8080/image?imageName=60",
			Location:  "北京市",
			Random:    9,
			Desc:      "《爱尔兰人》为马丁·斯科塞斯执导的传奇巨制，罗伯特·德尼罗、阿尔·帕西诺和乔·佩西主演。通过二战老兵弗兰克·希兰的视角，讲述了战后美国有组织犯罪的故事。",
		},
	}
	items := model.SameCityItems{
		Items1: items1,
		Items2: items2,
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
