package main

import (
	"fmt"
	"github.com/paw1a/digit-recognition/internal/dataset"
	"github.com/paw1a/digit-recognition/pkg/model"
	"path"
)

const (
	TrainLabels = "train-labels-idx1-ubyte"
	TrainImages = "train-images-idx3-ubyte"
	TestLabels  = "t10k-labels-idx1-ubyte"
	TestImages  = "t10k-images-idx3-ubyte"
)

func main() {
	mod := model.NewModel([]int{784, 800, 10}, 0.001, 5)

	trainImages, trainLabels, err := dataset.LoadDataset(
		path.Join("mnist", TrainImages), path.Join("mnist", TrainLabels))
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	testImages, testLabels, err := dataset.LoadDataset(
		path.Join("mnist", TestImages), path.Join("mnist", TestLabels))
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	mod.Fit(trainImages, trainLabels)

	mod.TestModel(testImages, testLabels)
}
