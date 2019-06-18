package input

import (
	"encoding/json"
	"io/ioutil"
)

// ArrangementLoader ...
type ArrangementLoader func(string) (interface{}, error)

type coordRepresentation struct {
	Coords [][]int `json:"coordinates"`
}

type concreteCoords struct{ coords [][]int }

// CoordinatesLoader ...
func CoordinatesLoader(fileName string) (interface{}, error) {
	f, err := ioutil.ReadFile(fileName)
	if err != nil {
		return concreteCoords{}, err
	}

	// Unmarshall data
	var cRep coordRepresentation
	if err = json.Unmarshal(f, &cRep); err != nil {
		return concreteCoords{}, err
	}

	return cRep.Coords, nil
}
