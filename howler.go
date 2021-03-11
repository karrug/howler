package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	var msg string
	http.HandleFunc("/m", func(w http.ResponseWriter, r *http.Request) {
		msg = r.FormValue("m")
		fmt.Fprintf(w, "ok")
	})
	http.HandleFunc("/g", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%s", msg)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
