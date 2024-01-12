package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	nRS "github.com/noctisbeta/shopping_list/src/room"
	nSLS "github.com/noctisbeta/shopping_list/src/shopping_list"
)

func setupRoomServiceRoutes(router *gin.Engine) {
	roomServiceHandler := nRS.GetRoomServiceHandlerInstance()

	router.POST("/room", roomServiceHandler.HandleCreateRoom)
	router.GET("/room/:code", roomServiceHandler.HandleGetRoom)
}

func setupShoppingListServiceRoutes(router *gin.Engine) {
	shoppingListServiceHandler := nSLS.GetShoppingListServiceHandlerInstance()

	router.GET("/items/:code", shoppingListServiceHandler.HandleGetItems)
	router.POST("/items", shoppingListServiceHandler.HandleAddItem)
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	// add in cors headers for all
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	})

	// add in cors to allow post request
	r.OPTIONS("/room", func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Next()
	})

	r.OPTIONS("/items", func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Next()
	})

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	setupRoomServiceRoutes(r)
	setupShoppingListServiceRoutes(r)

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	// r.Run("localhost:8080")
	r.Run()
}
