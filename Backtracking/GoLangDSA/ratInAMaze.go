package main

import "fmt"

func startMoving(maze [][]int, marked [][]int, allPathStore string, i, j int) {

	if maze[i][j] == maze[len(maze)-1][len(maze)-1] {

		fmt.Print(allPathStore)
		return
	}

	for i+1 < len(maze) && marked[i+1][j] == 0 && maze[i+1][j] == 1 {

		marked[i][j] = 1
		startMoving(maze, marked, allPathStore+"D", i+1, j)
		marked[i][j] = 0
	}

	for j-1 >= 0 && marked[i][j-1] == 0 && maze[j-1][i] == 1 {

		marked[i][j] = 1
		startMoving(maze, marked, allPathStore+"L", i, j-1)
		marked[i][j] = 0
	}

	for j+1 < len(maze) && marked[i][j+1] == 0 && maze[i][j+1] == 1 {

		marked[i][j] = 1
		startMoving(maze, marked, allPathStore+"R", i, j+1)
		marked[i][j] = 0
	}

	for i-1 >= 0 && marked[i-1][j] == 0 && maze[i-1][j] == 1 {

		marked[i][j] = 1
		startMoving(maze, marked, allPathStore+"R", i-1, j)
	}

}
func ratInAMaze(maze [][]int) string {

	var marked [][]int
	s := ""

	if maze[0][0] == 1 {

		startMoving(maze, marked, s, 0, 0)
	}

	return s
}
