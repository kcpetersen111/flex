package flexapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

const (
	LISTMOVIES int = iota
	GETMOVIEINFO
)

// Passed from the client to the server to tell it what action it wants performed
type Message struct {
	Message int    `json:"message"`
	Movie   string `json:"movie"`
}

func (s Server) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Need to perform an http upgrade and use some other library to handle web sockets here
	fmt.Println("Called the web socket route!")
	conn, _ := upgrader.Upgrade(w, r, nil) // error ignored for sake of simplicity

	for {
		// Read message from browser
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}
		fmt.Println("Message: ", string(msg))
		var test Message
		err = json.Unmarshal(msg, &test)

		// Print the message to the console
		fmt.Printf("%s sent: %v\n", conn.RemoteAddr(), test)

		switch test.Message {
		case LISTMOVIES:
			fmt.Println("Wanted to list movies!")
			var bytes []byte

			files, err := s.MovieHandler.ReadLocalDir(s.MovieHandler.MovieDir)
			if err != nil {
				log.Fatal(err)
			}

			for _, val := range files {
				for _, char := range val {
					bytes = append(bytes, byte(char))
				}
				bytes = append(bytes, byte(','))

				// fmt.Println(val)
			}

			if err = conn.WriteMessage(msgType, bytes); err != nil {
				return
			}
		case GETMOVIEINFO:
			fmt.Println("Wanted to get a specific movie information")
		}
		// Write message back to browser
		if err = conn.WriteMessage(msgType, msg); err != nil {
			return
		}
	}
}
