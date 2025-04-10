package ws

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Handler struct {
	hub *Hub
}

func NewHandler(hub *Hub) *Handler {
	return &Handler{
		hub: hub,
	}
}

type CreateRoomRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) CreateRoom(c *gin.Context) {
	var request CreateRoomRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.hub.Rooms[request.ID] = &Room{
		ID:      request.ID,
		Name:    request.Name,
		Clients: make(map[string]*Client),
	}

	c.JSON(http.StatusOK, request)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) JoinRoom(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//ws/JoinRoom/:roomID?userID=1&username=username
	roomID := c.Param("roomID")
	clientID := c.Query("userID")
	username := c.Query("username")

	cl := &Client{
		Conn:    conn,
		Message: make(chan *Message, 10),

		ID:       clientID,
		RoomID:   roomID,
		Username: username,
	}

	m := &Message{
		Content:  "A new user has joined the room",
		RoomID:   roomID,
		Username: username,
	}

	// register a new client through the register channel
	h.hub.Register <- cl

	// Broadcase the message
	h.hub.Broadcast <- m

	go cl.writeMessage()
	cl.readMessage(h.hub)
}

type RoomResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (h *Handler) GetRooms(c *gin.Context) {
	rooms := make([]RoomResponse, 0)

	for _, r := range h.hub.Rooms {
		rooms = append(rooms, RoomResponse{
			ID:   r.ID,
			Name: r.Name,
		})
	}

	c.JSON(http.StatusOK, rooms)
}

type ClientResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

func (h *Handler) GetClients(c *gin.Context) {
	var clients []ClientResponse
	roomID := c.Param("roomID")

	if _, exists := h.hub.Rooms[roomID]; !exists {
		clients = make([]ClientResponse, 0)
		c.JSON(http.StatusOK, clients)
	}

	// if already have client in room
	for _, c := range h.hub.Rooms[roomID].Clients {
		clients = append(clients, ClientResponse{
			ID:       c.ID,
			Username: c.Username,
		})
	}

	c.JSON(http.StatusOK, clients)
}
