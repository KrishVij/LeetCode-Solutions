package main

import (

	"fmt"

)

func rotate(matrix [][]int) {

	if len(matrix) == 0 {

		return
	}
	
	for j := 0; j < len(matrix[0]); j++ {

		for i := 0; i < len(matrix)/2; i++ {

			matrix[i][j],matrix[len(matrix) - 1 - i][j] = matrix[len(matrix) - 1 - i][j],matrix[i][j]
		}
	}
    
	
	for i := 0; i < len(matrix); i++ {

		for j := i + 1; j < len(matrix[i]); j++ {
			
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		
		}
	}

	for i := 0; i < len(matrix); i++ {

		for j := 0; j < len(matrix[i]); j++ {

			fmt.Print(matrix[i][j])
		}
	}

}
