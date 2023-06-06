package flexapi

import (
	"flex/movie"
	"fmt"
	"html"
	"log"
	"net/http"
)

// we will want to use a builder pattern to configure the server
type Server struct {
	//idk what to put in this rn but we will want it latter
	MovieHandler *movie.MovieHandler
}

func (s Server) BuildEndpoints() {
	http.HandleFunc("/getMovies", handleGetMovies)
	http.HandleFunc("/getInfo", handleGetMovieInfo)
	http.HandleFunc("/playFile", handlePlayFile)
	http.HandleFunc("/stopFile", handleStopFile)
	http.HandleFunc("/ws", s.handleWebSocket)
	// Test the websocket
	http.HandleFunc("/echo", s.handleWebSocket)
	http.HandleFunc("/", handleSendRoot)
}

func (Server) Serve() {
	// Start the server
	log.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleSendRoot(w http.ResponseWriter, r *http.Request) {
	// list out local movies
	fmt.Println("Called route!")
	http.ServeFile(w, r, "websockets.html")
}

func handleGetMovies(w http.ResponseWriter, r *http.Request) {
	// list out local movies
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	fmt.Println("Called get movies route!")
}

func handleGetMovieInfo(w http.ResponseWriter, r *http.Request) {
	// Get video file from request, and send back what we know about it.
	// Ex name of file with no extension, length, cover art, etc.
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	fmt.Println("Called get movie info route!")
}

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
