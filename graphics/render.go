package graphics

import (
	"fmt"
	"image/color"

	"github.com/fogleman/gg"
)

// RenderGraph рисует по заданным данным граф и сохраняет изображение в png-файл output
func RenderGraph(output string, options *RenderOptions, arrangeFn func([]vertex2D, RenderOptions), format ImageFormat, quality int) error {
	if options == nil {
		return fmt.Errorf("Nil argument passed")
	}

	// Initialize helper variables
	positions := make([]vertex2D, options.VertexCount)
	imgSide := (options.VertexCount + 1) * CellSide
	context := generateContext(imgSide, imgSide)

	// Set position of vertices by some rule
	arrangeFn(positions, *options)

	// Marks path vertices
	for _, val := range options.Path {
		positions[val].inPath = true
	}

	// Vertex and edges rendering
	for i := 0; i < options.VertexCount; i++ {
		r, g, b, a := pickColor(options.Colored, options.Colors, options.ColorCover[i][i], positions[i].inPath)
		drawVertex(context, options.Names[i], positions[i].x, positions[i].y, VertexRadius, r, g, b, a)

		for j := 0; j < options.VertexCount; j++ {
			if options.Matrix[i][j] == 1 {
				r, g, b, a := pickColor(options.Colored, options.Colors, options.ColorCover[i][j], false)
				drawEdge(context, positions[i].x, positions[i].y, positions[j].x, positions[j].y, VertexRadius, options.Directed, r, g, b, a)

				if options.Weighted {
					drawEdgeWeight(context, options.Weights[i][j], positions[i].x, positions[i].y, positions[j].x, positions[j].y)
				}
			}
		}
	}

	// Highlight first and last elements of path
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

	// Write information about rendered graph
	context.SetColor(color.Black)
	context.DrawString("Original Graph", 10, 15)
	outW, outH := context.MeasureString(output)
	context.Push()
	context.SetRGBA255(0, 0, 0, 125)
	context.DrawString(output, float64(imgSide)-outW-8, float64(imgSide)-outH+8)
	context.Pop()

	if format == Png {
		context.SavePNG(output)
	} else {
		gg.SaveJPG(output, context.Image(), quality)
	}

	return nil
}
