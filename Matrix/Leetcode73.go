package main

func fillZeroes(i int, j int, matrix *[][]int) {
    

    if matrix == nil || i < 0 || j < 0 {

        return
    }

    rows := len((*matrix))
    cols := len((*matrix)[0])

	for col := 0; col < cols; col++ {
         
		(*matrix)[i][col] = 0
	}
    
    for row := 0; row < rows; row++ {

		(*matrix)[row][j] = 0
	}
    
}

func setZeroes(matrix [][]int) {

	if len(matrix) == 0 {

		return
	}

    type Position struct {

        i,j int
    }
    
    var Zeroes []Position

	for i := 0; i < len(matrix); i++ {

		for j := 0; j < len(matrix[i]); j++ {

			if matrix[i][j] == 0 {

				Zeroes = append(Zeroes,Position{i,j})
			}
		}
	}

    for _,pos := range Zeroes {

        fillZeroes(pos.i,pos.j,&matrix)
    }
}
