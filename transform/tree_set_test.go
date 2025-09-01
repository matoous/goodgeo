package transform

import (
	"testing"

	"github.com/alecthomas/assert/v2"

	"github.com/matoous/goodgeo"
	"github.com/matoous/goodgeo/sorting"
)

func TestTree(t *testing.T) {
	set := NewTreeSet(goodgeo.XY, testCompare{})
	set.Insert([]float64{3, 1})
	set.Insert([]float64{3, 2})
	set.Insert([]float64{1, 2})
	set.Insert([]float64{4, 1})
	set.Insert([]float64{1, 1})
	set.Insert([]float64{6, 6})
	set.Insert([]float64{1, 1})
	set.Insert([]float64{3, 1})

	expected := []float64{
		1, 1, 1, 2,
		3, 1, 3, 2,
		4, 1, 6, 6,
	}

	actual := set.ToFlatArray()
	assert.Equal(t, expected, actual)
}

type testCompare struct{}

func (c testCompare) IsEquals(x, y goodgeo.Coord) bool {
	return x[0] == y[0] && x[1] == y[1]
}

func (c testCompare) IsLess(x, y goodgeo.Coord) bool {
	return sorting.IsLess2D(x, y)
}
