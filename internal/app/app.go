package app

import (
	"github.com/paw1a/digit-recognition/internal/dataset"
	"github.com/paw1a/digit-recognition/internal/desktop"
	"github.com/paw1a/digit-recognition/pkg/model"
	"log"
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
			log.Printf("failed to load model, new model was created: %v\n", err)
			mod = model.NewModel([]int{784, 250, 10}, 0.001, 5)
			go TrainModel(mod)
		}
	}

	app := desktop.NewApplication(mod)
	app.Run()
}

func TrainModel(mod *model.Model) {
	if !dataset.Exists() {
		log.Printf("Dataset dowloading...\n")
		err := dataset.DownloadDataset()
		if err != nil {
			log.Printf("failed to download dataset: %v\n", err)
			return
		}
	}

	err := os.MkdirAll(ModelDirectory, os.ModePerm)
	if err != nil {
		log.Printf("can't create model directory: %v", err)
		return
	}

	log.Printf("Dataset loading...\n")
	input, labels, err := dataset.LoadDataset(
		path.Join(dataset.Directory, dataset.TrainImages),
		path.Join(dataset.Directory, dataset.TrainLabels))

	if err != nil {
		log.Printf("failed to load dataset: %v\n", err)
		return
	}

	log.Printf("Model training...")
	mod.Fit(input, labels)

	log.Printf("Model serialization...")
	err = model.SerializeModel(path.Join(ModelDirectory, ModelFilename), mod)
	if err != nil {
		log.Printf("train error: %v\n", err)
	}
}
