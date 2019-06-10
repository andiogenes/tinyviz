package graphics

import (
	"fmt"
	"graph-labs/tinyviz/random"
	"image/color"
)

// RenderGraph рисует по заданным данным граф и сохраняет изображение в png-файл output
func RenderGraph(output string, options *RenderOptions) error {
	if options == nil {
		return fmt.Errorf("Nil argument passed")
	}

	positions := make([]vertex2D, options.VertexCount)
	combination := random.Combination(options.VertexCount*options.VertexCount, options.VertexCount)

	// imgSide := vertexCount * CellSide
	imgSide := (options.VertexCount + 1) * CellSide

	for i := 0; i < options.VertexCount; i++ {
		positions[i].x = float64(combination[i]%options.VertexCount+1) * CellSide
		positions[i].y = float64(combination[i]/options.VertexCount+1) * CellSide
		positions[i].inPath = false
	}

	for _, val := range options.Path {
		positions[val].inPath = true
	}

	context := generateContext(imgSide, imgSide)

	for i := 0; i < options.VertexCount; i++ {
		drawVertex(context, options.Names[i], positions[i].x, positions[i].y, VertexRadius, positions[i].inPath, options.Colored, options.Colors, options.ColorCover[i][i])
		for j := 0; j < options.VertexCount; j++ {
			if options.Matrix[i][j] == 1 {
				drawEdge(context, positions[i].x, positions[i].y, positions[j].x, positions[j].y, VertexRadius, options.Directed, options.Colored, options.Colors, options.ColorCover[i][j])
				if options.Weighted {
					drawEdgeWeight(context, options.Weights[i][j], positions[i].x, positions[i].y, positions[j].x, positions[j].y)
				}
			}
		}
	}

	if len(options.Path) > 0 {
		firstPathID := options.Path[0]
		lastPathID := options.Path[len(options.Path)-1]

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

	return nil
}
