package shopping_list_service

type ShoppingListItem struct {
	Id               int     `json:"id"`
	Name             string  `json:"name"`
	Price            float64 `json:"price"`
	Quantity         int     `json:"quantity"`
	ShoppingListCode string  `json:"shopping_list_id"`
}

type AddItemRequest struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
	Code     string  `json:"code"`
}

type AddItemsRequest struct {
	Items  []ShoppingListItem `json:"items"`
	RoomId string             `json:"room_id"`
}

type GetItemsRequest struct {
	Code string `json:"code"`
}
