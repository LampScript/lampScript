package main

import (
	"fmt"
)

/*
这题没做出来，有点惨

*/
//https://leetcode.com/problems/sudoku-solver/discuss/15752/Straight-Forward-Java-Solution-Using-Backtracking
func main() {
	c := [][]byte{
		{'5','3','.','.','7','.','.','.','.'},
		{'6','.','.','1','9','5','.','.','.'},
		{'.','9','8','.','.','.','.','6','.'},
		{'8','.','.','.','6','.','.','.','3'},
		{'4','.','.','8','.','3','.','.','1'},
		{'7','.','.','.','2','.','.','.','6'},
		{'.','6','.','.','.','.','2','8','.'},
		{'.','.','.','4','1','9','.','.','5'},
		{'.','.','.','.','8','.','.','7','9'}}
	solveSudoku(c)
	for i:=0;i<len(c);i++{
		fmt.Println(c[i])
	}

}

var bigCellExistMap [3][3]map[byte]bool
var rowExistMap [9]map[byte]bool
var columnExistMap [9]map[byte]bool
var all = []byte{'1','2','3','4','5','6','7','8','9'}
var repeat bool
var total int

func solveSudoku(board [][]byte)  {
	initAllMap(board)
	repeat=true
	for repeat {
		repeat = false
		for i:=0;i<9;i++  {
			for j:=0;j<9;j++ {
				//if board[i][j] == '.' {
				//	continue
				//}
				temp := getPossibleNum(j, i)
				if len(temp) == 1 {
					value := temp[0]
					rowExistMap[i][value] = true
					columnExistMap[j][value] = true
					bigCellExistMap[i/3][j/3][value] = true
					board[i][j] = value
					total++
					repeat = true
				}
				if i==0 && j == 8 {
					getPossibleNum(0,8)
				}
			}
		}
	}
}

func getPossibleNum(column, row int) []byte {
	rowMap := rowExistMap[row]
	columnMap := columnExistMap[column]
	bigCellMap := bigCellExistMap[column/3][row/3]

	possible := make([]byte,0, 9)
	for i:=0;i<len(all);i++{
		_,ok := rowMap[all[i]];if ok {
			continue
		}
		_,ok = columnMap[all[i]];if ok {
			continue
		}
		_,ok = bigCellMap[all[i]];if ok {
			continue
		}
		possible = append(possible, all[i])
	}
	
	return possible
}

func initAllMap(board [][]byte)  {

	for i:=0;i<9;i++  {
		for j:=0;j<9;j++ {
			value := board[i][j]
			if value == '.' {
				continue
			}
			total++
			if rowExistMap[i] == nil {
				rowExistMap[i] = make(map[byte]bool, 9)
			}
			rowExistMap[i][value] = true
			if columnExistMap[j] == nil {
				columnExistMap[j] = make(map[byte]bool, 9)
			}
			columnExistMap[j][value] = true
			if bigCellExistMap[i/3][j/3] == nil {
				bigCellExistMap[i/3][j/3] = make(map[byte]bool, 9)
			}
			bigCellExistMap[i/3][j/3][value] = true
		}
	}
}

