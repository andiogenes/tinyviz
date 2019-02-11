package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
)

// visualize визуализирует граф на основе информации из дескриптора
func visualize(fileName string) error {
	count, isDirected, isWeighted, names, path, matrix, weights, err := loadGraphData(fileName)
	if err != nil {
		return err
	}

	shuffleSeed()

	renderGraph(fmt.Sprintf("%s.viz.png", fileName), count, isDirected, isWeighted, names, path, matrix, weights)
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
	// Если конкретный дескриптор задан в аргументах командной строки,
	// то запускается визуализация только по нему
	// Иначе: визуализируются все графы по имеющимся в папке дескрипторам
	if len(os.Args) >= 2 {
		if err := visualize(os.Args[1]); err != nil {
			log.Fatal(err)
		}
	} else {
		visualizeFolder()
	}
}
