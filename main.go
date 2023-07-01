package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var i int = 0

func main() {
	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/api/incriment", func(w http.ResponseWriter, r *http.Request) {
		if (r.Method == "POST") {
			i++
		}
		fmt.Fprintf(w, strconv.Itoa(i))
	})
	http.ListenAndServe(":8080", nil)
	fmt.Println("server is running on 8080")
}