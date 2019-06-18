package graphics

// ImageFormat represents format of rendering image
type ImageFormat byte

const (
	// Png represents ...
	Png ImageFormat = iota
	// Jpeg represents ...
	Jpeg
)

// Stringify ...
func (f ImageFormat) Stringify() string {
	switch f {
	case Png:
		return "png"
	case Jpeg:
		return "jpeg"
	}

	return ""
}

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

// Vertex2D represents ...
type Vertex2D struct {
	x      float64
	y      float64
	inPath bool
}

// graphElement ...
type graphElement byte

const (
	geVertex graphElement = iota
	geEdge
)

// ArrangementFn ...
type ArrangementFn func([]Vertex2D, RenderOptions, interface{})
