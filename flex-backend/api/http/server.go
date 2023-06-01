package http

import (
	"log"
	"net/http"
)

// we will want to use a builder pattern to configure the server
type Server struct {
	//idk what to put in this rn but we will want it latter
}

// im not a huge fan of this but it will do for now
// I think it should just have function names where there are lambda funcs rn
func (Server) BuildEndpoints() {

	http.HandleFunc("/getMovies", func(w http.ResponseWriter, r *http.Request) {
		// list out local movies
		// fmt.Fprintf(w, "Here are the detected files! , %q", readLocalDir())
	})

	http.HandleFunc("/getInfo", func(w http.ResponseWriter, r *http.Request) {
		// Get video file from request, and send back what we know about it.
		// Ex name of file with no extension, length, cover art, etc.
	})

	http.HandleFunc("/playFile", func(w http.ResponseWriter, r *http.Request) {
		// Start the ffmpeg stream from the specified file in the request.
		// Would we want to attach a session id to the process so we know who is playing what in the future?
	})

	http.HandleFunc("/stopFile", func(w http.ResponseWriter, r *http.Request) {
		// Stop the ffmpeg stream from the specified file in the request.
		// Would we want to attach a session id to the process so we know how to stop the stream?
	})
}

func (Server) Serve() {
	// Start the server
	log.Fatal(http.ListenAndServe(":8081", nil))
}
