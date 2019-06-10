package main

import (
	"fmt"
	"graph-labs/tinyviz/graphics"
	"graph-labs/tinyviz/legacy"
	"graph-labs/tinyviz/random"
	"log"
	"sync"
)

// visualize визуализирует граф на основе информации из дескриптора
func visualize(fileName string) error {
	count, isDirected, isWeighted, isColored, names, path, matrix, weights, colors, colorMatrix, err := legacy.LoadGraphData(fileName)
	if err != nil {
		return err
	}

	random.ShuffleSeed()

	graphics.RenderGraph(fmt.Sprintf("%s.viz.png", fileName), count, isDirected, isWeighted, isColored, names, path, matrix, weights, colors, colorMatrix)
	fmt.Printf("%s visualizated\n", fileName)

	return nil
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
