package input

import (
	"encoding/json"
	"github.com/andiogenes/tinyviz/graphics"
	"io/ioutil"
	"strconv"
)

type representation struct {
	Count    int      `json:"count"`
	Directed bool     `json:"directed"`
	Weighted bool     `json:"weighted"`
	Colored  bool     `json:"colored"`
	Names    []string `json:"names"`
	Colors   []string `json:"colors"`
	Path     []int    `json:"path"`
	Matrix   [][]int  `json:"matrix"`
	Weights  [][]int  `json:"weights"`
	Cover    [][]int  `json:"cover"`
}

// LoadGraphData parses JSON-file and returns information about graph or error
func LoadGraphData(fileName string) (graphics.RenderOptions, error) {
	f, err := ioutil.ReadFile(fileName)
	if err != nil {
		return graphics.RenderOptions{}, err
	}

	// Unmarshal data try
	var rep representation
	if err = json.Unmarshal(f, &rep); err != nil {
		return graphics.RenderOptions{}, err
	}

	// Convert string slice to uint32 slice try
	colors := make([]uint32, 0)
	for _, v := range rep.Colors {
		color, err := strconv.ParseUint(v, 16, 32)
		if err != nil {
			return graphics.RenderOptions{}, err
		}

		colors = append(colors, uint32(color))
	}

	// Conversion of slice of slices of int to slice of slices of strings
	weights := make([][]string, len(rep.Weights))
	for i, v1 := range rep.Weights {
		line := make([]string, len(rep.Weights[i]))

		for j, v2 := range v1 {
			line[j] = strconv.Itoa(v2)
		}

		weights[i] = line
	}

	// Initializing return value
	retVal := graphics.RenderOptions{
		VertexCount: rep.Count,
		Directed:    rep.Directed,
		Weighted:    rep.Weighted,
		Colored:     rep.Colored,
		Names:       rep.Names,
		Colors:      colors,
		Path:        rep.Path,
		Matrix:      rep.Matrix,
		Weights:     weights,
		ColorCover:  rep.Cover,
	}

	return retVal, nil
}
