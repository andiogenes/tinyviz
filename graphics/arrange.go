package graphics

// arrange.go contains set of built-in arrangement rules for RenderGraph function of render.go

import "graph-labs/tinyviz/random"

// PutVertexInRandomFreeCell ...
func PutVertexInRandomFreeCell(positions []Vertex2D, options RenderOptions, additionalData interface{}) {
	combination := random.Combination(options.VertexCount*options.VertexCount, options.VertexCount)

	for i := 0; i < options.VertexCount; i++ {
		positions[i].x = float64(combination[i]%options.VertexCount+1) * CellSide
		positions[i].y = float64(combination[i]/options.VertexCount+1) * CellSide
		positions[i].inPath = false
	}
}
