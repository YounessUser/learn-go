package main

import (
	"fmt"
	"math"
	"os"
)

type Matrix struct {
  rows int
  cols int
  data [][]float64
}

var matrixes = make(map[string]Matrix)

func InitMatrixes(){
	var row int = 0
	var col int = 0

	fmt.Print("inset the first matrix size rows then cols : ")
	_, err := fmt.Scan(&row, &col)
	if err != nil {
		panic(err)
	}
  	var matrix1 Matrix = Matrix{rows: row, cols: col}
	fmt.Println("Define first matrix : ")
	matrix1.data = fillMatrix(row, col)
	DisplayMatrix(matrix1)
	matrixes["matrix1"] = matrix1

	fmt.Print("inset the second matrix size rows then cols : ")
	_, err2 := fmt.Scan(&row, &col)
	if err2 != nil {
		panic(err2)
	}
  	var matrix2 Matrix = Matrix{rows: row, cols: col}
	fmt.Println("Define second matrix : ")
	matrix2.data = fillMatrix(row, col)
	DisplayMatrix(matrix2)
	matrixes["matrix2"] = matrix2
}

func fillMatrix(row int , col int) [][]float64 {
	matrix := make([][]float64, row)
	for i := 0; i < row; i++ {
		matrix[i] = make([]float64, col)
		for j := 0; j < col; j++ {
			fmt.Printf("inset the value of %vth row and %vth column : ", i, j)
			fmt.Scan(&matrix[i][j])
		}
	}

	return matrix
}

func DisplayMatrix(matrix Matrix) {
	for i := 0; i < matrix.rows; i++ {
		fmt.Print("( ")
		for j := 0; j < matrix.cols; j++ {
			fmt.Printf(" %v ", matrix.data[i][j])
		}
		fmt.Println(" )")
	}
}

func AddMarixes(matrix1 Matrix, matrix2 Matrix) Matrix {
	if matrix1.rows != matrix2.rows && matrix1.cols != matrix2.cols {
		return Matrix{}
	}

	var matrix Matrix = Matrix{rows: matrix1.rows, cols: matrix1.cols, data: make([][]float64, matrix1.rows)}
	for i := 0; i < matrix1.rows; i++ {
		matrix.data[i] = make([]float64, matrix1.cols)
		for j := 0; j < matrix1.cols; j++ {
			matrix.data[i][j] = matrix1.data[i][j] + matrix2.data[i][j]
		}
	}

	return matrix
}

func SubstractMarixes(matrix1 Matrix, matrix2 Matrix) Matrix {
	if matrix1.rows != matrix2.rows && matrix1.cols != matrix2.cols {
		return Matrix{}
	}

	var matrix Matrix = Matrix{rows: matrix1.rows, cols: matrix1.cols, data: make([][]float64, matrix1.rows)}
	for i := 0; i < matrix1.rows; i++ {
		matrix.data[i] = make([]float64, matrix1.cols)
		for j := 0; j < matrix1.cols; j++ {
			matrix.data[i][j] = matrix1.data[i][j] - matrix2.data[i][j]
		}
	}

	return matrix
}

func MultiplyMarix(matrix1 Matrix, matrix2 Matrix) Matrix {
	if matrix1.rows != matrix2.cols || matrix1.cols != matrix2.rows {
		return Matrix{}
	}
	matrix := Matrix{rows: matrix1.rows, cols: matrix2.cols, data: make([][]float64, matrix1.rows)}
	sum := 0.0
	for i := 0; i < matrix1.rows; i++ {
		matrix.data[i] = make([]float64, matrix2.cols)
		for j := 0; j < matrix2.cols; j++ {
			sum = 0
			for k := 0; k < matrix1.cols; k++ {
				sum += (matrix1.data[i][k] * matrix2.data[k][j])
			}
			matrix.data[i][j] = sum
		}
	}

	return matrix
}

func DevideMarix(matrix1 Matrix, matrix2 Matrix) Matrix {
	if matrix1.rows != matrix2.cols || matrix1.cols != matrix2.rows {
		return Matrix{}
	}
	
	inversedMatrix2 := InverseMatrix(matrix2)
	matrix := MultiplyMarix(matrix1, inversedMatrix2)

	return matrix
}


func DeterminantMatrix(matrix Matrix) float64 {
	if matrix.rows != matrix.cols {
		return 0
	}

	detMatrix := 0.0;

	if matrix.rows == 2 {
		detMatrix = DetMatrixWithDementionTwo(matrix.data);
		return detMatrix
	}
	
	for i := 0; i < matrix.rows; i++ {
		result := MatrixMinor(matrix.data, 1, i)
		detMatrix += math.Pow(-1, float64(i)) * DetMatrixWithDementionTwo(result)
	}
	return detMatrix
}

func MatrixMinor(matrix [][]float64, detRow int, detCol int) [][]float64 {
	var size = len(matrix)
	var result [][]float64 = make([][]float64, size - 1)
	k := 0
	for i := 0; i < size; i++ {
		if i == detRow { continue }
		result[k] = make([]float64, size - 1)
		d := 0
		for j := 0; j < size ; j++ {
			if j == detCol { continue }
			result[k][d] = matrix[i][j];
			d++;
		}
		k++;
	}
	return result
}

func DetMatrixWithDementionTwo(matrix [][]float64) float64 {
	return  matrix[0][0]*matrix[1][1] - matrix[0][1]*matrix[1][0]
}

func GetMatrixMinors(matrix Matrix) Matrix {
	matrixMinors := Matrix{rows: matrix.rows, cols: matrix.cols, data: make([][]float64, matrix.rows)}
	for i := 0; i < matrix.rows; i++ {
		matrixMinors.data[i] = make([]float64, matrix.cols)
		for j := 0; j < matrix.cols ; j++ {
			result := MatrixMinor(matrix.data, i, j)
			matrixMinors.data[i][j] = math.Pow(-1, float64(i+j)) * DetMatrixWithDementionTwo(result)
		}
	}
	return matrixMinors
}

func InverseMatrix(matrix Matrix) Matrix {
	var det float64 = DeterminantMatrix(matrix)
	if ( det == 0 ) {
		return Matrix{}
	}

	inversedMatrix := Matrix{rows: matrix.rows, cols: matrix.cols, data: make([][]float64, matrix.rows)}
	result := GetMatrixMinors(matrix)
	adjugateMatrix := TransposeMatrix(result)
	fmt.Println("adjugate matrix : ")
	DisplayMatrix(adjugateMatrix)

	for i := 0; i < adjugateMatrix.rows; i++ {
		inversedMatrix.data[i] = make([]float64, matrix.cols)
		for j := 0; j < adjugateMatrix.cols ; j++ {
			inversedMatrix.data[i][j] = (1/det) * float64(adjugateMatrix.data[i][j])
		}
	}

	return inversedMatrix
}

func TransposeMatrix(matrix Matrix) Matrix {
	transposedMatrix := Matrix{rows: matrix.cols, cols: matrix.rows, data: make([][]float64, matrix.cols)}

	for i := 0; i < matrix.cols; i++ {
		transposedMatrix.data[i] = make([]float64, matrix.rows)
		for j := 0; j < matrix.rows; j++ {
			transposedMatrix.data[i][j] = matrix.data[j][i]
		}
	}

	return transposedMatrix
}

func menu() (choice int) {
	fmt.Println("Welcome to matrix math app :")
	fmt.Println("0) Exit")
	fmt.Println("1) Add")
	fmt.Println("2) Substract")
	fmt.Println("3) devide")
	fmt.Println("4) Multiply")
	fmt.Println("5) Transpose")
	fmt.Println("6) Inverse")
	fmt.Println("7) Det")

	fmt.Println("Choose your operaion :")
	_, err := fmt.Scanf("%d", &choice)
	if err != nil {
		panic(err)
	}
	return
}

func main() {
	for {
		var choice = menu()

		if (len(matrixes) == 0 && choice != 0) {
			InitMatrixes()
		}
		switch (choice) {
			case 1:
				result := AddMarixes(matrixes["matrix1"], matrixes["matrix2"])
				fmt.Println("result of Adding matrixes : ")
				DisplayMatrix(result)
			case 2:
				result := SubstractMarixes(matrixes["matrix1"], matrixes["matrix2"])
				fmt.Println("result of Substracting matrixes : ")
				DisplayMatrix(result)
			case 3:
				result := DevideMarix(matrixes["matrix1"], matrixes["matrix2"])
				fmt.Println("result of deviding : ")
				DisplayMatrix(result)
			case 4:
				result := MultiplyMarix(matrixes["matrix1"], matrixes["matrix2"])
				fmt.Println("result of multiplying : ")
				DisplayMatrix(result)
			case 5:
				result1 := TransposeMatrix(matrixes["matrix1"])
				fmt.Println("result of Transposed matrix 1 : ")
				DisplayMatrix(result1)

				result2 := TransposeMatrix(matrixes["matrix2"])
				fmt.Println("result of Transposed matrix 2 : ")
				DisplayMatrix(result2)
			case 6:
				result1 := InverseMatrix(matrixes["matrix1"])
				fmt.Println("result of Inversing matrix1 : ")
				DisplayMatrix(result1)

				result2 := InverseMatrix(matrixes["matrix2"])
				fmt.Println("result of Inversing matrix2 : ")
				DisplayMatrix(result2)
			case 7:
				fmt.Println("Det(matrix1) : ", DeterminantMatrix(matrixes["matrix1"]))
				fmt.Println("Det(matrix2) : ", DeterminantMatrix(matrixes["matrix2"]))
			default:
				os.Exit(1)
		}	
	}
}