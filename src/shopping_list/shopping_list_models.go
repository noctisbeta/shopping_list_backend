package shopping_list_service

type ShoppingListItem struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
	Code     string  `json:"code"`
	RoomID   int     `json:"room_id"`
}

type AddItemRequest struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
	Code     string  `json:"code"`
}

type GetItemsRequest struct {
	Code string `json:"code"`
}
