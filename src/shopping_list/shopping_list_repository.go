package shopping_list_service

import (
	"errors"
	"log"

	nPS "github.com/noctisbeta/shopping_list/src/postgres"
	nRR "github.com/noctisbeta/shopping_list/src/room"
)

type IShoppingListRepository interface {
	AddItem(request AddItemRequest) (*ShoppingListItem, error)
	GetItems(code string) (*[]ShoppingListItem, error)
	GetRoomIdByCode(code string) (int, error)
}

type shoppingListRepository struct {
	postgresService nPS.IPostgresService
	roomRepository  nRR.IRoomRepository
}

var shoppingListRepositoryInstance *shoppingListRepository

func GetShoppingListRepositoryInstance() *shoppingListRepository {

	if shoppingListRepositoryInstance == nil {
		shoppingListRepositoryInstance = newShoppingListRepository()
	}

	return shoppingListRepositoryInstance
}

func newShoppingListRepository() *shoppingListRepository {
	return &shoppingListRepository{
		postgresService: nPS.GetPostgresServiceInstance(),
		roomRepository:  nRR.GetRoomRepositoryInstance(),
	}
}

func (r *shoppingListRepository) AddItem(request AddItemRequest) (*ShoppingListItem, error) {

	room, err := r.roomRepository.GetRoomByCode(request.Code)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	query := "INSERT INTO items (name, price, quantity, room_id) VALUES ($1, $2, $3, $4) RETURNING id"

	var lastInsertId int

	err = r.postgresService.GetDB().QueryRow(query, request.Name, request.Price, request.Quantity, room.ID).Scan(&lastInsertId)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	if lastInsertId == 0 {
		log.Println("No rows inserted")
		return nil, errors.New("no rows inserted")
	}

	item := ShoppingListItem{
		ID:       lastInsertId,
		Name:     request.Name,
		Price:    request.Price,
		Quantity: request.Quantity,
		Code:     request.Code,
		RoomID:   lastInsertId,
	}

	return &item, nil
}

func (r *shoppingListRepository) GetItems(code string) (*[]ShoppingListItem, error) {
	room, err := r.roomRepository.GetRoomByCode(code)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	query := "SELECT * FROM items WHERE room_id = $1"

	rows, err := r.postgresService.GetDB().Query(query, room.ID)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	items := []ShoppingListItem{}

	for rows.Next() {
		var item ShoppingListItem
		item.Code = code
		err = rows.Scan(&item.ID, &item.Name, &item.Price, &item.Quantity, &item.RoomID)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		items = append(items, item)
	}

	return &items, nil
}

func (r *shoppingListRepository) GetRoomIdByCode(code string) (int, error) {
	room, err := r.roomRepository.GetRoomByCode(code)

	if err != nil {
		log.Println(err)
		return 0, err
	}

	return room.ID, nil
}
