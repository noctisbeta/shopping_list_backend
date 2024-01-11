package shopping_list_service

type IShoppingListService interface {
	AddItems(items []ShoppingListItem) (*[]ShoppingListItem, error)
	AddItem(request AddItemRequest) (*ShoppingListItem, error)
	GetItems(code string) (*[]ShoppingListItem, error)
}

type shoppingListService struct {
	shoppingListRepository IShoppingListRepository
}

var shoppingListServiceInstance *shoppingListService

func GetShoppingListServiceInstance() *shoppingListService {

	if shoppingListServiceInstance == nil {
		shoppingListServiceInstance = newShoppingListService()
	}

	return shoppingListServiceInstance
}

func newShoppingListService() *shoppingListService {
	return &shoppingListService{
		shoppingListRepository: GetShoppingListRepositoryInstance(),
	}
}

func (s *shoppingListService) AddItem(request AddItemRequest) (*ShoppingListItem, error) {
	item, err := s.shoppingListRepository.AddItem(request)

	if err != nil {
		return nil, err
	}

	return item, nil
}

func (s *shoppingListService) AddItems(items []ShoppingListItem) (*[]ShoppingListItem, error) {
	return nil, nil

	// items, err := s.shoppingListRepository.AddItems(items)

	// if err != nil {
	// 	return nil, err
	// }

	// return &items, nil
}

func (s *shoppingListService) GetItems(code string) (*[]ShoppingListItem, error) {
	items, err := s.shoppingListRepository.GetItems(code)

	if err != nil {
		return nil, err
	}

	return items, nil
}
