package model

type Me struct {
	Id string `json:"id"`
	Src string `json:"src"`
	Title string `json:"title"`
}

type Items struct {
	Items1 []Me `json:"items1"`
	Items2 []Me `json:"items2"`
	HeadPic string `json:"headPic"`
	Profile string 	`json:"profile"`
	MyCode string `json:"myCode"`
	Arrow string `json:"arrow"`
}