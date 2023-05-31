package main

import (
	"fmt"
	"log"
	"net/http"
)

func server() {

	http.Handle("/", http.FileServer(http.Dir("../website")))
	// http.Handle("/getMovies", myHelloWorldFunc())
	http.HandleFunc("/getMovies", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Here are the detected files! , %q", readLocalDir())
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
