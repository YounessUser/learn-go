// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math"
)

func InitMatrix(row int, col int) [][]int {
	matrix := make([][]int, row)
	for i := 0; i < row; i++ {
		matrix[i] = make([]int, col)
		for j := 0; j < col; j++ {
			fmt.Printf("inset the value of %vth row and %vth column : ", i, j)
			fmt.Scan(&matrix[i][j])
		}
	}

	return matrix
}

func DisplayMatrix(matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		fmt.Print("( ")
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf(" %v ", matrix[i][j])
		}
		fmt.Println(" )")
	}
}

func MultiplyMarix(matrix1 [][]int, matrix2 [][]int) {
	if len(matrix1) != len(matrix2[0]) && len(matrix1[0]) != len(matrix2) {
		return
	}
	matrix := make([][]int, len(matrix1))
	sum := 0
	for i := 0; i < len(matrix1); i++ {
		matrix[i] = make([]int, len(matrix2[0]))
		for j := 0; j < len(matrix2[0]); j++ {
			sum = 0
			for k := 0; k < len(matrix1[0]); k++ {
				sum += (matrix1[i][k] * matrix2[k][j])
			}
			matrix[i][j] = sum
		}
	}

	fmt.Println("result of multiplying : ")
	DisplayMatrix(matrix)
}

// https://www.youtube.com/watch?v=QGYvbsHDPxo
func Strassen(matrix1 [][]int, matrix2 [][]int) {
	if len(matrix1) != len(matrix2[0]) && len(matrix1[0]) != len(matrix2) {
		return
	}

	var maxLine float64 = math.Max(float64(len(matrix1)), float64(len(matrix2)))
	if CheckPower2(maxLine) && maxLine != 2 {

	} else {

	}

	matrix := make([][]int, len(matrix1))
	p1 := (matrix1[0][0] + matrix1[1][1]) * (matrix2[0][0] + matrix2[1][1])
	p2 := (matrix1[1][0] + matrix1[1][1]) * matrix2[0][0]
	p3 := matrix1[0][0] * (matrix2[0][1] - matrix2[1][1])
	p4 := matrix1[1][1] * (matrix2[1][0] - matrix2[0][0])
	p5 := (matrix1[0][0] + matrix1[0][1]) * matrix2[1][1]
	p6 := (matrix1[1][0] - matrix1[0][0]) * (matrix2[0][0] + matrix2[0][1])
	p7 := (matrix1[0][1] - matrix1[1][1]) * (matrix2[1][0] + matrix2[1][1])

	matrix[0][0] = p1 + p4 - p5 + p7
	matrix[0][1] = p3 + p5
	matrix[1][0] = p2 + p4
	matrix[1][1] = p1 + p3 - p2 + p6
}

func CheckPower2(num float64) bool {
	if num == 0 {
		return false
	}
	return math.Floor(math.Log2(num)) == math.Ceil(math.Log2(num))
}

func main() {
	var row int = 0
	var col int = 0

	fmt.Println("inset the your matrix size rows then cols : ")
	_, err := fmt.Scan(&row, &col)
	if err != nil {
		panic(err)
	}

	fmt.Println("Define first matrix : ")
	var matrix [][]int = InitMatrix(row, col)
	DisplayMatrix(matrix)

	fmt.Println("inset the your matrix size rows then cols : ")
	_, err2 := fmt.Scan(&row, &col)
	if err2 != nil {
		panic(err2)
	}

	fmt.Println("Define first matrix : ")
	var matrix2 [][]int = InitMatrix(row, col)
	DisplayMatrix(matrix2)

	MultiplyMarix(matrix, matrix2)
}
