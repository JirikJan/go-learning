package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Ahoj, tohle je moje první mikroslužba!")
	})

	http.ListenAndServe(":8080", nil)
}
