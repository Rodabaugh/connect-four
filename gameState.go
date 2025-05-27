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

func (gs *gameState) reset(w http.ResponseWriter, r *http.Request) {
	gs.initBoard()
	gs.currentPlayer = 1
	DrawBoard(&gs.board, false).Render(r.Context(), w)
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

	if gs.checkWin() {
		GameOver(&gs.board, gs.currentPlayer).Render(r.Context(), w)
		return
	}

	if gs.currentPlayer == 1 {
		gs.currentPlayer = 2
	} else {
		gs.currentPlayer = 1
	}

	DrawBoard(&gs.board, false).Render(r.Context(), w)
}

func (gs *gameState) checkWin() bool {
	// Check horizontal
	for row := 0; row < len(gs.board); row++ {
		for cell := 0; cell <= len(gs.board[0])-4; cell++ {
			if gs.board[row][cell] == gs.currentPlayer && gs.board[row][cell+1] == gs.currentPlayer && gs.board[row][cell+2] == gs.currentPlayer && gs.board[row][cell+3] == gs.currentPlayer {
				return true
			}
		}
	}

	// Check vertical
	for row := 0; row <= len(gs.board)-4; row++ {
		for cell := 0; cell < len(gs.board[0]); cell++ {
			if gs.board[row][cell] == gs.currentPlayer && gs.board[row+1][cell] == gs.currentPlayer && gs.board[row+2][cell] == gs.currentPlayer && gs.board[row+3][cell] == gs.currentPlayer {
				return true
			}
		}
	}
	// Check diagonals (top-left to bottom-right)
	for row := 0; row <= len(gs.board)-4; row++ {
		for cell := 0; cell <= len(gs.board[0])-4; cell++ {
			if gs.board[row][cell] == gs.currentPlayer && gs.board[row+1][cell+1] == gs.currentPlayer && gs.board[row+2][cell+2] == gs.currentPlayer && gs.board[row+3][cell+3] == gs.currentPlayer {
				return true
			}
		}
	}
	// Check diagonals (top-right to bottom-left)
	for row := 0; row <= len(gs.board)-4; row++ {
		for cell := 3; cell < len(gs.board[0]); cell++ {
			if gs.board[row][cell] == gs.currentPlayer && gs.board[row+1][cell-1] == gs.currentPlayer && gs.board[row+2][cell-2] == gs.currentPlayer && gs.board[row+3][cell-3] == gs.currentPlayer {
				return true
			}
		}
	}
	return false
}
