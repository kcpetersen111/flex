package movie

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// A list of file types that we should support
// For now let's just start with the basic mp4 file format
var supportedFileExtensions = []string{"*.mp4"}

type Movie struct {
	movieDir string
}

func NewMovie(movieDir string) *Movie {
	return &Movie{
		movieDir: movieDir,
	}
}

func (m Movie) readLocalDir(root string) ([]string, error) {
	// return a slice of just the file names within a movieDir directory
	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		for _, pattern := range supportedFileExtensions {
			// Check each file to check the extension. Might not be the most efficient
			if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
				return err
			} else if matched {
				matches = append(matches, path)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}

func (m Movie) ListMovies() {
	_, err := os.Stat(m.movieDir)
	if os.IsNotExist(err) {
		err = os.Mkdir(m.movieDir, 0750)
		if err != nil {
			log.Fatal(err)
			return
		}
	}

	files, err := m.readLocalDir(m.movieDir)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("Found these movies!\n%v\n", files)
}

func (m Movie) getMovieInfo() {
	fmt.Println("What would you like to watch?")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t := scanner.Text()

	moviePath := fmt.Sprintf("%s/%s", m.movieDir, t)
	movie, err := os.Stat(moviePath)
	if err != nil {
		log.Fatalf("There was an error: %v\n", err)
	}
	fmt.Printf("Movie name: %s\nMovie size: %d\nFile mode: %s\nLast modified: %s\n", movie.Name(), movie.Size(), movie.Mode(), movie.ModTime())

}
