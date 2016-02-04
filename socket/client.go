package socket

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second
	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second
	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

// generic base client for websocket communication
type Client struct {
	ws   *websocket.Conn
	send chan []byte
}

// function creates a client for special group
func CreateClient(ws *websocket.Conn) *Client {
	return &Client{
		ws:   ws,
		send: make(chan []byte, 256),
	}
}

// function removes a client from the hub
func (c *Client) removeClient() {
	MusicHub.unregister <- c

	c.ws.Close()
}

// function reads info from client's socket
func (c *Client) ReadPump() {
	defer c.removeClient()

	c.ws.SetReadLimit(maxMessageSize)
	c.ws.SetReadDeadline(time.Now().Add(pongWait))
	c.ws.SetPongHandler(func(string) error {
		c.ws.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	// listen for socket
	for {
		_, _, err := c.ws.ReadMessage()
		if err != nil {
			log.Print(err.Error())
			break
		}
	}
}

// function writes to client's socket
func (c *Client) write(messageType int, message []byte) error {
	c.ws.SetWriteDeadline(time.Now().Add(writeWait))
	return c.ws.WriteMessage(messageType, message)
}

// function gets messages from client's channel
// and can write to client by calling write function
func (c *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.ws.Close()
	}()

	// write to socket
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}

			err := c.write(websocket.BinaryMessage, message)
			if err != nil {
				return
			}

		case <-ticker.C:
			err := c.write(websocket.PingMessage, []byte{})
			if err != nil {
				return
			}
		}
	}
}
