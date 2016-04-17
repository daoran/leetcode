// Source : https://oj.leetcode.com/problems/word-search/
// Author : daoran
// Date   : 2016-04-17

/**********************************************************************************
*
* Given a 2D board and a word, find if the word exists in the grid.
*
* The word can be constructed from letters of sequentially adjacent cell,
* where "adjacent" cells are those horizontally or vertically neighboring.
* The same letter cell may not be used more than once.
*
* For example,
* Given board =
*
* [
*   ["ABCE"],
*   ["SFCS"],
*   ["ADEE"]
* ]
*
* word = "ABCCED", -> returns true,
* word = "SEE", -> returns true,
* word = "ABCB", -> returns false.
*
*
**********************************************************************************/

package main

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(word) == 0 {
		return false
	}

	row := len(board)
	col := len(board[0])

	mask := make([][]int, row)
	for i := 0; i < len(mask); i++ {
		mask[i] = make([]int, col)
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == word[0] {
				m := mask
				if existHelper(board, word, 0, i, j, m) {
					return true
				}
			}
		}
	}

	return false
}

func existHelper(board [][]byte, word string, idx int, row int, col int, mask [][]int) bool {
	i := row
	j := col

	if board[i][j] == word[idx] && mask[i][j] == 0 {
		mask[i][j] = 1

		if idx+1 == len(word) {
			return true
		}

		idx++
		if (i+1 < len(board) && existHelper(board, word, idx, i+1, j, mask)) ||
			(i > 0 && existHelper(board, word, idx, i-1, j, mask)) ||
			(j+1 < len(board[i]) && existHelper(board, word, idx, i, j+1, mask)) ||
			(j > 0 && existHelper(board, word, idx, i, j-1, mask)) {
			return true
		}

		mask[i][j] = 0 //cannot find any successful solution, clear the mark. (backtracking)

	}

	return false
}

func main() {

}
