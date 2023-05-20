package main

import (
	"log"
	"net/http"
)

const PAGE = "<h1>B 228 SPM . . .</h1>"

func main() {
	log.Println("Started server on 0.0.0.0:8080, url: http://localhost:8080")

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte(PAGE))
	// })

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	http.ListenAndServe(":8080", nil)
}
