package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		log.Println("handling request from", r.RemoteAddr)
		time.Sleep(5 * time.Second)
		w.Write([]byte("hello"))
	})

	if err := http.ListenAndServe(":8088", nil); err != nil {
		log.Fatal(err)
	}
}
