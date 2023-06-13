package movieHandler

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// A list of file types that we should support
// For now let's just start with the basic mp4 file format
var supportedFileExtensions = []string{"*.mp4"}

type MovieHandler struct {
	MovieDir string
}

func NewMovieHandler(MovieDir string) *MovieHandler {
	return &MovieHandler{
		MovieDir: MovieDir,
	}
}

func (m MovieHandler) ReadLocalDir(root string) ([]string, error) {
	// return a slice of just the file names within a MovieDir directory
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

func (m MovieHandler) ListMovies() {
	_, err := os.Stat(m.MovieDir)
	if os.IsNotExist(err) {
		err = os.Mkdir(m.MovieDir, 0750)
		if err != nil {
			log.Fatal(err)
			return
		}
	}

	files, err := m.ReadLocalDir(m.MovieDir)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("Found these movies!\n%v\n", files)
}

func (MovieHandler) GetMovieInfo(moviePath string) (os.FileInfo, error) {
	movie, err := os.Stat(moviePath)
	if err != nil {
		return nil, err
	}
	return movie, nil

}
