package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()

	rows := 6
	cols := 7

	board := make([][]int, rows)
	for i := range board {
		board[i] = make([]int, cols)
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		MainPage(&board).Render(r.Context(), w)
	})

	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Starting Server on port 8080")
	log.Fatal(server.ListenAndServe())
}
