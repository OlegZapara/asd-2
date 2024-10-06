package alg

import (
	"math"
	"time"
)

func heuristic(board *[8][8]int) int {
	conflicts := 0

	for i := 0; i < 8; i++ {
		rowSum := 0
		for j := 0; j < 8; j++ {
			rowSum += board[i][j]
		}
		if rowSum > 1 {
			conflicts += rowSum - 1
		}
	}

	for j := 0; j < 8; j++ {
		colSum := 0
		for i := 0; i < 8; i++ {
			colSum += board[i][j]
		}
		if colSum > 1 {
			conflicts += colSum - 1
		}
	}

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if board[i][j] == 1 {
				for d := 1; i+d < 8 && j+d < 8; d++ {
					if board[i+d][j+d] == 1 {
						conflicts++
					}
				}
				for d := 1; i+d < 8 && j-d >= 0; d++ {
					if board[i+d][j-d] == 1 {
						conflicts++
					}
				}
			}
		}
	}

	return conflicts
}

func RBFS(board *[8][8]int, row int, fLimit int, startPos [2]int, delay time.Duration) (bool, int) {
	if row == LIMIT {
		return true, 0
	}

	bestScore := math.MaxInt32
	for col := 0; col < 8; col++ {
		if isValid(board, row, col) {
			board[row][col] = 1
			time.Sleep(delay)

			h := heuristic(board)
			if h < bestScore {
				bestScore = h
			}

			if bestScore <= fLimit {
				found, score := RBFS(board, row+1, fLimit, startPos, delay)
				if found {
					return true, score
				}
			}

			board[row][col] = 0
		}
	}

	return false, bestScore
}
