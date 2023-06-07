package flexapi

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// This is temporary to just test the websockets out. Need to change the structure
func (s Server) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Websocket will be created when the client wants to play a movie.
	// This is how we are going to stream the movie from the server to the client

	// Upgrade the plain http connection to a web socket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	// This will read a value from the socket, print the message to the console, and write it back to the client
	// Basically just a ping / pong to make sure that the connection works for now
	for {
		// Read message from browser
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}

		// Print the message to the console
		fmt.Printf("%s sent: %v\n", conn.RemoteAddr(), string(msg))

		// Write message back to browser
		if err = conn.WriteMessage(msgType, msg); err != nil {
			return
		}
	}
}
