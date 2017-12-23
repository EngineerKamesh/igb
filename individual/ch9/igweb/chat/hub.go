// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package chat

import (
	"github.com/EngineerKamesh/igb/igweb/bot"
)

type ClientMessage struct {
	client  *Client
	message []byte
}

// hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// A bot that implements the bot.Bot interface
	chatbot bot.Bot

	// Registered clients.
	clients map[*Client]bool

	// Inbound messages also containing the client who sent the message
	broadcastmsg chan *ClientMessage

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

func NewHub(chatbot bot.Bot) *Hub {
	return &Hub{
		chatbot:      chatbot,
		broadcastmsg: make(chan *ClientMessage),
		register:     make(chan *Client),
		unregister:   make(chan *Client),
		clients:      make(map[*Client]bool),
	}
}

func (h *Hub) SendMessage(client *Client, message []byte) {
	client.send <- message
}

func (h *Hub) ChatBot() bot.Bot {
	return h.chatbot
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
			greeting := h.chatbot.Greeting()
			h.SendMessage(client, []byte(greeting))

		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case clientmsg := <-h.broadcastmsg:
			client := clientmsg.client
			reply := h.chatbot.Reply(string(clientmsg.message))
			h.SendMessage(client, []byte(reply))
		}
	}
}
