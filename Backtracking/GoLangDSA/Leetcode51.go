package main

func solve(coloumn int, result *[][]string, ans []string, leftRow []int, upperDiagonal []int, lowerDiagnal []int, n int) {

	if coloumn == n {

		subsetcpy := append([]string{}, ans...)
		*result = append(*result, subsetcpy)
		return
	}

	for row := 0; row < n; row++ {

		if leftRow[row] == 0 && lowerDiagnal[row+coloumn] == 0 && upperDiagonal[n-1+coloumn-row] == 0 {

			rowChars := []rune(ans[row])
			rowChars[coloumn] = 'Q'
			ans[row] = string(rowChars)

			leftRow[row] = 1
			lowerDiagnal[row+coloumn] = 1
			upperDiagonal[n-1+coloumn-row] = 1
			solve(coloumn+1, result, ans, leftRow, upperDiagonal, lowerDiagnal, n)

			rowChars[coloumn] = '.'
			ans[row] = string(rowChars)
			leftRow[row] = 0
			lowerDiagnal[row+coloumn] = 0
			upperDiagonal[n-1+coloumn-row] = 0
		}
	}
}
func solveNQueens(n int) [][]string {

	var result [][]string
	ans := make([]string, n)

	for i := 0; i < n; i++ {
		ans[i] = string(make([]rune, n, n))
		for j := range ans[i] {
			ans[i] = ans[i][:j] + "." + ans[i][j+1:]
		}
	}

	leftRow := make([]int, n)
	upperDiagnal := make([]int, 2*n-1)
	lowerDiagnal := make([]int, 2*n-1)
	solve(0, &result, ans, leftRow, upperDiagnal, lowerDiagnal, n)

	return result
}
