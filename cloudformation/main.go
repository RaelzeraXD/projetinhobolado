package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "teste cloudformation100")
	})
	http.ListenAndServe(":80", nil)
}
