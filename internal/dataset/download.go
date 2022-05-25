package dataset

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

const (
	Directory = "mnist"

	TrainLabels = "train-labels-idx1-ubyte"
	TrainImages = "train-images-idx3-ubyte"
	TestLabels  = "t10k-labels-idx1-ubyte"
	TestImages  = "t10k-images-idx3-ubyte"
)

var (
	filenames = []string{
		TrainImages,
		TrainLabels,
		TestImages,
		TestLabels,
	}

	urls = []string{
		"http://yann.lecun.com/exdb/mnist/train-labels-idx1-ubyte.gz",
		"http://yann.lecun.com/exdb/mnist/train-images-idx3-ubyte.gz",
		"http://yann.lecun.com/exdb/mnist/t10k-labels-idx1-ubyte.gz",
		"http://yann.lecun.com/exdb/mnist/t10k-images-idx3-ubyte.gz",
	}
)

func DatasetExists() bool {
	if _, err := os.Stat(Directory); err != nil {
		return false
	}

	for _, filename := range filenames {
		if _, err := os.Stat(filename); err != nil {
			return false
		}
	}

	return true
}

func DownloadFile(url string, savePath string) error {
	file, err := os.Create(savePath)
	if err != nil {
		return fmt.Errorf("can't create file %s: %v", savePath, err)
	}
	defer file.Close()

	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("can't get file from url %s: %v", url, err)
	}
	defer response.Body.Close()

	gzipReader, err := gzip.NewReader(response.Body)
	if err != nil {
		return fmt.Errorf("can't uncompress dataset file %s: %v", savePath, err)
	}
	defer gzipReader.Close()

	_, err = io.Copy(file, gzipReader)
	if err != nil {
		return fmt.Errorf("can't write dataset to file %s: %v", savePath, err)
	}

	return nil
}

func DownloadDataset() error {
	err := os.MkdirAll(Directory, os.ModePerm)
	if err != nil {
		return fmt.Errorf("can't create dataset directory: %v", err)
	}

	for i, filename := range filenames {
		if _, err := os.Stat(filename); err != nil {
			err = DownloadFile(urls[i], path.Join(Directory, filenames[i]))
			if err != nil {
				return err
			}
		}
	}

	return nil
}
