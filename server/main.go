package main

import (
	// "net"
	"bufio"
	"fmt"
	"log"
	"os"
	// "time"
	"io/ioutil"
	"strings"
)

func main() {

	_, err := os.Stat("Movie")
	if os.IsNotExist(err) {
		os.Mkdir("Movie", 0750)
	}

	files, err := ioutil.ReadDir("Movie")
	if err != nil {
		log.Fatal(err)
	}
	movieList := make([]string, 0)
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".mkv") || strings.HasSuffix(file.Name(), ".mp4") { //Maybe we should create an array of known filetypes that we can use for video? Or should we just try and throw it in the transcoder and see if it gives an error?
			fmt.Printf("Found movie %s\n", file.Name())
			movieList = append(movieList, file.Name())
		}
	}

	fmt.Println("What would you like to watch?")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t := scanner.Text()

	moviePath := fmt.Sprintf("Movie/%s", t)
	movie, err := os.Stat(moviePath)
	if err != nil {
		log.Fatalf("There was an error: %v\n", err)
	}
	fmt.Printf("Movie name: %s\nMovie size: %d\nFile mode: %s\nLast modified: %s\n", movie.Name(),movie.Size(),movie.Mode(),movie.ModTime())

	return

}
