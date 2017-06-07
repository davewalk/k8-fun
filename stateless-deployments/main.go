package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var (
	version = "0.0.5"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		panic("PORT has to be set")
	}

	router := mux.NewRouter()
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("got a request")
		w.Write([]byte(version))
	})
	http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}
