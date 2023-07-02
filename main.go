package main

import (
	"fmt"
	"net/http"
)

var i int = 0

func main() {
	mux := InitHandlers()
	fmt.Println("server is running on 8080")
	http.ListenAndServe(":8080", mux)
}
