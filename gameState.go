package main

import (
	"net/http"
	"strconv"
)

func (gs *gameState) initBoard() {
	gs.board = make([][]int, gs.rows)
	for i := range gs.board {
		gs.board[i] = make([]int, gs.cols)
	}
}

func (gs *gameState) makeMove(w http.ResponseWriter, r *http.Request) {
	row, err := strconv.ParseInt(r.PathValue("row"), 10, 0)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Row", err)
	}
	col, err := strconv.ParseInt(r.PathValue("col"), 10, 0)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Col", err)
	}

	if row > gs.rows || row < 0 {
		respondWithError(w, http.StatusBadRequest, "Invalid Row", nil)
	}

	if row > gs.cols || col < 0 {
		respondWithError(w, http.StatusBadRequest, "Invalid Col", nil)
	}

	if gs.board[row][col] == 1 || gs.board[row][col] == 2 {
		respondWithError(w, http.StatusBadRequest, "Invalid Move", nil)
	}

	gs.board[row][col] = gs.currentPlayer

	if gs.currentPlayer == 1 {
		gs.currentPlayer = 2
	} else {
		gs.currentPlayer = 1
	}

	DrawBoard(&gs.board).Render(r.Context(), w)
}
