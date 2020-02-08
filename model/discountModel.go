package model

type DiscountModel struct {
	Id string `json:"id"`
	DiscountItemSrc string `json:"discountItemSrc"`
	DiscountItemHotel string `json:"discountItemHotel"`
	DiscountItemGoodsPrice int32 `json:"discountItemGoodsPrice"`
	DiscountItemTitle string `json:"discountItemTitle"`
	DiscountItemIcon string `json:"discountItemIcon"`
	DiscountItemPrice int32 `json:"discountItemPrice"`
	Discount string `json:"discount"`
}
