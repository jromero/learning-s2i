package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const port = "8080"

func main() {
	log.Println("Starting app on port:", port)
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, "This is a test app!")
	})
	r.HandleFunc("/host", func(w http.ResponseWriter, r *http.Request) {
		var name, _ = os.Hostname()
		_, _ = fmt.Fprintf(w, "This request was processed by host: %s\n", name)
	})
	r.HandleFunc("/hello/{object}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		_, _ = fmt.Fprintf(w, "Hello %v!\n", vars["object"])
	})
	http.Handle("/", r)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
	log.Println("Shutting down")
}
