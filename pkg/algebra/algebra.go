package algebra

import (
	"math/rand"
	"time"
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
	rand.Seed(time.Now().Unix())

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
	rand.Seed(time.Now().Unix())

	vector := make([]float64, size)

	for i := 0; i < size; i++ {
		vector[i] = rand.Float64()*2 - 1
	}

	return vector
}

func DotMatrixVector(matrix [][]float64, vector []float64) []float64 {
	if len(matrix) == 0 || len(matrix) != len(vector) {
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

func Max(n1 int, n2 int) int {
	if n1 > n2 {
		return n1
	} else {
		return n2
	}
}

func Min(n1 int, n2 int) int {
	if n1 < n2 {
		return n1
	} else {
		return n2
	}
}
