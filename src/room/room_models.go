package room_service

type RoomCreateRequest struct {
	Code string `json:"code"`
}

type RoomGetRequest struct {
	Code string `json:"code"`
}

type GetRoomDB struct {
	ID   int
	Code string
}
