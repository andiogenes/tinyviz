package main

import (
	"fmt"
	"log"
	"sync"
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
