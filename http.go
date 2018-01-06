package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const (
	httpEndpoint = "0.0.0.0:3000"
)

func runHttp() {
	r := mux.NewRouter()

	r.HandleFunc("/api/time", httpGetTime).Methods("GET")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./www")))

	go func() {
		fmt.Fprintf(os.Stdout, "http: listening on \"%s\"\n", httpEndpoint)
		http.ListenAndServe(httpEndpoint, r)
	}()
}

func httpGetTime(w http.ResponseWriter, r *http.Request) {
	time := timeGet()

	bytes, err := json.Marshal(time)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(bytes)
	if err != nil {
		panic(err)
	}
}