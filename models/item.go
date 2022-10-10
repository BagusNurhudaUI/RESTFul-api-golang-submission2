package models

type Item struct {
	Item_id     uint   `json:"item_id" gorm:"primaryKey"`
	Item_code   string `json:"item_code"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	Order_id    string `json:"-"`
}
