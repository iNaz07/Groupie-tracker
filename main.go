package main

import (
	"net/http"

	"./domain/artists"
	"./domain/locations"
	"./domain/relations"
)

func main() {
	fs := http.FileServer(http.Dir("static/public"))
	http.Handle("/", fs)

	http.HandleFunc("/info", handler)

	http.HandleFunc("/api/artists", artists.Handler)
	http.HandleFunc("/api/relations", relations.Handler)
	http.HandleFunc("/api/locations", locations.Handler)

	http.ListenAndServe(":9090", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/public/index.html")
}
