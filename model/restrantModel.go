package model

type GoodRestaurantNavModel struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Desc string `json:"desc"`
	Src string `json:"src"`
}


type GoodRestaurantBillBoard struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Desc string `json:"desc"`
	Src string `json:"src"`
}

type GoodRestaurantTabItem struct {

	Id string `json:"id"`
	Src string `json:"src"`
	Title string `json:"title"`
	Star float32 `json:"star"`
	Price string `json:"price"`
	Area string `json:"area"`
	Type string `json:"type"`
	Desc string `json:"desc"`
	Team string `json:"team"`
	Quan string `json:"quan"`
}