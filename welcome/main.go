package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	name := os.Getenv("NAME")
	helloMsg := fmt.Sprintf("hello %s", name)
	http.HandleFunc("/hello", hello(helloMsg))
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}

func hello(msg string) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w,
			msg)
	}
}
