package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var MovieDir = "Movie"

func readLocalDir() []string {
	// return a slice of just the file names within a directory ( curently './Movie' )
	// Might want to implement this with a parameter which is the folder name to scan
	// and then recursively scan all directories in the movie root folder
	dirFiles, err := os.ReadDir(MovieDir)
	if err != nil {
		log.Fatal(err)
	}

	files := make([]string, 0)
	for _, val := range dirFiles {
		files = append(files, val.Name())
	}
	return files
}

func isValidFile(file string) bool {
	return (strings.HasSuffix(file, ".mkv") || strings.HasSuffix(file, ".mp4") || strings.HasSuffix(file, ".m4v"))
}

func main() {

	go func() {
		// Run the Http server
		server()
	}()

	_, err := os.Stat(MovieDir)
	if os.IsNotExist(err) {
		os.Mkdir(MovieDir, 0750)
	}

	files := readLocalDir()

	movieList := make([]string, 0)
	for _, file := range files {
		if isValidFile(file) { // Maybe we should create an array of known filetypes that we can use for video? Or should we just try and throw it in the transcoder and see if it gives an error?
			fmt.Printf("Found movie %s\n", file)
			movieList = append(movieList, file)
		}
	}

	fmt.Println("What would you like to watch?")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t := scanner.Text()

	moviePath := fmt.Sprintf("%s/%s", MovieDir, t)
	movie, err := os.Stat(moviePath)
	if err != nil {
		log.Fatalf("There was an error: %v\n", err)
	}
	fmt.Printf("Movie name: %s\nMovie size: %d\nFile mode: %s\nLast modified: %s\n", movie.Name(), movie.Size(), movie.Mode(), movie.ModTime())

	return

}
