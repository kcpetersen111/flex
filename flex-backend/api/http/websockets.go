package flexapi

import (
	"context"
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
}

type connection struct {
	ctx    context.Context
	cancel context.CancelFunc
	conn   *websocket.Conn
}

type message struct {
	Test string `json:"test"`
}

// This is temporary to just test the websockets out. Need to change the structure
func (s Server) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	ws := &connection{
		ctx:    ctx,
		cancel: cancel,
		conn:   conn,
	}
	go ws.receive()
	go ws.send()
}

func (w *connection) receive() {
	defer w.cancel()
	msg := &message{}
	for {
		err := w.conn.ReadJSON(msg)
		if err != nil {
			//probably handle the connection closing differently than other errors that we may want to be more minor
			return
		}
		log.Printf("message received: %v\n", msg)
	}
}

func (w *connection) send() {
	defer w.cancel()

	m := message{
		Test: "success",
	}
	count := 0
	for {
		// eventually this should probably have a different go routine read the video file and send it over a channel to this
		select {
		case <-w.ctx.Done():
			return
		default:
			w.conn.WriteJSON(m)
			if count > 5 {
				return
			}
			count++
		}
	}
}

// I think the play and stop handlers need to be implemented within the websocket
func handlePlayFile(w http.ResponseWriter, r *http.Request) {
	// Start the ffmpeg stream from the specified file in the request.
	// Would we want to attach a session id to the process so we know who is playing what in the future?
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	fmt.Println("Called play movie route!")
}

func handleStopFile(w http.ResponseWriter, r *http.Request) {
	// Stop the ffmpeg stream from the specified file in the request.
	// Would we want to attach a session id to the process so we know how to stop the stream?
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	fmt.Println("Called stop movie route!")
}
