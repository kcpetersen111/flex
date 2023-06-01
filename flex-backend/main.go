package main

import (
	"flag"
	"flex/movie"
	"os"
)

func main() {

	var sAddr string
	var port int
	flag.StringVar(&sAddr, "s", "0.0.0.0", "Server Address")
	flag.IntVar(&port, "p", 8080, "Server Port")
	// Read env variable for the movie directory
	path := os.Getenv("dataDir")

	if len(path) == 0 {
		// Environment variable doesn't exist, set default
		path = "/Movie"
	}

	movie := movie.NewMovie(path)

	movie.ListMovies()

	return

}
