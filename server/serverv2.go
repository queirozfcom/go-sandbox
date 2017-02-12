package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handlerv2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome, %!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handlerv2)
	http.HandleFunc("/about/", about)
	http.ListenAndServe(":8080", nil)
}

type Message struct {
	Text string
}

func about(w http.ResponseWriter, r *http.Request) {
	m := Message{"foo-bar baz"}
	b, err := json.Marshal(m)

	if err != nil {
		panic(err)
	}

	w.Write(b)
}
