package main

import (
	"fmt"
	"net/http"
	"os"
	
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "This is a test app!")
	})
	r.HandleFunc("/host", func(w http.ResponseWriter, r *http.Request) {
		var name, _ = os.Hostname()
		fmt.Fprintf(w, "This request was processed by host: %s\n", name)
	})
	r.HandleFunc("/hello/{object}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Fprintf(w, "Hello %v!\n", vars["object"])
	})
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}