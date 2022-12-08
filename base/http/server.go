package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "http http")
	w.Write([]byte("hello"))
}

func main() {
	http.HandleFunc("/hello", hello)

	http.ListenAndServe("127.0.0.1:7000", nil)
}
