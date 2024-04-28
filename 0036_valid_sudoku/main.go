package main

func main() {

}

func isValidSudoku(board [][]byte) bool {
	xSet := make([]map[byte]bool, 9)
	ySet := make([]map[byte]bool, 9)
	squareSet := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		xSet[i] = make(map[byte]bool)
		ySet[i] = make(map[byte]bool)
		squareSet[i] = make(map[byte]bool)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			currentValue := board[i][j]
			row := i
			col := j
			boxIndex := (row/3)*3 + (col / 3)

			if currentValue == '.' {
				continue
			}

			if xSet[row][currentValue] || ySet[col][currentValue] || squareSet[boxIndex][currentValue] {
				return false
			}

			xSet[row][currentValue] = true
			ySet[row][currentValue] = true
			squareSet[row][currentValue] = true
		}
	}
	return true
}
