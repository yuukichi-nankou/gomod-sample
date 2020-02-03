package main

import (
	"fmt"
	"net/http"
	"io"

	"rsc.io/quote"
)

func handle(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, quote.Hello())
}

func main() {
	portNumber := "8080"
	http.HandleFunc("/", handle)
	fmt.Println("Server listening on port ", portNumber)
	http.ListenAndServe(":"+portNumber, nil)
}
