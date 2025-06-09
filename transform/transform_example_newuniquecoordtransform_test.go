package transform_test

import (
	"fmt"

	"github.com/matoous/goodgeo"
	"github.com/matoous/goodgeo/sorting"
	"github.com/matoous/goodgeo/transform"
)

type coordTransformExampleCompare struct{}

func (c coordTransformExampleCompare) IsEquals(x, y goodgeo.Coord) bool {
	return x[0] == y[0] && x[1] == y[1]
}

func (c coordTransformExampleCompare) IsLess(x, y goodgeo.Coord) bool {
	return sorting.IsLess2D(x, y)
}

func ExampleUniqueCoords() {
	coordData := []float64{0, 0, 1, 1, 1, 1, 3, 3, 0, 0}
	layout := goodgeo.XY

	filteredCoords := transform.UniqueCoords(layout, coordTransformExampleCompare{}, coordData)
	fmt.Println(filteredCoords)
	// Output: [0 0 1 1 3 3]
}
