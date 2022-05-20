package algebra

import "math/rand"

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
		return nil
	}

	result := make([]float64, len(matrix))

	for i := 0; i < len(matrix); i++ {
		var sum float64
		for j := 0; j < len(matrix[0]); j++ {
			sum += matrix[i][j] * vector[i]
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
