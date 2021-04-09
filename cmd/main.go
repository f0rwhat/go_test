package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/hello-world", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, "Hello world")
	})

	port, exists := os.LookupEnv("PORT")
	if exists == false {
		fmt.Printf("NO PORT")
	}
	http.ListenAndServe(":"+port, nil)

}
