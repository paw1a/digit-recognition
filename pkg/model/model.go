package model

import (
	"github.com/paw1a/digit-recognition/pkg/algebra"
)

type layer struct {
	weights    [][]float64
	biases     []float64
	activation func(x []float64) []float64
	derivative func(x []float64) []float64
}

type model struct {
	layers []layer
}

func NewModel(sizes []int) *model {
	model := &model{
		layers: make([]layer, len(sizes)),
	}

	for i := 0; i < len(sizes)-1; i++ {
		model.layers[i] = layer{
			weights:    algebra.RandomMatrix(sizes[i], sizes[i+1]),
			biases:     algebra.RandomVector(sizes[i+1]),
			activation: ReluActivation,
			derivative: ReluDerivative,
		}
	}

	model.layers[len(model.layers)-1].activation = SoftmaxActivation

	return model
}

func (m *model) FeedForward(inputs []float64) []float64 {
	activated := inputs

	for i := 0; i < len(m.layers)-1; i++ {
		weightedSum := algebra.DotMatrixVector(m.layers[i].weights, activated)
		weightedSum = algebra.AddVector(weightedSum, m.layers[i].biases)
		activated = m.layers[i+1].activation(weightedSum)
	}

	return activated
}
