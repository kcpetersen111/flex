package flexapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	movie "github.com/kcpetersen111/flex/movieHandler"
)

// we will want to use a builder pattern to configure the server
type Server struct {
	//idk what to put in this rn but we will want it latter
	MovieHandler *movie.MovieHandler
}

func (s Server) BuildEndpoints() {
	http.HandleFunc("/getMovies", s.handleGetMovies)
	http.HandleFunc("/getInfo", s.handleGetMovieInfo)
	http.HandleFunc("/playFile", handlePlayFile)
	http.HandleFunc("/stopFile", handleStopFile)
	http.HandleFunc("/ws", s.handleWebSocket)
	// Used to test the websocket
	http.HandleFunc("/", handleSendRoot)
}

func (Server) Serve(port int) {
	// Start the server
	log.Printf("Starting server on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), nil))
}

func handleSendRoot(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "websockets.html")
}

func (s Server) handleGetMovies(w http.ResponseWriter, r *http.Request) {
	// Send back the filepaths to client as json
	files, err := s.MovieHandler.ReadLocalDir(s.MovieHandler.MovieDir)
	if err != nil {
		log.Fatal(err)
	}
	jData, err := json.Marshal(files)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}

func (s Server) handleGetMovieInfo(w http.ResponseWriter, r *http.Request) {
	// Get video file from request, and send back what we know about it.
	// Query param will need to have the full file path, since that is what we are sending to the client when it asks for the list of movies
	var movie string
	for k, v := range r.URL.Query() {
		// Get the first value of the query param that matches the string 'movie'
		// localhost:8080/getInfo?movie=1&movie=2
		// This would cause test to have 2 values, [1 2] instead of just 1
		if k == "movie" {
			movie = v[0]
			break
		}
	}

	movieInfo, err := s.MovieHandler.GetMovieInfo(string(movie))
	if err != nil {
		log.Println(err)
	}

	// Size is the number of bytes within the file (1000000 bytes = 1MB)
	type movieInfoResponse struct {
		Name string
		Size int64
		Mode string
	}

	movieResponse := movieInfoResponse{}
	movieResponse.Name = movieInfo.Name()
	movieResponse.Size = movieInfo.Size()
	// Convert the mode from rwx-rwx-rwx to octal (777 for example)
	movieResponse.Mode = fmt.Sprintf("%o", movieInfo.Mode())

	// Send the response back to the client
	jData, err := json.Marshal(movieResponse)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}
