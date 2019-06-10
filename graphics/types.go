package graphics

// imageFormat represents format of rendering image
type imageFormat byte

const (
	png imageFormat = iota
	jpeg
)

// RenderOptions ...
type RenderOptions struct {
	// Count of vertices in graph
	VertexCount int
	// Is graph directed or not
	Directed bool
	// Is graph weighted or not
	Weighted bool
	// Is graph colored or not
	Colored bool
	// Names of graph's vertices
	Names []string
	// Colors used in graph
	Colors []uint32
	// ...
	Path []int
	// Adjacency matrix of graph
	Matrix [][]int
	// Weight matrix of graph
	Weights [][]string
	// Color matrix of graph
	ColorCover [][]int
}

// vertex2D represents ...
type vertex2D struct {
	x      float64
	y      float64
	inPath bool
}
