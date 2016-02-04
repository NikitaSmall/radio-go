package socket

import (
	"log"
	"os"
	"time"

	"github.com/tcolgate/mp3"
)

// simple hub that manages all the socket connections as a clients
type Hub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
	unregister chan *Client
}

// creates a hub
func newHub() Hub {
	return Hub{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

// main hub of this app. To send something to a client use this hub
var MusicHub = newHub()

// function registers new client to a hub
func (hub *Hub) Register(client *Client) {
	hub.register <- client
}

// function removes a client from a hub
func (hub *Hub) unregisterClient(c *Client) {
	if _, ok := hub.clients[c]; ok {
		delete(hub.clients, c)
		close(c.send)
	}
}

// function checks is hub empty
func (hub *Hub) isEmpty() bool {
	return len(hub.clients) == 0
}

// function returns number of clients in a hub
func (hub *Hub) count() int {
	return len(hub.clients)
}

// function makes preparations to broadcast message to all clients
// if there are clients in a hub
func (hub *Hub) SendMessage(message []byte) {
	if hub.isEmpty() {
		return
	}

	hub.broadcast <- message
}

// main hub process
func (hub *Hub) Run() {
	go hub.streamStep()

	for {
		select {
		case c := <-hub.register:
			hub.clients[c] = true
		case c := <-hub.unregister:
			hub.unregisterClient(c)
		case m := <-hub.broadcast:
			hub.broadcastMessage(m)
		}
	}
}

// function broadcasts prepared message to all the clients in the hub
func (hub *Hub) broadcastMessage(m []byte) {
	for c := range hub.clients {
		select {
		case c.send <- m:
		default:
			close(c.send)
			delete(hub.clients, c)
		}
	}
}

func (hub *Hub) streamStep() {
	r, err := os.Open("music/ФинродЗонг - Привал.mp3")
	if err != nil {
		log.Println(err)
		return
	}

	d := mp3.NewDecoder(r)
	var f mp3.Frame

	for {
		if err := d.Decode(&f); err != nil {
			log.Println(err)
			if err.Error() == "EOF" {
				return
			} else {
				continue
			}
		}
		b := make([]byte, f.Size())

		f.Reader().Read(b)
		hub.SendMessage(b)

		time.Sleep(f.Duration())
	}
}
