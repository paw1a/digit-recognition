package model

import (
	"math"
)

func ReluActivation(x []float64) []float64 {
	result := make([]float64, len(x))

	for i := 0; i < len(x); i++ {
		if x[i] <= 0 {
			result[i] = 0
		} else {
			result[i] = x[i]
		}
	}

	return result
}

func ReluDerivative(x float64) float64 {
	if x <= 0 {
		return 0
	} else {
		return 1
	}
}

func SoftmaxActivation(x []float64) []float64 {
	result := make([]float64, len(x))
	var sum float64

	for i := 0; i < len(x); i++ {
		sum += math.Exp(x[i])
	}

	for i := 0; i < len(x); i++ {
		result[i] = math.Exp(x[i]) / sum
	}

	return result
}

func SoftmaxDerivativeStub(x float64) float64 {
	return 1
}
