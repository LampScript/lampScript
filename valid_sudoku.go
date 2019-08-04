package main

import "fmt"

func main() {


	a := []byte{1,2,3,}
	fmt.Println(a)

}

/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Sudoku.
Memory Usage: 2.8 MB, less than 77.78% of Go online submissions for Valid Sudoku.

暴力查就完事了

*/

func isValidSudoku(board [][]byte) bool {
	for i:=0;i<len(board);i++  {
		if !checkRow(board, i) {
			return false
		}

		if !checkColumn(board, i) {
			return false
		}
		if i%3 == 0 {
			if !checkCell(board, i, 0) {
				return false
			}
			if !checkCell(board, i, 3) {
				return false
			}
			if !checkCell(board, i, 6) {
				return false
			}
		}
	}
	return true
}

func checkCell(board [][]byte, row, column int) bool {

	numMap := make(map[byte]bool)

	for i:=row;i < row + 3 ; i++ {
		for j:=column; j < column + 3;j++ {
			temp := board[i][j]
			if temp == '.' {
				continue
			}
			_, ok := numMap[temp];if ok {
				return false
			}
			numMap[temp] = true
		}
	}
	return true
}




func checkRow(board [][]byte, column int) bool {

	numMap := make(map[byte]bool)

	for j:=0; j < 9; j++ {
		temp := board[column][j]
		if temp == '.' {
			continue
		}
		_, ok := numMap[temp];if ok {
			return false
		}
		numMap[temp] = true
	}

	return true
}



func checkColumn(board [][]byte, row int) bool {

	numMap := make(map[byte]bool)

	for j:=0; j < 9; j++ {
		temp := board[j][row]
		if temp == '.' {
			continue
		}
		_, ok := numMap[temp];if ok {
			return false
		}
		numMap[temp] = true
	}

	return true
}



