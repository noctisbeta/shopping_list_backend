package room_service

import (
	"errors"
	"log"

	nPS "github.com/noctisbeta/shopping_list/src/postgres_service"
)

type IRoomRepository interface {
	CreateRoom(code string) (*Room, error)
	GetRoom(code string) (*Room, error)
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

func (rr *roomRepository) CreateRoom(code string) (*Room, error) {
	log.Println("Creating room with code: " + code)

	room := Room{Code: code}

	query := "INSERT INTO rooms (access_code) VALUES ($1)"
	result, err := rr.postgresService.GetDB().Exec(query, code)

	log.Println("HERE")

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

func (rr *roomRepository) GetRoom(code string) (*Room, error) {
	room := Room{}

	query := "SELECT access_code FROM rooms WHERE access_code = $1"
	err := rr.postgresService.GetDB().QueryRow(query, code).Scan(&room.Code)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &room, nil
}
