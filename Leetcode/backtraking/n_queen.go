package main

import (
	"strings"
)

func main() {
	solveNQueens(4)
}

func solveNQueens(n int) [][]string {
	chase := make([][]rune, n)
	for i := 0; i < n; i++ {
		chase[i] = make([]rune, n)
		for j := 0; j < n; j++ {
			chase[i][j] = '.'
		}
	}
	result := [][]string{}
	solveNQueensHelper(0, chase, &result)
	return result
}

func solveNQueensHelper(curRowNum int, chase [][]rune, result *[][]string) {
	if curRowNum == len(chase) {
		addToResult(chase, result)
		return
	}
	for columnNum := 0; columnNum < len(chase); columnNum++ {
		if isValid(curRowNum, columnNum, chase) {
			chase[curRowNum][columnNum] = 'Q'
			solveNQueensHelper(curRowNum+1, chase, result)
			chase[curRowNum][columnNum] = '.'
		}
	}
	return
}

func isValid(curRow int, curCol int, board [][]rune) bool {
	for i := 0; i < curRow; i++ {
		if board[i][curCol] == 'Q' {
			return false
		}
	}
	for upRow, leftColumn := curRow - 1, curCol - 1; upRow >= 0 && leftColumn >= 0; upRow ,leftColumn= upRow-1, leftColumn-1 {
		if board[upRow][leftColumn] == 'Q' {
			return false
		}
	}
	for upRow,rightColumn := curRow - 1,curCol + 1 ; upRow >= 0 && rightColumn < len(board); upRow ,rightColumn = upRow-1,rightColumn+1 {
		if board[upRow][rightColumn] == 'Q' {
			return false
		}
	}
	return true
}

func addToResult(board [][]rune, res *[][]string) {
    n := len(board)
    rows := []string{}
    
    for r := 0; r < n; r++ {
        row := strings.Builder{}
        for c := 0; c < n; c++ {
            row.WriteRune(board[r][c])
        }
        rows = append(rows, row.String())
    } 
    *res = append(*res, rows)
}