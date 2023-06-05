package main

import (
	"flag"
	server "flex/api/http"
	"flex/movie"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	var sAddr string
	var port int
	flag.StringVar(&sAddr, "s", "0.0.0.0", "Server Address")
	flag.IntVar(&port, "p", 8080, "Server Port")

	// Read env variable for the movie directory from .env first, then look at the global environment variables
	// If the env file doesn't exist, then let's just log the message anyways
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}
	path := os.Getenv("dataDir")

	if len(path) == 0 {
		// Environment variable doesn't exist, set default
		path = "/Movies"
	}

	movie := movie.NewMovie(path)

	movie.ListMovies()

	server := server.Server{}

	server.BuildEndpoints()
	server.Serve()

	return

}
