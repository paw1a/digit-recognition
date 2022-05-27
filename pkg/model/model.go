package model

import (
	"encoding/gob"
	"fmt"
	"github.com/paw1a/digit-recognition/pkg/algebra"
	"os"
)

type Layer struct {
	Weights     [][]float64
	Biases      []float64
	weightedSum []float64
	activated   []float64
	activation  func(x []float64) []float64
	derivative  func(x float64) float64
}

type Model struct {
	Layers       []Layer
	LearningRate float64
	Epochs       int
	Trained      bool
}

func (m *Model) SerializeModel(filePath string, mod *Model) error {
	file, err := os.Create(filePath)
	defer file.Close()

	if err != nil {
		return fmt.Errorf("can't create serialization file %s: %v", filePath, err)
	}

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(mod)
	if err != nil {
		return fmt.Errorf("serialization error: %v", err)
	}

	return err
}

func (m *Model) DeserializeModel(filePath string, mod *Model) error {
	file, err := os.Open(filePath)
	defer file.Close()

	if err != nil {
		return fmt.Errorf("can't open deserialization file %s: %v", filePath, err)
	}

	decoder := gob.NewDecoder(file)
	err = decoder.Decode(mod)
	if err != nil {
		return fmt.Errorf("deserialization error: %v", err)
	}

	for i := 0; i < len(mod.Layers); i++ {
		mod.Layers[i].activation = ReluActivation
		mod.Layers[i].derivative = ReluDerivative
	}

	mod.Layers[len(mod.Layers)-1].activation = SoftmaxActivation
	mod.Layers[len(mod.Layers)-1].derivative = SoftmaxDerivativeStub

	return err
}

func (m *Model) FeedForward(input []float64) []float64 {
	m.Layers[0].activated = input

	for i := 0; i < len(m.Layers)-1; i++ {
		m.Layers[i].weightedSum = algebra.DotMatrixVector(m.Layers[i].Weights, m.Layers[i].activated)
		m.Layers[i].weightedSum = algebra.AddVector(m.Layers[i].weightedSum, m.Layers[i].Biases)
		m.Layers[i+1].activated = m.Layers[i+1].activation(m.Layers[i].weightedSum)
	}

	return m.Layers[len(m.Layers)-1].activated
}

func (m *Model) BackPropagation(output []float64, real []float64) {
	delta := make([]float64, len(output))

	for i := 0; i < len(delta); i++ {
		delta[i] = output[i] - real[i]
	}

	for i := len(m.Layers) - 2; i >= 0; i-- {
		for j := 0; j < len(m.Layers[i].Weights); j++ {
			for k := 0; k < len(m.Layers[i].Weights[0]); k++ {
				if j == 0 {
					m.Layers[i].Biases[k] -= m.LearningRate * delta[k]
				}
				m.Layers[i].Weights[j][k] -= m.LearningRate * delta[k] * m.Layers[i].activated[j] *
					m.Layers[i].derivative(m.Layers[i].weightedSum[k])
			}
		}

		tempDelta := delta
		delta = make([]float64, len(m.Layers[i].Weights))

		for j := 0; j < len(m.Layers[i].Weights); j++ {
			var sum float64
			for k := 0; k < len(m.Layers[i].Weights[0]); k++ {
				sum += tempDelta[k] * m.Layers[i].Weights[j][k]
			}
			delta[j] = sum
		}
	}
}

func (m Model) PredictDigit(input []float64) (digit int, output []float64) {
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

func NewModel(sizes []int, learningRate float64, epochs int) *Model {
	model := &Model{
		Layers:       make([]Layer, len(sizes)),
		LearningRate: learningRate,
		Epochs:       epochs,
	}

	for i := 0; i < len(sizes)-1; i++ {
		model.Layers[i] = Layer{
			Weights:    algebra.RandomMatrix(sizes[i], sizes[i+1]),
			Biases:     algebra.RandomVector(sizes[i+1]),
			activation: ReluActivation,
			derivative: ReluDerivative,
		}
	}

	model.Layers[len(model.Layers)-1].activation = SoftmaxActivation
	model.Layers[len(model.Layers)-1].derivative = SoftmaxDerivativeStub

	return model
}
