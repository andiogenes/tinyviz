package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// loadNames загружает список имен вершин из файла
func loadNames(fileName string) ([]string, error) {
	f, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	names := strings.Split(string(f), "\r\n")
	if names[len(names)-1] == "" {
		names = names[:len(names)-1]
	}

	return names, nil
}

// loadMatrix загружает матрицу смежности графа из файла
func loadMatrix(fileName string, matrixSize int) ([][]int, error) {
	f, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	str := strings.Replace(string(f), " ", "", -1)
	str = strings.Replace(str, "\r\n", "", -1)
	values := strings.Split(str, ",")
	if values[len(values)-1] == "" {
		values = values[:len(values)-1]
	}

	// Длина среза с элементами
	valuesCount := len(values)

	// matrixSize := int(math.Sqrt(float64(len(values))))

	matrix := make([][]int, matrixSize)

	// Заполнение матрицы
	for i := 0; i < matrixSize; i++ {
		matrix[i] = make([]int, matrixSize)
		for j := 0; j < matrixSize; j++ {
			// Проверка на принадлежность границам массива
			if valuesCount <= i*matrixSize+j {
				err = errors.New("Mismatch between graph descriptor and adjacency matrix")
				return nil, err
			}
			strValue := values[i*matrixSize+j]
			matrix[i][j], err = strconv.Atoi(strValue)
			if err != nil {
				return nil, err
			}
		}
	}

	return matrix, nil
}

// loadPath загружает список вершин в маршруте из файла
func loadPath(fileName string, vertexCount int) ([]int, error) {
	f, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	strPath := strings.Split(string(f), " ")

	var path []int

	if len(strPath) == 1 && strPath[0] == "" {
		return path, nil
	}

	for _, val := range strPath {
		id, err := strconv.Atoi(val)
		if err != nil {
			return nil, err
		}

		if id < 0 || id >= vertexCount {
			err = fmt.Errorf("Vertex id in traverse path is out of range")
			return nil, err
		}

		path = append(path, id)
	}

	return path, nil
}

// Загружает размер, флаг ориентированности, имена вершин, список пути и матрицу смежности графа
func loadGraphData(fileName string) (int, bool, []string, []int, [][]int, error) {
	// Загрузка файла-дескриптора
	descr, err := ioutil.ReadFile(fileName)
	if err != nil {
		return 0, false, nil, nil, nil, err
	}

	// Парсинг дескриптора
	str := strings.Split(string(descr), "\r\n")
	if len(str) < 2 {
		return 0, false, nil, nil, nil, err
	}

	// Извлечение количества вершин из файла
	vertexCount, err := strconv.Atoi(str[0])
	if err != nil {
		return 0, false, nil, nil, nil, err
	}

	// Извлечение флага ориентированности
	directionFlag, err := (strconv.Atoi(str[1]))
	if err != nil {
		return 0, false, nil, nil, nil, err
	}
	concreteDirectionFlag := (directionFlag == 1)

	// Загрузка имен вершин
	names, err := loadNames(strings.Join([]string{fileName, ".names"}, ""))
	if err != nil {
		return 0, false, nil, nil, nil, err
	}

	if namesCount := len(names); namesCount != vertexCount {
		err = fmt.Errorf("Mismatch between number of vertices and it's names: %d vs %d", vertexCount, namesCount)
		return 0, false, nil, nil, nil, err
	}

	path, err := loadPath(strings.Join([]string{fileName, ".path"}, ""), vertexCount)
	if err != nil {
		//return 0, false, nil, nil, nil, err
		fmt.Println(err.Error())
		path = make([]int, 0)
	}

	// Загрузка матрицы смежности
	matrix, err := loadMatrix(strings.Join([]string{fileName, ".matrix"}, ""), vertexCount)
	if err != nil {
		return 0, false, nil, nil, nil, err
	}

	return vertexCount, concreteDirectionFlag, names, path, matrix, nil
}
