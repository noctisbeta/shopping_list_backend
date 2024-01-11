package shopping_list_service

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IShoppingListServiceHandler interface {
	HandleAddItems(c *gin.Context)
	HandleGetItems(c *gin.Context)
	HandleAddItem(c *gin.Context)
}

type shoppingListServiceHandler struct {
	shoppingListService IShoppingListService
}

var shoppingListServiceHandlerInstance *shoppingListServiceHandler

func GetShoppingListServiceHandlerInstance() *shoppingListServiceHandler {

	if shoppingListServiceHandlerInstance == nil {
		shoppingListServiceHandlerInstance = newShoppingListServiceHandler()
	}

	return shoppingListServiceHandlerInstance
}

func newShoppingListServiceHandler() *shoppingListServiceHandler {
	return &shoppingListServiceHandler{
		shoppingListService: GetShoppingListServiceInstance(),
	}
}

func (h *shoppingListServiceHandler) HandleAddItems(c *gin.Context) {
	var request AddItemsRequest

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	items, err := h.shoppingListService.AddItems(request.Items)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"code": items})
}

func (h *shoppingListServiceHandler) HandleAddItem(c *gin.Context) {
	var request AddItemRequest

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	item, err := h.shoppingListService.AddItem(request)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// c.JSON(200, gin.H{"code": items})
	c.JSON(200, gin.H{"item": item})
}

func (h *shoppingListServiceHandler) HandleGetItems(c *gin.Context) {
	log.Println("HandleGetItems")
	var request GetItemsRequest

	// get code from url items/:code
	request.Code = c.Param("code")

	items, err := h.shoppingListService.GetItems(request.Code)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"items": items})
}
