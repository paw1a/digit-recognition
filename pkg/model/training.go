package model

import (
	"fmt"
	"math"
	"math/rand"
)

func (m *model) Fit(input [][]float64, real [][]float64) {
	if len(input) != len(real) {
		return
	}

	for i := 0; i < m.epochs; i++ {
		var loss float64
		var correct int

		for j := 0; j < len(input); j++ {
			inputIndex := rand.Intn(len(input))

			digit, output := m.PredictDigit(input[inputIndex])

			if real[inputIndex][digit] == 1 {
				correct++
			}

			for k := 0; k < len(output); k++ {
				if real[inputIndex][k] == 1 {
					loss += -math.Log(output[k])
					break
				}
			}

			m.BackPropagation(output, real[inputIndex])
		}
		printLearningStat(i+1, loss/float64(len(input)), float64(correct)/float64(len(input)))
	}
}

func (m *model) TestModel(input [][]float64, real [][]float64) {
	var correct int
	var loss float64

	for i := 0; i < len(input); i++ {
		output := m.FeedForward(input[i])

		maxProbability := output[0]
		maxProbabilityIndex := 0
		for k := 0; k < len(output); k++ {
			if output[k] > maxProbability {
				maxProbability = output[k]
				maxProbabilityIndex = k
			}
		}

		if real[i][maxProbabilityIndex] == 1 {
			correct++
		}

		for k := 0; k < len(output); k++ {
			if real[i][k] == 1 {
				loss += -math.Log(output[k])
				break
			}
		}
	}

	printLearningStat(100, loss/float64(len(input)), float64(correct)/float64(len(input)))
}

func printLearningStat(epoch int, loss float64, accuracy float64) {
	fmt.Printf("Epoch: %d. Loss = %f. Accuracy = %f\n", epoch, loss, accuracy)
}
