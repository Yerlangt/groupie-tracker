package main

import (
	"fmt"
	"log"
	"net/http"

	handle "groupie-tracker/handle"
)

// var groupie = handle.Groupie{}

func main() {
	groupie := handle.GroupieCreator()
	// fmt.Println(groupie.Artists[1])
	fmt.Println("Data analyzed, listing o host: http://localhost:8080")
	mux := http.NewServeMux()
	mux.HandleFunc("/", groupie.MainHandler)
	mux.HandleFunc("/artist/", groupie.ArtistHandler)
	mux.HandleFunc("/search", groupie.SearchHandler)
	mux.Handle("/image/", http.StripPrefix("/image", http.FileServer(http.Dir("./web/image"))))
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./web/static"))))
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
