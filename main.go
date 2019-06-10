package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/urfave/cli"
)

// visualize визуализирует граф на основе информации из дескриптора
func visualize(fileName string) error {
	count, isDirected, isWeighted, isColored, names, path, matrix, weights, colors, colorMatrix, err := loadGraphData(fileName)
	if err != nil {
		return err
	}

	shuffleSeed()

	renderGraph(fmt.Sprintf("%s.viz.png", fileName), count, isDirected, isWeighted, isColored, names, path, matrix, weights, colors, colorMatrix)
	fmt.Printf("%s visualizated\n", fileName)

	return nil
}

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

// visualizeFolder ищет все дескрипторы графов в папке и визуализирует все графы по ним
func visualizeFolder() {
	descriptors, err := getDescriptors()
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	wg.Add(len(descriptors))

	for _, val := range descriptors {
		go func(fileName string) {
			if err := visualize(fileName); err != nil {
				log.Printf("Unable to process %s: %s\n", fileName, err.Error())
			}
			wg.Done()
		}(val)
	}

	wg.Wait()
}

func main() {
	var format string
	var quality int

	app := cli.NewApp()

	app.Name = "tinyviz"
	app.Version = "0.2.0"
	app.Description = "Graph visualization tool for educational purposes"
	app.Authors = []cli.Author{
		{Name: "Anton", Email: "megadeathlightsaber@gmail.com"},
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "format, f",
			Value:       "png",
			Usage:       "image output format (jpg/png)",
			Destination: &format,
		},
		cli.IntFlag{
			Name:        "quality, q",
			Value:       80,
			Usage:       "jpeg image ouput quality (from 0 to 100)",
			Destination: &quality,
		},
	}

	app.Action = func(c *cli.Context) error {
		if format != "png" && format != "jpg" && format != "jpeg" {
			fmt.Println("Unknown format \"", format, "\", reset to \"png\"")
			format = "png"
		}

		if quality <= 0 || quality > 100 {
			fmt.Println("Unbounded quality value ", quality, ", reset to 80")
			quality = 80
		}

		if c.NArg() > 0 {
			err := visualize(c.Args()[0])
			return err
		}

		visualizeFolder()

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
