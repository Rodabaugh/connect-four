package main

import (
	"log"
	"net/http"
	"time"
)

type gameState struct {
	rows          int64
	cols          int64
	board         [][]int
	currentPlayer int
}

func main() {
	mux := http.NewServeMux()

	gs := gameState{
		rows:          6,
		cols:          7,
		currentPlayer: 1,
	}

	gs.initBoard()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		MainPage(&gs.board).Render(r.Context(), w)
	})

	mux.HandleFunc("POST /move/{col}", gs.makeMove)
	mux.HandleFunc("POST /reset", gs.reset)

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
