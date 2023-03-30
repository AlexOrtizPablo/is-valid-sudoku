package isValidSudoku

import "fmt"

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		rowMap := map[byte]bool{}

		for j := 0; j < 9; j++ {
			if board[i][j] != 46 {
				if _, ok := rowMap[board[i][j]]; !ok {
					rowMap[board[i][j]] = true
				} else {
					return false
				}
			}
		}
	}

	for i := 0; i < 9; i++ {
		columnMap := map[byte]bool{}

		for j := 0; j < 9; j++ {
			if board[j][i] != 46 {
				if _, ok := columnMap[board[j][i]]; !ok {
					columnMap[board[j][i]] = true
				} else {
					return false
				}
			}
		}
	}

	if !isAllBoxesValid(board) {
		return false
	}

	return true
}

func isAllBoxesValid(board [][]byte) bool {
	if !isValidBox(0, 0, board) {
		return false
	}

	if !isValidBox(0, 3, board) {
		return false
	}

	if !isValidBox(0, 6, board) {
		return false
	}

	if !isValidBox(3, 0, board) {
		return false
	}

	return true
}

func isValidBox(startX int, startY int, board [][]byte) bool {
	boxMap := map[byte]bool{}

	for i := startX; i < (startX + 3); i++ {
		for j := startY; j < (startY + 3); j++ {
			if board[i][j] != 46 {
				fmt.Println(board[i][j])

				if _, ok := boxMap[board[i][j]]; !ok {
					boxMap[board[i][j]] = true
				} else {
					return false
				}
			}
		}
	}

	return true
}
