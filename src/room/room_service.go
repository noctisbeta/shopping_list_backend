package room_service

import "log"

type IRoomService interface {
	CreateRoom(code string) (GetRoomDB, error)
	GetRoom(code string) (GetRoomDB, error)
}

type roomService struct {
	roomRepository IRoomRepository
}

var roomServiceInstance *roomService

func GetRoomServiceInstance() *roomService {
	if roomServiceInstance == nil {
		roomServiceInstance = newRoomService()
	}
	return roomServiceInstance
}

func newRoomService() *roomService {
	return &roomService{
		roomRepository: GetRoomRepositoryInstance(),
	}
}

func (rs *roomService) CreateRoom(code string) (GetRoomDB, error) {
	room, err := rs.roomRepository.CreateRoom(code)
	if err != nil {
		log.Println(err)
		return GetRoomDB{}, err
	}
	return *room, nil
}

func (rs *roomService) GetRoom(code string) (GetRoomDB, error) {
	room, err := rs.roomRepository.GetRoomByCode(code)
	if err != nil {
		log.Println(err)
		return GetRoomDB{}, err
	}
	return *room, nil
}
