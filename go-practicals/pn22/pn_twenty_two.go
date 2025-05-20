package pn22

import "fmt"

func SearchIn2DMatrix() {
	var matrix [][]int = 
	element, row, column := fetchElementFrom2DMatrix(matrix, 10)
	fmt.Println("element: ",element," found in 2D matrix ",matr)

}

func fetchElementFrom2DMatrix(matrix [][]int, element int) (int, int, int) {
	for i := 0; i <= len(matrix); i++ {
		for j := 0; j <= len(matrix[i]); j++ {
			if matrix[i][j] == element {
				return matrix[i][j], i, j
			}
		}
	}
	return 0, -1, -1
}
