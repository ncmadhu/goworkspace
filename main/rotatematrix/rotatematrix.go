package rotatematrix

import "fmt"

//Rotate - rotate given matrix 90 deg clockwise
func Rotate(matrix [][]int) {
	n := len(matrix)
	i := 0
	var temp int
	for i < n/2 {
		j := 0
		for j < n-n/2 {
			matrix[n-1-(n-1-j)][n-1-i], temp = matrix[i][j], matrix[n-1-(n-1-j)][n-1-i]
			matrix[n-1-i][n-1-j], temp = temp, matrix[n-1-i][n-1-j]
			matrix[n-1-j][i], temp = temp, matrix[n-1-j][i]
			matrix[i][j] = temp
			j++
		}
		i++
	}
	fmt.Println(matrix)
}
