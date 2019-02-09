package main

import (
	"image/color"
	"math"

	"github.com/fogleman/gg"
)

const (
	// CellSide - коэффициент размера визуализации
	CellSide = 48
	// VertexRadius - радиус вершины
	VertexRadius = 16
	// VertexColorRed - компонента красного цвета вершины
	VertexColorRed = 15
	// VertexColorGreen - компонента зеленого цвета вершины
	VertexColorGreen = 164
	// VertexColorBlue - компонента синего цвета вершины
	VertexColorBlue = 125
	// VertexColorAlpha - значение альфа-канала цвета вершины
	VertexColorAlpha = 125
)

// generateContext создает контекст для рисования заданного размера
func generateContext(width, height int) *gg.Context {
	context := gg.NewContext(width, height)
	context.SetColor(color.White)
	context.Clear()

	return context
}

// drawVertex рисует вершину графа
func drawVertex(context *gg.Context, name string, x, y, r float64, inPath bool) {
	context.DrawCircle(x, y, r)
	context.SetColor(color.Black)
	context.SetLineWidth(2)
	context.StrokePreserve()
	if !inPath {
		context.SetRGBA255(VertexColorRed, VertexColorGreen, VertexColorBlue, VertexColorAlpha)
	} else {
		context.SetRGBA255(255-VertexColorRed, 255-VertexColorGreen, 255-VertexColorBlue, VertexColorAlpha)
	}
	context.Fill()
	context.SetColor(color.Black)

	strW, strH := context.MeasureString(name)

	context.DrawString(name, x+r/2.-strW*1.5, y+r/2.-strH/2.)
	// context.DrawString(name, x-strW/2., y-r/2.+strH/2.)
}

// drawEdge рисует ребро графа
func drawEdge(context *gg.Context, x1, y1, x2, y2, r float64, isDirected bool) {
	vecX, vecY := float64(x2-x1), float64(y2-y1)
	vecLen := math.Sqrt(vecX*vecX + vecY*vecY)
	// Преобразование длины вектора в r
	vecX, vecY = (vecX/vecLen)*r, (vecY/vecLen)*r

	context.SetColor(color.Black)
	context.SetLineWidth(1.2)
	context.DrawLine(x1+vecX, y1+vecY, x2-vecX, y2-vecY)
	context.Stroke()

	if isDirected {
		var normX, normY float64

		if vecX == 0 {
			normX = -1
		} else {
			normX = 1 / vecX
		}

		if vecY == 0 {
			normY = -1
		} else {
			normY = -1 / vecY
		}

		normLen := math.Sqrt(normX*normX + normY*normY)
		normX, normY = (normX/normLen)*r*0.2, (normY/normLen)*r*0.2

		context.LineTo(x2-vecX, y2-vecY)
		context.LineTo(x2-vecX*1.5+normX, y2-vecY*1.5+normY)
		context.LineTo(x2-vecX*1.5-normX, y2-vecY*1.5-normY)
		context.LineTo(x2-vecX, y2-vecY)
		context.SetColor(color.Black)
		context.Fill()
	}
}

// drawWeightInfo отображает информацию о весе ребра
func drawEdgeWeight(context *gg.Context, weight int, x1, y1, x2, y2 float64) {
	medianX, medianY := (x2-x1)/2., (y2-y1)/2.

	strW, strH := context.MeasureString(fmt.Sprintf("%d", weight))

	context.SetRGBA255(255, 135, 245, 175)
	context.DrawRectangle(x1+medianX, y1+medianY-strH, strW, strH)
	context.Fill()
	context.SetRGB255(145, 35, 185)
	context.DrawString(fmt.Sprintf("%d", weight), x1+medianX, y1+medianY)
}

type vertex2D struct {
	x      float64
	y      float64
	inPath bool
}

// renderGraph рисует по заданным данным граф и сохраняет изображение в png-файл output
func renderGraph(output string, vertexCount int, isDirected bool, names []string, path []int, matrix [][]int) {
	positions := make([]vertex2D, vertexCount)
	combination := randomCombination(vertexCount*vertexCount, vertexCount)

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
		drawVertex(context, names[i], positions[i].x, positions[i].y, VertexRadius, positions[i].inPath)
		for j := 0; j < vertexCount; j++ {
			if matrix[i][j] != 0 {
				drawEdge(context, positions[i].x, positions[i].y, positions[j].x, positions[j].y, VertexRadius, isDirected)
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
