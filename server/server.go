package main

import (
	"fmt"
	"log"
	"net/http"
)

func server() {

	// Send the index.html file to the client
	http.Handle("/", http.FileServer(http.Dir("../website")))
	http.HandleFunc("/getMovies", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Here are the detected files! , %q", readLocalDir())
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

	// Start the server
	log.Fatal(http.ListenAndServe(":8081", nil))
}
