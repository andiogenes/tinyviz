package graphics

import (
	"graph-labs/tinyviz/random"
	"image/color"
)

// RenderGraph рисует по заданным данным граф и сохраняет изображение в png-файл output
func RenderGraph(output string, vertexCount int, isDirected bool, isWeighted bool, isColored bool, names []string, path []int, matrix [][]int, weights [][]string, colors []uint32, colorCover [][]int) {
	positions := make([]vertex2D, vertexCount)
	combination := random.Combination(vertexCount*vertexCount, vertexCount)

	// imgSide := vertexCount * CellSide
	imgSide := (vertexCount + 1) * CellSide

	for i := 0; i < vertexCount; i++ {
		positions[i].x = float64(combination[i]%vertexCount+1) * CellSide
		positions[i].y = float64(combination[i]/vertexCount+1) * CellSide
		positions[i].inPath = false
	}

	for _, val := range path {
		positions[val].inPath = true
	}

	context := generateContext(imgSide, imgSide)

	for i := 0; i < vertexCount; i++ {
		drawVertex(context, names[i], positions[i].x, positions[i].y, VertexRadius, positions[i].inPath, isColored, colors, colorCover[i][i])
		for j := 0; j < vertexCount; j++ {
			if matrix[i][j] == 1 {
				drawEdge(context, positions[i].x, positions[i].y, positions[j].x, positions[j].y, VertexRadius, isDirected, isColored, colors, colorCover[i][j])
				if isWeighted {
					drawEdgeWeight(context, weights[i][j], positions[i].x, positions[i].y, positions[j].x, positions[j].y)
				}
			}
		}
	}

	if len(path) > 0 {
		firstPathID := path[0]
		lastPathID := path[len(path)-1]

		context.SetRGB255(0, 0, 255)
		context.DrawCircle(positions[firstPathID].x, positions[firstPathID].y, VertexRadius)
		context.Stroke()

		context.SetRGB255(255, 0, 0)
		context.DrawCircle(positions[lastPathID].x, positions[lastPathID].y, VertexRadius)
		context.Stroke()
	}

	context.SetColor(color.Black)
	context.DrawString("Original Graph", 10, 15)
	outW, outH := context.MeasureString(output)
	context.Push()
	context.SetRGBA255(0, 0, 0, 125)
	context.DrawString(output, float64(imgSide)-outW-8, float64(imgSide)-outH+8)
	context.Pop()

	context.SavePNG(output)
}
