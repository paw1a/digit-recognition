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
	mod := model.NewModel([]int{784, 128, 64, 10}, 0.001, 5)

	images, labels, err := dataset.LoadDataset(
		path.Join("mnist", TrainImages), path.Join("mnist", TrainLabels))
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	mod.Fit(images, labels)
}
