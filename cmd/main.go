package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello-world", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, "Hello world")
	})

	http.ListenAndServe(":8000", nil)

}
