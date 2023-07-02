package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

var i int = 0

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./public")))
	mux.HandleFunc("/api/incriment", func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			http.SetCookie(w, &http.Cookie{Name: "token", Value: "1", Expires: time.Now().Add(365 * 24 * time.Hour)})
			fmt.Println("error getting cookie, setting it up")
		} else {
			fmt.Println(cookie.Value)
		}
		if r.Method == "POST" {
			i++
		}
		w.Write([]byte(strconv.Itoa(i)))
	})
	fmt.Println("server is running on 8080")
	http.ListenAndServe(":8080", mux)
}
