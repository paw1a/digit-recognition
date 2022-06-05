package model

import (
	"log"
	"math"
	"math/rand"
)

func (m *Model) Fit(input [][]float64, labels [][]float64) {
	m.TrainingState.DatasetSize = len(input)

	if len(input) != len(labels) {
		return
	}

	for i := 0; i < m.Epochs; i++ {
		var loss float64
		var correct int

		m.TrainingState.CurrentEpoch = i
		m.TrainingState.CurrentLoss = loss

		for j := 0; j < len(input); j++ {
			m.TrainingState.CurrentIteration = j + 1

			inputIndex := rand.Intn(len(input))

			digit, output := m.PredictDigit(input[inputIndex])

			if labels[inputIndex][digit] == 1 {
				correct++
			}

			for k := 0; k < len(output); k++ {
				if labels[inputIndex][k] == 1 {
					loss += -math.Log(output[k])
					break
				}
			}

			m.BackPropagation(output, labels[inputIndex])
		}
		printLearningStat(i+1, loss/float64(len(input)), float64(correct)/float64(len(input)))
	}

	m.Trained = true
}

func (m *Model) TestModel(input [][]float64, labels [][]float64) {
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

		if labels[i][maxProbabilityIndex] == 1 {
			correct++
		}

		for k := 0; k < len(output); k++ {
			if labels[i][k] == 1 {
				loss += -math.Log(output[k])
				break
			}
		}
	}

	printLearningStat(100, loss/float64(len(input)), float64(correct)/float64(len(input)))
}

func printLearningStat(epoch int, loss float64, accuracy float64) {
	log.Printf("Epoch: %d. Loss = %f. Accuracy = %f\n", epoch, loss, accuracy)
}
