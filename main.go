package main

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
)

type Feature struct {
	path     string
	fileInfo os.FileInfo
}

func NewFeature(path string, fileInfo os.FileInfo) Feature {
	return Feature{
		path:     path,
		fileInfo: fileInfo,
	}
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("%+v", r)
			os.Exit(-1)
		}
	}()

	config, err := NewConfig("./configuration.yaml")
	if err != nil {
		panic(err)
	}

	features := make([]Feature, 1)

	err = filepath.Walk(config.FeaturesDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return errors.Wrapf(err, `Failed to read feature file '%v'`, path)
		}

		features = append(features, NewFeature(path, info))

		return nil
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf(`%+v`, features)
}
