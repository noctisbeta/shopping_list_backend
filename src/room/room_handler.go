package room_service

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IRoomServiceHandler interface {
	HandleCreateRoom(c *gin.Context)
	HandleGetRoom(c *gin.Context)
}

type roomServiceHandler struct {
	roomService IRoomService
}

var roomServiceHandlerInstance *roomServiceHandler

func GetRoomServiceHandlerInstance() *roomServiceHandler {

	if roomServiceHandlerInstance == nil {
		roomServiceHandlerInstance = newRoomServiceHandler()
	}

	return roomServiceHandlerInstance
}

func newRoomServiceHandler() *roomServiceHandler {
	return &roomServiceHandler{
		roomService: GetRoomServiceInstance(),
	}
}

func (h *roomServiceHandler) HandleCreateRoom(c *gin.Context) {
	var request RoomCreateRequest

	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	room, err := h.roomService.CreateRoom(request.Code)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"code": room.Code})
}

func (h *roomServiceHandler) HandleGetRoom(c *gin.Context) {
	var request RoomGetRequest

	// get code from url room/:code
	request.Code = c.Param("code")

	room, err := h.roomService.GetRoom(request.Code)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"code": room.Code})
}
