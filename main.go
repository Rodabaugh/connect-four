package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type gameState struct {
	rows          int64
	cols          int64
	board         [][]int
	currentPlayer int
	gameOver      bool
	winner        int
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Using enviroment variables.")
	} else {
		log.Println("Loaded .env file.")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()

	gs := gameState{
		rows:          6,
		cols:          7,
		currentPlayer: 1,
		gameOver:      false,
	}

	gs.initBoard()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		MainPage(&gs.board).Render(r.Context(), w)
	})

	mux.HandleFunc("/get-refresh", gs.handleGetRefresh)

	mux.HandleFunc("/ws", handleWebSocket)

	mux.HandleFunc("POST /move/{col}", gs.makeMove)
	mux.HandleFunc("POST /reset", gs.reset)

	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Starting connect-four on port: %s\n", port)
	log.Fatal(server.ListenAndServe())
}
