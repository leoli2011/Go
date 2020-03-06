package main

import (
	"net/http"
	"fmt"
	"log"
)


func myhandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "helloworld!!!")
}

func gohandler(w http.ResponseWriter, r *http.Request)  {
	fmt.Println("got the request")
	fmt.Fprintf(w, "gogogo!!!")
}

func main() {
	http.HandleFunc("/hello", myhandler)
	http.HandleFunc("/go", gohandler)
	err :=http.ListenAndServe("127.0.0.1:9008", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
