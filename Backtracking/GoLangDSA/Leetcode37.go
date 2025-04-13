package main

func isValid(board [][]byte, row int, coloumn int, c byte) bool {

	for i := 0; i < 9; i++ {

		if board[row][i] == c {

			return false
		}

		if board[i][coloumn] == c {

			return false
		}

		if board[3*(row/3)+i/3][3*(coloumn/3)+i%3] == c {

			return false
		}
	}

	return true
}
func startSolving(board [][]byte) bool {

	for i := 0; i < 9; i++ {

		for j := 0; j < 9; j++ {

			if board[i][j] == '.' {

				for c := byte('1'); c <= byte('9'); c++ {

					if isValid(board, i, j, c) {

						board[i][j] = c
					}

					if startSolving(board) == true {

						return true
					} else {

						board[i][j] = '.'
					}

				}

				return false
			}

		}
	}

	return false
}
func solveSudoku(board [][]byte) {

	startSolving(board)
}
