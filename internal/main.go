package main

import (
	"fmt"
	"github.com/paw1a/digit-recognition/pkg/algebra"
	"github.com/paw1a/digit-recognition/pkg/model"
)

func main() {
	mod := model.NewModel([]int{784, 128, 64, 10})

	input := algebra.RandomVector(784)
	output := mod.FeedForward(input)

	fmt.Printf("%v", output)
}
