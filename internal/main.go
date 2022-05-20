package main

import (
	"fmt"
	"github.com/paw1a/digit-recognition/pkg/algebra"
)

func main() {
	matrix := algebra.RandomMatrix(10, 10)
	vector := make([]float64, 10)
	for i := 0; i < 10; i++ {
		vector[i] = float64(i * 10)
	}

	fmt.Printf("%v", algebra.DotMatrixVector(matrix, vector))
}
