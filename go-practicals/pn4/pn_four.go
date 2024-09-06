package pn4

import "fmt"

// Practical4 Description
/**
@initialAuthor: monks_mojo
@lastUpdatedBY: monks_mojo
Practical Four:
Complete the createMatrix function. It takes a number of rows and columns and returns a 2D slice of integers where the value of each cell is i * j where i and j are the indexes of the row and column respectively.
For example, a 10x10 matrix would look like this:

input row int -> 10, column int -> 10

output row X column matrix

*/

// The function "CreateMatrix" prompts the user to enter the number of rows and columns for a matrix
// and then calls the "create2DMatrix" function.
func CreateMatrix() {
	row, column := 0, 0
	fmt.Println("Enter the total number of rows of the matrix")
	fmt.Scanln(&row)
	fmt.Println("Enter the total number of columns of the matrix")
	fmt.Scanln(&column)
	create2DMatrix(row, column)
}

// The function creates a 2D matrix of size rowNo x columnNo and fills it with values obtained by
// multiplying the row and column indices.
func create2DMatrix(rowNo, columnNo int) {
	var matrix [][]int

	for i := 0; i < rowNo; i++ {
		var row []int
		for j := 0; j < columnNo; j++ {
			row = append(row, i*j)
		}
		matrix = append(matrix, row)
	}
	fmt.Printf("a %d X %d matrix will be \n", rowNo, columnNo)
	fmt.Println(matrix)
}
