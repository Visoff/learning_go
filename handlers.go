package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func InitHandlers() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("./public")))
	mux.HandleFunc("/api/incriment", func(w http.ResponseWriter, r *http.Request) {
		/*cookie, err := r.Cookie("token")
		if err != nil {
			http.SetCookie(w, &http.Cookie{Name: "token", Value: "1", Expires: time.Now().Add(365 * 24 * time.Hour)})
			fmt.Println("error getting cookie, setting it up")
		}*/
		if r.Method == "POST" {
			i++
		} else if r.Method == "PATCH" {
			body_bynary, _ := io.ReadAll(r.Body)
			body := string(body_bynary)
			fmt.Println(body)
			val, err := strconv.Atoi(body)
			if err == nil {
				i = val
			}
		}
		w.Write([]byte(strconv.Itoa(i)))
	})

	return mux
}
