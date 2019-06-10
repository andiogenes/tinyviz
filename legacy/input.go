package legacy

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

	names := strings.Split(strings.Replace(string(f), "\r", "", -1), "\n")
	if names[len(names)-1] == "" {
		names = names[:len(names)-1]
	}

	return names, nil
}

// loadMatrix загружает матрицу смежности/весов графа из файла
func loadMatrix(fileName string, matrixSize int) ([][]int, error) {
	f, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	str := strings.Replace(string(f), " ", "", -1)
	str = strings.Replace(str, "\r\n", "", -1)
	str = strings.Replace(str, "\n", "", -1)
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
				err = errors.New("Mismatch between graph descriptor and matrix")
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

// loadStringMatrix загружает матрицу со строчными значениями из файла
func loadStringMatrix(fileName string, matrixSize int) ([][]string, error) {
	f, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	str := strings.Replace(string(f), " ", "", -1)
	str = strings.Replace(str, "\r\n", "", -1)
	str = strings.Replace(str, "\n", "", -1)
	values := strings.Split(str, ",")
	if values[len(values)-1] == "" {
		values = values[:len(values)-1]
	}

	// Длина среза с элементами
	valuesCount := len(values)

	// matrixSize := int(math.Sqrt(float64(len(values))))

	matrix := make([][]string, matrixSize)

	// Заполнение матрицы
	for i := 0; i < matrixSize; i++ {
		matrix[i] = make([]string, matrixSize)
		for j := 0; j < matrixSize; j++ {
			// Проверка на принадлежность границам массива
			if valuesCount <= i*matrixSize+j {
				err = errors.New("Mismatch between graph descriptor and matrix")
				return nil, err
			}
			// Записывает в ячейку i,j соответствующее значение
			matrix[i][j] = values[i*matrixSize+j]
		}
	}

	return matrix, nil
}

// stringifyMatrix конвертирует числовую матрицу в строчную
func stringifyMatrix(matrix [][]int) [][]string {
	out := make([][]string, len(matrix))

	lenMatrix := len(matrix)

	for i := 0; i < lenMatrix; i++ {
		out[i] = make([]string, lenMatrix)
		for j := 0; j < lenMatrix; j++ {
			out[i][j] = strconv.Itoa(matrix[i][j])
		}
	}

	return out
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

// loadColors загружает список цветов вершин/ребер в формате ARGB и матрицу цветов графа
func loadColors(colorsFileName string, matrixFileName string, vertexCount int) ([]uint32, [][]int, error) {
	// Загрузка списка цветов
	f, err := ioutil.ReadFile(colorsFileName)
	if err != nil {
		return nil, nil, err
	}

	str := strings.Split(string(f), " ")
	if str[len(str)-1] == "" {
		str = str[:len(str)-1]
	}

	colors := make([]uint32, 0)

	for _, v := range str {
		color, err := strconv.ParseUint(v, 16, 32)
		if err != nil {
			return nil, nil, err
		}

		colors = append(colors, uint32(color))
	}

	// Загрузка матрицы
	matrix, err := loadMatrix(matrixFileName, vertexCount)
	if err != nil {
		return nil, nil, err
	}

	// Проверка на совпадение индексов цветов в матрице
	for i := 0; i < vertexCount; i++ {
		for j := 0; j < vertexCount; j++ {
			if matrix[i][j] < 0 || (matrix[i][j] > len(colors) && len(colors) != 0) {
				return nil, nil, fmt.Errorf("Color id %d in color matrix at [%d, %d] is out of range", matrix[i][j], i, j)
			}
		}
	}

	return colors, matrix, nil
}

// LoadGraphData загружает размер, флаг ориентированности, флаг взвешенности, флаг окрашенности,
// имена вершин, список пути, матрицу смежности графа, матрицу весов, список цветов, матрицу цветов
func LoadGraphData(fileName string) (int, bool, bool, bool, []string, []int, [][]int, [][]string, []uint32, [][]int, error) {
	// Загрузка файла-дескриптора
	descr, err := ioutil.ReadFile(fileName)
	if err != nil {
		return 0, false, false, false, nil, nil, nil, nil, nil, nil, err
	}

	// Парсинг дескриптора
	str := strings.Split(strings.Replace(string(descr), "\r", "", -1), "\n")
	// Если в выражении меньше двух значений, то оно некорректно
	if len(str) < 2 {
		return 0, false, false, false, nil, nil, nil, nil, nil, nil, err
	}

	// Извлечение количества вершин из файла
	vertexCount, err := strconv.Atoi(str[0])
	if err != nil {
		return 0, false, false, false, nil, nil, nil, nil, nil, nil, err
	}

	// Извлечение флага ориентированности
	directionFlag, err := (strconv.Atoi(str[1]))
	if err != nil {
		return 0, false, false, false, nil, nil, nil, nil, nil, nil, err
	}
	concreteDirectionFlag := (directionFlag == 1)

	// Извлечение флага взвешенности, если он имеется
	weightFlag := false
	if len(str) >= 3 {
		value, err := strconv.Atoi(str[2])
		if err != nil {
			return 0, false, false, false, nil, nil, nil, nil, nil, nil, err
		}
		weightFlag = (value == 1)
	}

	// Извлечение флага окрашенности, если он имеется
	colorFlag := false
	if len(str) >= 4 {
		value, err := strconv.Atoi(str[3])
		if err != nil {
			return 0, false, false, false, nil, nil, nil, nil, nil, nil, err
		}
		colorFlag = (value == 1)
	}

	// Загрузка имен вершин
	names, err := loadNames(strings.Join([]string{fileName, ".names"}, ""))
	if err != nil {
		return 0, false, false, false, nil, nil, nil, nil, nil, nil, err
	}

	if namesCount := len(names); namesCount != vertexCount {
		err = fmt.Errorf("Mismatch between number of vertices and it's names: %d vs %d", vertexCount, namesCount)
		return 0, false, false, false, nil, nil, nil, nil, nil, nil, err
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
		return 0, false, false, false, nil, nil, nil, nil, nil, nil, err
	}

	// Загрузка матрицы весов
	weights, err := loadStringMatrix(strings.Join([]string{fileName, ".weights"}, ""), vertexCount)
	if err != nil {
		fmt.Println(err.Error())
		weights = stringifyMatrix(matrix)
	}

	// Загрузка матрицы цветов
	colors, colorsMatrix, err := loadColors(strings.Join([]string{fileName, ".colors"}, ""),
		strings.Join([]string{fileName, ".cmatrix"}, ""),
		vertexCount)

	if err != nil {
		fmt.Println(err.Error())
		colors = make([]uint32, 0)
		colorsMatrix = matrix
	}

	return vertexCount, concreteDirectionFlag, weightFlag, colorFlag, names, path, matrix, weights, colors, colorsMatrix, nil
}
