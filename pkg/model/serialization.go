package model

import (
	"encoding/gob"
	"fmt"
	"os"
)

func SerializeModel(filePath string, mod *Model) error {
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

func DeserializeModel(filePath string, mod *Model) error {
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
