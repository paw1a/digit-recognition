package model

import (
	"github.com/paw1a/digit-recognition/pkg/algebra"
)

type layer struct {
	weights     [][]float64
	biases      []float64
	weightedSum []float64
	activated   []float64
	activation  func(x []float64) []float64
	derivative  func(x float64) float64
}

type model struct {
	layers       []layer
	learningRate float64
	epochs       int
}

func (m *model) FeedForward(input []float64) []float64 {
	m.layers[0].activated = input

	for i := 0; i < len(m.layers)-1; i++ {
		m.layers[i].weightedSum = algebra.DotMatrixVector(m.layers[i].weights, m.layers[i].activated)
		m.layers[i].weightedSum = algebra.AddVector(m.layers[i].weightedSum, m.layers[i].biases)
		m.layers[i+1].activated = m.layers[i+1].activation(m.layers[i].weightedSum)
	}

	return m.layers[len(m.layers)-1].activated
}

func (m *model) BackPropagation(output []float64, real []float64) {
	delta := make([]float64, len(output))

	for i := 0; i < len(delta); i++ {
		delta[i] = output[i] - real[i]
	}

	for i := len(m.layers) - 2; i >= 0; i-- {
		for j := 0; j < len(m.layers[i].weights); j++ {
			for k := 0; k < len(m.layers[i].weights[0]); k++ {
				if j == 0 {
					m.layers[i].biases[k] -= m.learningRate * delta[k]
				}
				m.layers[i].weights[j][k] -= m.learningRate * delta[k] * m.layers[i].activated[j] *
					m.layers[i].derivative(m.layers[i].weightedSum[k])
			}
		}

		tempDelta := delta
		delta = make([]float64, len(m.layers[i].weights))

		for j := 0; j < len(m.layers[i].weights); j++ {
			var sum float64
			for k := 0; k < len(m.layers[i].weights[0]); k++ {
				sum += tempDelta[k] * m.layers[i].weights[j][k]
			}
			delta[j] = sum
		}
	}
}

func (m model) PredictDigit(input []float64) (digit int, output []float64) {
	output = m.FeedForward(input)

	maxProbability := output[0]
	maxProbabilityIndex := 0
	for k := 0; k < len(output); k++ {
		if output[k] > maxProbability {
			maxProbability = output[k]
			maxProbabilityIndex = k
		}
	}

	return maxProbabilityIndex, output
}

func NewModel(sizes []int, learningRate float64, epochs int) *model {
	model := &model{
		layers:       make([]layer, len(sizes)),
		learningRate: learningRate,
		epochs:       epochs,
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
	model.layers[len(model.layers)-1].derivative = SoftmaxDerivativeStub

	return model
}
