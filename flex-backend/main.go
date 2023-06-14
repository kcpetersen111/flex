package main

import (
	"flag"
	"log"
	"os"

	server "github.com/kcpetersen111/flex/api/http"
	"github.com/kcpetersen111/flex/movieHandler"

	"github.com/joho/godotenv"
)

func main() {
	// Get line numbers in log messages
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var sAddr, dbAddr, dbUser, dbPassword string
	var port int
	flag.StringVar(&sAddr, "s", "0.0.0.0", "Server Address")
	flag.IntVar(&port, "p", 8080, "Server Port")
	flag.StringVar(&dbAddr, "dbAddr", "0.0.0.0:5432", "Database Address")
	flag.StringVar(&dbUser, "dbUser", "root", "Database User")
	flag.StringVar(&dbPassword, "dbPassword", "password", "Database Password")

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
	log.Printf("The path is: %s\n", path)

	MovieHandler := movieHandler.NewMovieHandler(path)

	MovieHandler.ListMovies()

	server := server.Server{MovieHandler}

	server.BuildEndpoints()
	server.Serve(port)

	return
}
