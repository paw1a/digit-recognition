package app

import (
	"fmt"
	"github.com/paw1a/digit-recognition/internal/dataset"
	"github.com/paw1a/digit-recognition/internal/desktop"
	"github.com/paw1a/digit-recognition/pkg/model"
	"os"
	"path"
)

const (
	ModelDirectory = "model"
	ModelFilename  = "model.gob"
)

func Run() {
	mod := new(model.Model)

	if _, err := os.Stat(path.Join(ModelDirectory, ModelFilename)); err != nil {
		mod = model.NewModel([]int{784, 250, 10}, 0.001, 5)
		go TrainModel(mod)
	} else {
		err := model.DeserializeModel(path.Join(ModelDirectory, ModelFilename), mod)
		if err != nil {
			fmt.Printf("failed to load model, new model was created: %v\n", err)
			mod = model.NewModel([]int{784, 250, 10}, 0.001, 5)
			go TrainModel(mod)
		}
	}

	app := desktop.NewApplication(mod)
	app.Run()
}

func TrainModel(mod *model.Model) {
	if !dataset.Exists() {
		fmt.Printf("Dataset dowloading...\n")
		err := dataset.DownloadDataset()
		if err != nil {
			fmt.Printf("failed to download dataset: %v\n", err)
			return
		}
	}

	err := os.MkdirAll(ModelDirectory, os.ModePerm)
	if err != nil {
		fmt.Printf("can't create model directory: %v", err)
		return
	}

	fmt.Printf("Dataset loading...\n")
	input, labels, err := dataset.LoadDataset(
		path.Join(dataset.Directory, dataset.TrainImages),
		path.Join(dataset.Directory, dataset.TrainLabels))

	if err != nil {
		fmt.Printf("failed to load dataset: %v\n", err)
		return
	}

	mod.Fit(input, labels)

	err = model.SerializeModel(path.Join(ModelDirectory, ModelFilename), mod)
	if err != nil {
		fmt.Printf("train error: %v\n", err)
	}
}
