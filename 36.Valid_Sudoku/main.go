package main

func isValidSudoku(board [][]byte) bool {
	n := 9
	rows := make([]map[byte]bool, n)
	cols := make([]map[byte]bool, n)
	boxes := make([]map[byte]bool, n)

	for i := 0; i < n; i++ {
		rows[i] = map[byte]bool{}
		cols[i] = map[byte]bool{}
		boxes[i] = map[byte]bool{}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			val := board[i][j]
			boxIndex := (i/3)*3 + j/3
			if val == '.' {
				continue
			}
			if rows[i][val] || cols[j][val] || boxes[boxIndex][val] {
				return false
			}

			rows[i][val] = true
			cols[i][val] = true
			boxes[boxIndex][val] = true
		}
	}

	return true
}

func main() {

}
