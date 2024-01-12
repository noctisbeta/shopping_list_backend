package room_service

import (
	"errors"
	"log"

	nPS "github.com/noctisbeta/shopping_list/src/postgres"
)

type IRoomRepository interface {
	CreateRoom(code string) (*GetRoomDB, error)
	GetRoomByCode(code string) (*GetRoomDB, error)
	GetRoomByID(code int) (*GetRoomDB, error)
}

type roomRepository struct {
	postgresService nPS.IPostgresService
}

var roomRepositoryInstance *roomRepository

func GetRoomRepositoryInstance() *roomRepository {
	if roomRepositoryInstance == nil {
		roomRepositoryInstance = newRoomRepository()
	}
	return roomRepositoryInstance
}

func newRoomRepository() *roomRepository {
	return &roomRepository{
		postgresService: nPS.GetPostgresServiceInstance(),
	}
}

func (rr *roomRepository) CreateRoom(code string) (*GetRoomDB, error) {
	log.Println("Creating room with code: " + code)

	room := GetRoomDB{Code: code}

	query := "INSERT INTO rooms (code) VALUES ($1)"
	result, err := rr.postgresService.GetDB().Exec(query, code)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if rowsAffected == 0 {
		log.Println("No rows inserted")
		return nil, errors.New("no rows inserted")
	}

	return &room, nil
}

func (rr *roomRepository) GetRoomByCode(code string) (*GetRoomDB, error) {
	room := GetRoomDB{}

	query := "SELECT * FROM rooms WHERE code = $1"
	err := rr.postgresService.GetDB().QueryRow(query, code).Scan(&room.ID, &room.Code)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &room, nil
}

func (rr *roomRepository) GetRoomByID(id int) (*GetRoomDB, error) {
	room := GetRoomDB{}

	query := "SELECT (id, code) FROM rooms WHERE id = $1"
	err := rr.postgresService.GetDB().QueryRow(query, id).Scan(&room.ID, &room.Code)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &room, nil
}
