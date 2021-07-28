package main

import (
	"data/numerals"
	"fmt"
	"html"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		urlelements := strings.Split(r.URL.Path, "/")
		if urlelements[1] == "roman_number" {
			number, _ := strconv.Atoi(strings.TrimSpace(urlelements[2]))
			if number == 0 || number > 10 {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte("404-Not Found"))
			} else {
				fmt.Fprintf(w, "%q", html.EscapeString(numerals.Roman(number)))
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400-Bad Request"))
		}
	})
	s := &http.Serve{
		Addr:           ":8000",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
