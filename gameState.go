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
	gs.gameOver = false
	DrawBoard(&gs.board, false).Render(r.Context(), w)
	BroadcastBoardRefresh()
}

func (gs *gameState) makeMove(w http.ResponseWriter, r *http.Request) {
	col, err := strconv.ParseInt(r.PathValue("col"), 10, 0)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Col", err)
		return
	}

	if col > gs.cols || col < 0 {
		respondWithError(w, http.StatusBadRequest, "Invalid Col", nil)
		return
	}

	valid := false
	for row := len(gs.board) - 1; row >= 0; row-- {
		if gs.board[row][col] != 1 && gs.board[row][col] != 2 {
			gs.board[row][col] = gs.currentPlayer
			valid = true
			break
		}
	}

	if !valid {
		respondWithError(w, http.StatusBadRequest, "Invalid Col", nil)
		return
	}

	if gs.checkWin() {
		gs.gameOver = true
		gs.winner = gs.currentPlayer
		GameOver(&gs.board, gs.currentPlayer).Render(r.Context(), w)
		BroadcastBoardRefresh()
		return
	}

	if gs.boardFull() {
		gs.gameOver = true
		gs.winner = 0
		GameOver(&gs.board, 0).Render(r.Context(), w)
		BroadcastBoardRefresh()
		return
	}

	if gs.currentPlayer == 1 {
		gs.currentPlayer = 2
	} else {
		gs.currentPlayer = 1
	}

	DrawBoard(&gs.board, false).Render(r.Context(), w)
	BroadcastBoardRefresh()
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

func (gs *gameState) handleGetRefresh(w http.ResponseWriter, r *http.Request) {
	if gs.gameOver {
		GameOver(&gs.board, gs.winner).Render(r.Context(), w)
	} else {
		DrawBoard(&gs.board, false).Render(r.Context(), w)
	}
}

func (gs *gameState) boardFull() bool {
	for _, row := range gs.board {
		for _, cell := range row {
			if cell == 0 {
				return false
			}
		}
	}
	return true
}
