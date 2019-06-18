package graphics

// arrange.go contains set of built-in arrangement rules for RenderGraph function of render.go

import (
	"fmt"
	"graph-labs/tinyviz/random"
)

// PutVertexInRandomFreeCell ...
func PutVertexInRandomFreeCell(positions []Vertex2D, options RenderOptions, additionalData interface{}) error {
	combination := random.Combination(options.VertexCount*options.VertexCount, options.VertexCount)

	for i := 0; i < options.VertexCount; i++ {
		positions[i].x = float64(combination[i]%options.VertexCount+1) * CellSide
		positions[i].y = float64(combination[i]/options.VertexCount+1) * CellSide
		positions[i].inPath = false
	}

	return nil
}

// PutVertexAtPosition ...
func PutVertexAtPosition(positions []Vertex2D, options RenderOptions, additionalData interface{}) error {
	data, correct := additionalData.([][]int)
	if !correct {
		return fmt.Errorf("Type assertion failed - given data doesn't represent type [][]int")
	}

	for i := 0; i < options.VertexCount; i++ {
		positions[i].x = float64(data[i][0])
		positions[i].y = float64(data[i][1])
		positions[i].inPath = false
	}

	return nil
}
