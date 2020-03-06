package main

import (
	"net/http"
	"fmt"
)

func HandConn(w http.ResponseWriter, r *http.Request) {
	fmt.Println("r.Method", r.Method)
	fmt.Println("r.HEAD", r.Header)
	fmt.Println("r.URL", r.URL)
	fmt.Println("r.Body", r.Body)

	w.Write([]byte("<h1>HELLO WORLD<h1>"))
}

func main() {
	http.HandleFunc("/", HandConn)
	http.ListenAndServe("127.0.0.1:8000", nil)
}
