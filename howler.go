package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	var m string
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "favicon.ico")
	})
	http.HandleFunc("/g", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", m)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		m = r.URL.Path[1:]
	})
	log.Fatal(http.ListenAndServe(":8050", nil))
}
