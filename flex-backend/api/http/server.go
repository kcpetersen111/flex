package flexapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	movie "github.com/kcpetersen111/flex/movieHandler"
)

// we will want to use a builder pattern to configure the server
type Server struct {
	//idk what to put in this rn but we will want it latter
	MovieHandler *movie.MovieHandler
}

type movieBody struct {
	MoviePath string `json:moviePath`
}

func (s Server) BuildEndpoints() {
	http.HandleFunc("/getMovies", s.handleGetMovies)
	http.HandleFunc("/getInfo", s.handleGetMovieInfo)
	http.HandleFunc("/playFile", handlePlayFile)
	http.HandleFunc("/stopFile", handleStopFile)
	http.HandleFunc("/ws", s.handleWebSocket)
}

func (Server) Serve(port int) {
	// Start the server
	log.Printf("Starting server on port %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), nil))
}

func (Server) sendServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 - Something bad happened!"))
	w.Header().Set("Content-Type", "application/json")
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
		s.sendServerError(w)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}

func (s Server) handleGetMovieInfo(w http.ResponseWriter, r *http.Request) {
	// Get video file from request, and send back what we know about it.
	// Request body needs to be in the form of {"moviePath": "dir"}

	requestBody := movieBody{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		s.sendServerError(w)
	}

	if err := json.Unmarshal(body, &requestBody); err != nil {
		log.Println(err)
		s.sendServerError(w)
		return
	}

	// Decode the string just in case :)
	requestBody.MoviePath, err = url.QueryUnescape(requestBody.MoviePath)
	if err != nil {
		log.Println(err)
		s.sendServerError(w)
	}

	movieInfo, err := s.MovieHandler.GetMovieInfo(requestBody.MoviePath)
	if err != nil {
		log.Println(err)
	}

	// // Size is the number of bytes within the file (1000000 bytes = 1MB)
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
		s.sendServerError(w)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}
