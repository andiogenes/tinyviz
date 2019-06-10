package main

import (
	"fmt"
	"graph-labs/tinyviz/graphics"
	"os"
	"path/filepath"
)

// getDescriptors возвращает срез имен всех файлов-дескрипторов в текущей папке
func getDescriptors() ([]string, error) {
	curPath, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	var files []string

	err = filepath.Walk(curPath, func(path string, f os.FileInfo, errout error) error {
		if errout != nil {
			return errout
		}

		if !f.IsDir() {
			if filepath.Ext(path) == ".descr" {
				files = append(files, f.Name())
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

// pickFormat ...
func pickFormat(extension string) (graphics.ImageFormat, error) {
	switch extension {
	case "png":
		return graphics.Png, nil
	case "jpg", "jpeg":
		return graphics.Jpeg, nil
	default:
		return graphics.Png, fmt.Errorf("Incorrect argument: excepted \"png\", \"jpg\" or \"jpeg\", given \"%s\"", extension)
	}
}
