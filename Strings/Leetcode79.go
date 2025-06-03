package main

import "fmt"

func wordMatchInMatrix(board [][]byte, word,match string, row,coloumn int) bool {

	if row == len(board) - 1 && coloumn == len(board) - 1 {

		if word == matched {

			return true
		}
	}

	if coloumn == len(board) - 1 {

		match += string(board[row,coloumn])
		wordMatchInMatrix(board,word,match,row + 1,0)
	}else {

		match += string(board[i][j])
	    wordMatchInMatrix(word,match,row,coloumn + 1)
	}

	return false
	
}

func exist(board [][]byte, word string) bool {

	if len(board) == 0 {

		return false
	}

	var result bool

	result = wordMatchInMatrix(board,word,"",0,0)

	return result
}
