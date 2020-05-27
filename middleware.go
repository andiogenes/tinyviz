package main

import (
	"fmt"
	"github.com/andiogenes/tinyviz/graphics"
	"github.com/andiogenes/tinyviz/input"
	"github.com/andiogenes/tinyviz/legacy"
	"log"
	"sync"
)

// visualize визуализирует граф на основе информации из дескриптора
func visualize(fileName string, format graphics.ImageFormat, quality int, arrangeFn graphics.ArrangementFn, dataLoaderFn input.ArrangementLoader) error {
	options, err := input.LoadGraphData(fileName)
	if err != nil {
		options, err = legacy.LoadGraphData(fileName)
	}

	if err != nil {
		return err
	}

	// Tries to load additional data from file with name fileName by dataLoaderFn
	var data interface{}
	if dataLoaderFn != nil {
		data, err = dataLoaderFn(fileName)
		if err != nil {
			return err
		}
	} else {
		data = nil
	}

	err = graphics.RenderGraph(fmt.Sprintf("%s.viz.%s", fileName, format.Stringify()), &options, arrangeFn, data, format, quality)
	if err != nil {
		return err
	}
	fmt.Printf("%s visualized\n", fileName)

	return nil
}

// visualizeFolder ищет все дескрипторы графов в папке и визуализирует все графы по ним
func visualizeFolder(format graphics.ImageFormat, quality int, arrangeFn graphics.ArrangementFn, dataLoaderFn input.ArrangementLoader) {
	descriptors, err := getDescriptors()
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	wg.Add(len(descriptors))

	for _, val := range descriptors {
		go func(fileName string) {
			if err := visualize(fileName, format, quality, arrangeFn, dataLoaderFn); err != nil {
				log.Printf("Unable to process %s: %s\n", fileName, err.Error())
			}
			wg.Done()
		}(val)
	}

	wg.Wait()
}
