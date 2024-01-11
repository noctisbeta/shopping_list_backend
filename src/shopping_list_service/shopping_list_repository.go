package shopping_list_service

import (
	"errors"
	"log"

	nPS "github.com/noctisbeta/shopping_list/src/postgres_service"
)

type IShoppingListRepository interface {
	AddItems(items []ShoppingListItem) (*[]ShoppingListItem, error)
	AddItem(request AddItemRequest) (*ShoppingListItem, error)
	GetItems(code string) (*[]ShoppingListItem, error)
}

type shoppingListRepository struct {
	postgresService nPS.IPostgresService
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

func (r *shoppingListRepository) AddItems(items []ShoppingListItem) (*[]ShoppingListItem, error) {

	return nil, errors.New("not implemented")

	// query := "INSERT INTO items (id, name, price, quantity, shopping_list_id) VALUES ($1, $2, $3, $4, $5)"
	// result, err := r.postgresService.GetDB().Exec(query, code)

	// log.Println("HERE")

	// if err != nil {
	// 	log.Println(err)
	// 	return nil, err
	// }

	// rowsAffected, err := result.RowsAffected()
	// if err != nil {
	// 	log.Println(err)
	// 	return nil, err
	// }

	// if rowsAffected == 0 {
	// 	log.Println("No rows inserted")
	// 	return nil, errors.New("no rows inserted")
	// }

	// return &room, nil
}

func (r *shoppingListRepository) AddItem(request AddItemRequest) (*ShoppingListItem, error) {

	query := "INSERT INTO items (name, price, quantity, shopping_list_id) VALUES ($1, $2, $3, $4) RETURNING id"

	result, err := r.postgresService.GetDB().Exec(query, request.Name, request.Price, request.Quantity, request.Code)

	log.Println(result)

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

	item := ShoppingListItem{
		Name:             request.Name,
		Price:            request.Price,
		Quantity:         request.Quantity,
		ShoppingListCode: request.Code,
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
		err = rows.Scan(&item.Id, &item.Name, &item.Price, &item.Quantity, &item.ShoppingListCode)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		items = append(items, item)
	}

	return &items, nil
}
