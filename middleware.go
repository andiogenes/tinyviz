package main

import (
	"fmt"
	"graph-labs/tinyviz/graphics"
	"graph-labs/tinyviz/input"
	"graph-labs/tinyviz/legacy"
	"log"
	"sync"
)

// visualize визуализирует граф на основе информации из дескриптора
func visualize(fileName string, format graphics.ImageFormat, quality int) error {
	options, err := input.LoadGraphData(fileName)
	if err != nil {
		options, err = legacy.LoadGraphData(fileName)
	}

	if err != nil {
		return err
	}

	err = graphics.RenderGraph(fmt.Sprintf("%s.viz.%s", fileName, format.Stringify()), &options, graphics.PutVertexInRandomFreeCell, format, quality)
	if err != nil {
		return err
	}
	fmt.Printf("%s visualizated\n", fileName)

	return nil
}

// visualizeFolder ищет все дескрипторы графов в папке и визуализирует все графы по ним
func visualizeFolder(format graphics.ImageFormat, quality int) {
	descriptors, err := getDescriptors()
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	wg.Add(len(descriptors))

	for _, val := range descriptors {
		go func(fileName string) {
			if err := visualize(fileName, format, quality); err != nil {
				log.Printf("Unable to process %s: %s\n", fileName, err.Error())
			}
			wg.Done()
		}(val)
	}

	wg.Wait()
}