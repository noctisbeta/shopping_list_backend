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
	}
}

func (r *shoppingListRepository) AddItem(request AddItemRequest) (*ShoppingListItem, error) {

	id, err := r.GetRoomIdByCode(request.Code)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	query := "INSERT INTO items (name, price, quantity, room_id) VALUES ($1, $2, $3, $4) RETURNING id"

	var lastInsertId int

	err = r.postgresService.GetDB().QueryRow(query, request.Name, request.Price, request.Quantity, id).Scan(&lastInsertId)

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
		RoomCode: request.Code,
		RoomID:   id,
	}

	return &item, nil
}

func (r *shoppingListRepository) GetItems(code string) (*[]ShoppingListItem, error) {

	query := "SELECT * FROM items WHERE shopping_list_id = $1"
	rows, err := r.postgresService.GetDB().Query(query, code)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	items := []ShoppingListItem{}

	for rows.Next() {
		var item ShoppingListItem
		err = rows.Scan(&item.ID, &item.Name, &item.Price, &item.Quantity, &item.RoomCode)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		items = append(items, item)
	}

	return &items, nil
}

func (r *shoppingListRepository) GetRoomIdByCode(code string) (int, error) {
	// query := "SELECT id FROM rooms WHERE code = $1"

	// err := r.postgresService.GetDB().QueryRow(query, code).Scan(&id)

	room, err := r.roomRepository.GetRoomByCode(code)

	if err != nil {
		log.Println(err)
		return 0, err
	}

	return room.ID, nil
}
