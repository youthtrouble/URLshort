package main

import (
	"fmt"
	"net/http"

	"github.com/youthtrouble/URLshort/handlers"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/mytwitter":  "https://twitter.com/damndeji",
		"/mylinkedin": "https://www.linkedin.com/in/ayodeji-ajibola/?lipi=urn%3Ali%3Apage%3Ad_flagship3_feed%3B3JRtjxzbRVSv8Ama6RbOmw%3D%3D&licu=urn%3Ali%3Acontrol%3Ad_flagship3_feed-nav.settings_view_profile",
	}
	mapHandler := handlers.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml := `
- Path: /thisrepo
  URL: https://github.com/youthtrouble/URLshort.git
- Path: /mygithub
  URL: https://github.com/youthtrouble
`
	yamlHandler, err := handlers.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
