package ws

type Room struct {
	ID      string             `json:"id"`
	Name    string             `json:"name"`
	Clients map[string]*Client `json:"clients"`
}

type Hub struct {
	Rooms map[string]*Room

	Register   chan *Client
	Unregister chan *Client
	Broadcast  chan *Message
}

func NewHub() *Hub {
	return &Hub{
		Rooms: make(map[string]*Room),

		Register:   make(chan *Client), // recives client info
		Unregister: make(chan *Client),
		Broadcast:  make(chan *Message, 5),
	}
}

func (h *Hub) Run() {
	for {
		select { // recives client info through this channels:
		case cl := <-h.Register:
			if room, exists := h.Rooms[cl.RoomID]; exists { // if room exists
				// add client to the room if not already present
				if _, ok := room.Clients[cl.ID]; !ok {
					room.Clients[cl.ID] = cl
				}
			}

		case cl := <-h.Unregister:
			if room, ok := h.Rooms[cl.RoomID]; ok {
				if _, ok := room.Clients[cl.ID]; ok {
					if len(room.Clients) > 0 {
						h.Broadcast <- &Message{
							Content:  "user left the chat",
							RoomID:   cl.RoomID,
							Username: cl.Username,
						}
					}

					delete(room.Clients, cl.ID)
					close(cl.Message)
					// if cl.Message != nil {
					// 	close(cl.Message)
					// }
				}
			}

		case m := <-h.Broadcast:
			if room, ok := h.Rooms[m.RoomID]; ok {
				for _, cl := range room.Clients { // sent to each client in that room
					// if cl.Message != nil {
					// 	cl.Message <- m
					// }
					cl.Message <- m
				}
			}
		}
	}
}
