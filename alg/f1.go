package alg

// Check if there is any blocking queen between two positions in the same row
func isRowVisible(board *[N][N]int, row, col1, col2 int) bool {
	if col1 > col2 {
		col1, col2 = col2, col1
	}
	for col := col1 + 1; col < col2; col++ {
		if board[row][col] == 1 {
			return false // there's a blocking queen
		}
	}
	return true
}

// Check if there is any blocking queen between two positions in the same column
func isColVisible(board *[N][N]int, col, row1, row2 int) bool {
	if row1 > row2 {
		row1, row2 = row2, row1
	}
	for row := row1 + 1; row < row2; row++ {
		if board[row][col] == 1 {
			return false // there's a blocking queen
		}
	}
	return true
}

// Check if there is any blocking queen between two positions in the same diagonal
func isDiagVisible(board *[N][N]int, row1, col1, row2, col2 int) bool {
	// Moving diagonally, check for blocking queens
	rowStep := 1
	colStep := 1
	if row1 > row2 {
		rowStep = -1
	}
	if col1 > col2 {
		colStep = -1
	}

	for r, c := row1+rowStep, col1+colStep; r != row2 && c != col2; r, c = r+rowStep, c+colStep {
		if board[r][c] == 1 {
			return false // there's a blocking queen
		}
	}
	return true
}

// Count pairs of queens that attack each other considering visibility (F1)
func F1(board *[N][N]int) int {
	count := 0

	// Iterate over all pairs of queens on the board
	for row1 := 0; row1 < N; row1++ {
		for col1 := 0; col1 < N; col1++ {
			if board[row1][col1] == 1 {
				// Check for another queen on the same row
				for col2 := col1 + 1; col2 < N; col2++ {
					if board[row1][col2] == 1 && isRowVisible(board, row1, col1, col2) {
						count++
					}
				}

				// Check for another queen on the same column
				for row2 := row1 + 1; row2 < N; row2++ {
					if board[row2][col1] == 1 && isColVisible(board, col1, row1, row2) {
						count++
					}
				}

				// Check for another queen on the diagonals
				for row2, col2 := row1+1, col1+1; row2 < N && col2 < N; row2, col2 = row2+1, col2+1 {
					if board[row2][col2] == 1 && isDiagVisible(board, row1, col1, row2, col2) {
						count++
					}
				}
				for row2, col2 := row1+1, col1-1; row2 < N && col2 >= 0; row2, col2 = row2+1, col2-1 {
					if board[row2][col2] == 1 && isDiagVisible(board, row1, col1, row2, col2) {
						count++
					}
				}
			}
		}
	}

	return count
}
