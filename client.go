package main

import (
	"fmt"

	"github.com/gorilla/websocket"
)

// FindHandler func
type FindHandler func(string) (Handler, bool)

// Message Name and Data
type Message struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

// Client Name and Data
type Client struct {
	send        chan Message
	socket      *websocket.Conn
	findHandler FindHandler
}

func (client *Client) Read() {
	var message Message
	for {
		if err := client.socket.ReadJSON(&message); err != nil {
			break
		}
		// what function to call?
		if handler, found := client.findHandler(message.Name); found {
			handler(client, message.Data)
		}
	}
	client.socket.Close()
}

func (client *Client) Write() {
	for msg := range client.send {
		fmt.Printf("%#v\n", msg)
		if err := client.socket.WriteJSON(msg); err != nil {
			break
		}
	}
	client.socket.Close()
}

// NewClient Method
func NewClient(socket *websocket.Conn, findHandler FindHandler) *Client {
	return &Client{
		send:        make(chan Message),
		socket:      socket,
		findHandler: findHandler,
	}
}
