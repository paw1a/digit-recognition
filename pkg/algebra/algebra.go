package algebra

import (
	"fmt"
	"math/rand"
)

func AddVector(vector1 []float64, vector2 []float64) []float64 {
	if len(vector1) != len(vector2) {
		return nil
	}

	result := make([]float64, len(vector1))

	for i := 0; i < len(vector1); i++ {
		result[i] = vector1[i] + vector2[i]
	}

	return result
}

func RandomMatrix(height int, width int) [][]float64 {
	matrix := make([][]float64, height)

	for i := 0; i < height; i++ {
		matrix[i] = make([]float64, width)
		for j := 0; j < width; j++ {
			matrix[i][j] = rand.Float64()*2 - 1
		}
	}

	return matrix
}

func RandomVector(size int) []float64 {
	vector := make([]float64, size)

	for i := 0; i < size; i++ {
		vector[i] = rand.Float64()*2 - 1
	}

	return vector
}

func DotMatrixVector(matrix [][]float64, vector []float64) []float64 {
	if len(matrix) == 0 || len(matrix) != len(vector) {
		fmt.Printf("err")
		return nil
	}

	result := make([]float64, len(matrix[0]))

	for i := 0; i < len(matrix[0]); i++ {
		var sum float64
		for j := 0; j < len(matrix); j++ {
			sum += matrix[j][i] * vector[j]
		}
		result[i] = sum
	}

	return result
}

func DotVector(vector1 []float64, vector2 []float64) []float64 {
	if len(vector1) != len(vector2) {
		return nil
	}

	result := make([]float64, len(vector1))

	for i := 0; i < len(vector1); i++ {
		result[i] = vector1[i] * vector2[i]
	}

	return result
}

func MultiplyMatrix(matrix [][]float64, number float64) {
	if len(matrix) == 0 {
		return
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[i][j] *= number
		}
	}
}

func MultiplyVectors(vector1 []float64, vector2 []float64) [][]float64 {
	matrix := make([][]float64, len(vector1))

	for i := 0; i < len(vector1); i++ {
		matrix[i] = make([]float64, len(vector2))
		for j := 0; j < len(vector2); j++ {
			matrix[i][j] = vector1[i] * vector2[j]
		}
	}

	return matrix
}

func MultiplyVector(vector []float64, number float64) []float64 {
	result := make([]float64, len(vector))

	for i := 0; i < len(vector); i++ {
		result[i] = vector[i] * number
	}

	return result
}

func DiffVector(vector []float64, delta []float64) {
	if len(vector) != len(delta) {
		return
	}

	for i := 0; i < len(vector); i++ {
		vector[i] -= delta[i]
	}
}

func DiffMatrix(matrix1 [][]float64, matrix2 [][]float64) {
	if len(matrix1) == 0 {
		return
	}

	for i := 0; i < len(matrix1); i++ {
		for j := 0; j < len(matrix1[0]); j++ {
			matrix1[i][j] -= matrix2[i][j]
		}
	}
}
