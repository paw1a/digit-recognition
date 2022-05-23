package dataset

import (
	"encoding/binary"
	"fmt"
	"os"
)

type ImagesHeader struct {
	MagicNumber int32
	ImagesCount int32
	Rows        int32
	Cols        int32
}

type LabelsHeader struct {
	MagicNumber int32
	LabelsCount int32
}

func LoadDataset(imagesPath string, labelsPath string) ([][]float64, [][]float64, error) {
	imagesFile, err := os.Open(imagesPath)
	if err != nil {
		return nil, nil, fmt.Errorf("can't open file %s: %v", imagesPath, err)
	}

	labelsFile, err := os.Open(labelsPath)
	if err != nil {
		return nil, nil, fmt.Errorf("can't open file %s: %v", labelsPath, err)
	}

	defer imagesFile.Close()
	defer labelsFile.Close()

	imagesHeader, err := readImagesHeader(imagesFile)
	if err != nil {
		return nil, nil, err
	}

	labelsHeader, err := readLabelsHeader(labelsFile)
	if err != nil {
		return nil, nil, err
	}

	images, err := loadImages(imagesFile, imagesHeader)
	if err != nil {
		return nil, nil, err
	}

	labels, err := loadLabels(labelsFile, labelsHeader)
	if err != nil {
		return nil, nil, err
	}

	return images, labels, nil
}

func loadImages(file *os.File, imagesHeader *ImagesHeader) ([][]float64, error) {
	imageBytes := make([][]byte, imagesHeader.ImagesCount)
	size := int(imagesHeader.Rows * imagesHeader.Cols)

	for i := 0; i < len(imageBytes); i++ {
		imageBytes[i] = make([]byte, size)

		err := binary.Read(file, binary.BigEndian, imageBytes[i])

		if err != nil {
			return nil, fmt.Errorf("can't read image data: %v", err)
		}
	}

	images := make([][]float64, imagesHeader.ImagesCount)

	for i := 0; i < len(imageBytes); i++ {
		images[i] = make([]float64, size)
		for j := 0; j < size; j++ {
			images[i][j] = float64(imageBytes[i][j]) / 255
		}
	}

	return images, nil
}

func loadLabels(file *os.File, labelsHeader *LabelsHeader) ([][]float64, error) {
	labelBytes := make([]byte, labelsHeader.LabelsCount)

	err := binary.Read(file, binary.BigEndian, labelBytes)

	if err != nil {
		return nil, fmt.Errorf("can't read label data: %v", err)
	}

	labels := make([][]float64, len(labelBytes))

	for i := 0; i < len(labels); i++ {
		labels[i] = make([]float64, 10)
		labels[i][labelBytes[i]] = 1
	}

	return labels, nil
}

func readImagesHeader(file *os.File) (*ImagesHeader, error) {
	imagesHeader := new(ImagesHeader)

	err := binary.Read(file, binary.BigEndian, imagesHeader)
	if err != nil {
		return nil, fmt.Errorf("can't read images header from file %s: %v",
			file.Name(), err)
	}

	return imagesHeader, nil
}

func readLabelsHeader(file *os.File) (*LabelsHeader, error) {
	labelsHeader := new(LabelsHeader)

	err := binary.Read(file, binary.BigEndian, labelsHeader)
	if err != nil {
		return nil, fmt.Errorf("can't read labels header from file %s: %v",
			file.Name(), err)
	}

	return labelsHeader, nil
}
