package main

import (
	"fmt"
	"github.com/paw1a/digit-recognition/pkg/algebra"
	"github.com/paw1a/digit-recognition/pkg/model"
)

func main() {
	mod := model.NewModel([]int{784, 128, 64, 10}, 0.001)

	input := algebra.RandomVector(784)
	output := mod.FeedForward(input)

	mod.BackPropagation(output, []float64{0, 0, 1, 0, 0, 0, 0, 0, 0, 0})

	fmt.Printf("%v", output)
}
