package graphics

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

// convertColor разделяет цвет из RGBA255 по соответствующим компонентам (Red, Green, Blue, alpha)
func convertColor(colorRgba uint32) (int, int, int, int) {
	return int((colorRgba >> 24) & 0x000000ff),
		int((colorRgba >> 16) & 0x000000ff),
		int((colorRgba >> 8) & 0x000000ff),
		int(colorRgba & 0x000000ff)
}

// drawVertex рисует вершину графа
func drawVertex(context *gg.Context, name string, x, y, r float64, inPath bool, isColored bool, colors []uint32, colorIndex int) {
	context.DrawCircle(x, y, r)
	context.SetColor(color.Black)
	context.SetLineWidth(2)
	context.StrokePreserve()
	if !inPath {
		if isColored && colorIndex != 0 {
			context.SetRGBA255(convertColor(colors[colorIndex-1]))
		} else {
			context.SetRGBA255(VertexColorRed, VertexColorGreen, VertexColorBlue, VertexColorAlpha)
		}
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
func drawEdge(context *gg.Context, x1, y1, x2, y2, r float64, isDirected bool, isColored bool, colors []uint32, colorIndex int) {
	vecX, vecY := float64(x2-x1), float64(y2-y1)
	vecLen := math.Sqrt(vecX*vecX + vecY*vecY)
	// Преобразование длины вектора в r
	vecX, vecY = (vecX/vecLen)*r, (vecY/vecLen)*r

	if isColored && colorIndex != 0 {
		context.SetRGBA255(convertColor(colors[colorIndex-1]))
	} else {
		context.SetColor(color.Black)
	}
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
		if isColored && colorIndex != 0 {
			context.SetRGBA255(convertColor(colors[colorIndex-1]))
		} else {
			context.SetColor(color.Black)
		}
		context.Fill()
	}
}

// drawWeightInfo отображает информацию о весе ребра
func drawEdgeWeight(context *gg.Context, weight string, x1, y1, x2, y2 float64) {
	medianX, medianY := (x2-x1)/2., (y2-y1)/2.

	strW, strH := context.MeasureString(weight)

	context.SetRGBA255(255, 135, 245, 175)
	context.DrawRectangle(x1+medianX, y1+medianY-strH, strW, strH)
	context.Fill()
	context.SetRGB255(145, 35, 185)
	context.DrawString(weight, x1+medianX, y1+medianY)
}

type vertex2D struct {
	x      float64
	y      float64
	inPath bool
}
