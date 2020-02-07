package main

import (
	"fmt"
	"net/http"

	"github.com/anandagireesh/urlshort/controller"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/mag":     "http://magiscloudstech.com/",
		"/mag-fet": "http://magiscloudstech.com/#feature",
	}
	mapHandler := controller.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml := `
- path: /mag
  url: "http://magiscloudstech.com/"
- path: /mag-fet
  url: "http://magiscloudstech.com/#feature"
`

	fmt.Println(yaml)
	yamlHandler, err := controller.YAMLHandler([]byte(yaml), mapHandler)
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
