package model

type GuessItem struct {
	Id        string `json:"id"`
	Src       string `json:"src"`
	Title     string `json:"title"`
	Desc      string `json:"desc"`
	GoodPrice string `json:"goodPrice"`
	Price     string `json:"price"`
	SoldNum   string `json:"soldNum"`
}
