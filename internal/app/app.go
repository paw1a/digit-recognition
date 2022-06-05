package app

import (
	"fmt"
	"github.com/paw1a/digit-recognition/internal/dataset"
	"github.com/paw1a/digit-recognition/pkg/model"
	"path"
)

func Run() {
	if dataset.Exists() {
		err := dataset.DownloadDataset()
		if err != nil {
			fmt.Printf("download dataset error: %v", err)
		}
	}

	mod := model.NewModel([]int{784, 800, 10}, 0.001, 5)

	trainImages, trainLabels, err := dataset.LoadDataset(
		path.Join(dataset.Directory, dataset.TrainImages),
		path.Join(dataset.Directory, dataset.TrainLabels))
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	testImages, testLabels, err := dataset.LoadDataset(
		path.Join(dataset.Directory, dataset.TestImages),
		path.Join(dataset.Directory, dataset.TestLabels))
	if err != nil {
		fmt.Printf("%v", err)
		return
	}

	mod.Fit(trainImages, trainLabels)

	mod.TestModel(testImages, testLabels)
}
