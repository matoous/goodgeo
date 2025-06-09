package transform

import (
	"github.com/matoous/goodgeo"
)

// UniqueCoords creates a new coordinate array (with the same layout as the inputs) that
// contains each unique coordinate in the coordData.  The ordering of the coords are the
// same as the input.
func UniqueCoords(layout goodgeo.Layout, compare Compare, coordData []float64) []float64 {
	set := NewTreeSet(layout, compare)
	stride := layout.Stride()
	uniqueCoords := make([]float64, 0, len(coordData))
	numCoordsAdded := 0
	for i := 0; i < len(coordData); i += stride {
		coord := coordData[i : i+stride]
		added := set.Insert(goodgeo.Coord(coord))

		if added {
			uniqueCoords = append(uniqueCoords, coord...)
			numCoordsAdded++
		}
	}
	return uniqueCoords[:numCoordsAdded*stride]
}
