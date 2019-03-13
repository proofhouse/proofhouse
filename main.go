package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/cucumber/gherkin-go"
	"github.com/pkg/errors"
	"io/ioutil"
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

	sql.Register()

	config, err := NewConfig("./configuration.yaml")
	if err != nil {
		panic(err)
	}

	fmt.Println(config.Plugins)

	var features []Feature

	err = filepath.Walk(config.FeaturesDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return errors.Wrapf(err, "Failed to read feature file '%v'", path)
		}
		if info.IsDir() || filepath.Ext(path) != ".feature" {
			return nil
		}

		features = append(features, NewFeature(path, info))

		return nil
	})
	if err != nil {
		panic(err)
	}

	for _, f := range features {
		data, err := ioutil.ReadFile(f.path)
		if err != nil {
			panic(errors.Wrapf(err, "Failed to read feature file '%v'", f.path))
		}

		doc, err := gherkin.ParseGherkinDocument(bytes.NewReader(data))
		if err != nil {
			panic(errors.Wrap(err, "Failed to parse feature"))
		}

		_ = doc
	}
}
