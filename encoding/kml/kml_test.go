package kml

import (
	"encoding/xml"
	"strings"
	"testing"

	"github.com/alecthomas/assert/v2"

	"github.com/matoous/goodgeo"
)

func Test(t *testing.T) {
	for _, tc := range []struct {
		g    goodgeo.T
		want string
	}{
		{
			g:    goodgeo.NewPoint(goodgeo.XY),
			want: `<Point><coordinates>0,0</coordinates></Point>`,
		},
		{
			g:    goodgeo.NewPoint(goodgeo.XY).MustSetCoords(goodgeo.Coord{0, 0}),
			want: `<Point><coordinates>0,0</coordinates></Point>`,
		},
		{
			g:    goodgeo.NewPoint(goodgeo.XYZ).MustSetCoords(goodgeo.Coord{0, 0, 0}),
			want: `<Point><coordinates>0,0</coordinates></Point>`,
		},
		{
			g:    goodgeo.NewPoint(goodgeo.XYZ).MustSetCoords(goodgeo.Coord{0, 0, 1}),
			want: `<Point><coordinates>0,0,1</coordinates></Point>`,
		},
		{
			g:    goodgeo.NewPoint(goodgeo.XYM).MustSetCoords(goodgeo.Coord{0, 0, 1}),
			want: `<Point><coordinates>0,0</coordinates></Point>`,
		},
		{
			g:    goodgeo.NewPoint(goodgeo.XYZM).MustSetCoords(goodgeo.Coord{0, 0, 0, 1}),
			want: `<Point><coordinates>0,0</coordinates></Point>`,
		},
		{
			g:    goodgeo.NewPoint(goodgeo.XYZM).MustSetCoords(goodgeo.Coord{0, 0, 1, 1}),
			want: `<Point><coordinates>0,0,1</coordinates></Point>`,
		},
		{
			g: goodgeo.NewMultiPoint(goodgeo.XY).MustSetCoords([]goodgeo.Coord{{1, 2}, {3, 4}, {5, 6}}),
			want: `<MultiGeometry>` +
				`<Point>` +
				`<coordinates>1,2</coordinates>` +
				`</Point>` +
				`<Point>` +
				`<coordinates>3,4</coordinates>` +
				`</Point>` +
				`<Point>` +
				`<coordinates>5,6</coordinates>` +
				`</Point>` +
				`</MultiGeometry>`,
		},
		{
			g: goodgeo.NewLineString(goodgeo.XY).MustSetCoords([]goodgeo.Coord{
				{0, 0},
				{1, 1},
			}),
			want: `<LineString><coordinates>0,0 1,1</coordinates></LineString>`,
		},
		{
			g: goodgeo.NewLineString(goodgeo.XYZ).MustSetCoords([]goodgeo.Coord{
				{1, 2, 3},
				{4, 5, 6},
			}),
			want: `<LineString><coordinates>1,2,3 4,5,6</coordinates></LineString>`,
		},
		{
			g: goodgeo.NewLineString(goodgeo.XYM).MustSetCoords([]goodgeo.Coord{
				{1, 2, 3},
				{4, 5, 6},
			}),
			want: `<LineString><coordinates>1,2 4,5</coordinates></LineString>`,
		},
		{
			g: goodgeo.NewLineString(goodgeo.XYZM).MustSetCoords([]goodgeo.Coord{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
			}),
			want: `<LineString><coordinates>1,2,3 5,6,7</coordinates></LineString>`,
		},
		{
			g: goodgeo.NewMultiLineString(goodgeo.XY).MustSetCoords([][]goodgeo.Coord{
				{{1, 2}, {3, 4}, {5, 6}, {7, 8}},
			}),
			want: `<MultiGeometry>` +
				`<LineString>` +
				`<coordinates>1,2 3,4 5,6 7,8</coordinates>` +
				`</LineString>` +
				`</MultiGeometry>`,
		},
		{
			g: goodgeo.NewMultiLineString(goodgeo.XY).MustSetCoords([][]goodgeo.Coord{
				{{1, 2}, {3, 4}, {5, 6}, {7, 8}},
				{{9, 10}, {11, 12}, {13, 14}},
			}),
			want: `<MultiGeometry>` +
				`<LineString>` +
				`<coordinates>1,2 3,4 5,6 7,8</coordinates>` +
				`</LineString>` +
				`<LineString>` +
				`<coordinates>9,10 11,12 13,14</coordinates>` +
				`</LineString>` +
				`</MultiGeometry>`,
		},
		{
			g: goodgeo.NewPolygon(goodgeo.XY).MustSetCoords([][]goodgeo.Coord{
				{{1, 2}, {3, 4}, {5, 6}, {1, 2}},
			}),
			want: `<Polygon>` +
				`<outerBoundaryIs>` +
				`<LinearRing>` +
				`<coordinates>1,2 3,4 5,6 1,2</coordinates>` +
				`</LinearRing>` +
				`</outerBoundaryIs>` +
				`</Polygon>`,
		},
		{
			g: goodgeo.NewPolygon(goodgeo.XYZ).MustSetCoords([][]goodgeo.Coord{
				{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {1, 2, 3}},
			}),
			want: `<Polygon>` +
				`<outerBoundaryIs>` +
				`<LinearRing>` +
				`<coordinates>1,2,3 4,5,6 7,8,9 1,2,3</coordinates>` +
				`</LinearRing>` +
				`</outerBoundaryIs>` +
				`</Polygon>`,
		},
		{
			g: goodgeo.NewPolygon(goodgeo.XYZ).MustSetCoords([][]goodgeo.Coord{
				{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {1, 2, 3}},
				{{0.4, 0.5, 0.6}, {0.7, 0.8, 0.9}, {0.1, 0.2, 0.3}, {0.4, 0.5, 0.6}},
			}),
			want: `<Polygon>` +
				`<outerBoundaryIs>` +
				`<LinearRing>` +
				`<coordinates>1,2,3 4,5,6 7,8,9 1,2,3</coordinates>` +
				`</LinearRing>` +
				`</outerBoundaryIs>` +
				`<innerBoundaryIs>` +
				`<LinearRing>` +
				`<coordinates>0.4,0.5,0.6 0.7,0.8,0.9 0.1,0.2,0.3 0.4,0.5,0.6</coordinates>` +
				`</LinearRing>` +
				`</innerBoundaryIs>` +
				`</Polygon>`,
		},
		{
			g: goodgeo.NewMultiPolygon(goodgeo.XYZ).MustSetCoords([][][]goodgeo.Coord{
				{
					{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {1, 2, 3}},
					{{0.4, 0.5, 0.6}, {0.7, 0.8, 0.9}, {0.1, 0.2, 0.3}, {0.4, 0.5, 0.6}},
				},
			}),
			want: `<MultiGeometry>` +
				`<Polygon>` +
				`<outerBoundaryIs>` +
				`<LinearRing>` +
				`<coordinates>1,2,3 4,5,6 7,8,9 1,2,3</coordinates>` +
				`</LinearRing>` +
				`</outerBoundaryIs>` +
				`<innerBoundaryIs>` +
				`<LinearRing>` +
				`<coordinates>0.4,0.5,0.6 0.7,0.8,0.9 0.1,0.2,0.3 0.4,0.5,0.6</coordinates>` +
				`</LinearRing>` +
				`</innerBoundaryIs>` +
				`</Polygon>` +
				`</MultiGeometry>`,
		},
		{
			g: goodgeo.NewMultiPolygon(goodgeo.XYZ).MustSetCoords([][][]goodgeo.Coord{
				{
					{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {1, 2, 3}},
				},
				{
					{{0.4, 0.5, 0.6}, {0.7, 0.8, 0.9}, {0.1, 0.2, 0.3}, {0.4, 0.5, 0.6}},
				},
			}),
			want: `<MultiGeometry>` +
				`<Polygon>` +
				`<outerBoundaryIs>` +
				`<LinearRing>` +
				`<coordinates>1,2,3 4,5,6 7,8,9 1,2,3</coordinates>` +
				`</LinearRing>` +
				`</outerBoundaryIs>` +
				`</Polygon>` +
				`<Polygon>` +
				`<outerBoundaryIs>` +
				`<LinearRing>` +
				`<coordinates>0.4,0.5,0.6 0.7,0.8,0.9 0.1,0.2,0.3 0.4,0.5,0.6</coordinates>` +
				`</LinearRing>` +
				`</outerBoundaryIs>` +
				`</Polygon>` +
				`</MultiGeometry>`,
		},
		{
			g: goodgeo.NewGeometryCollection().MustPush(
				goodgeo.NewLineString(goodgeo.XY).MustSetCoords([]goodgeo.Coord{
					{-122.4425587930444, 37.80666418607323},
					{-122.4428379594768, 37.80663578323093},
				}),
				goodgeo.NewLineString(goodgeo.XY).MustSetCoords([]goodgeo.Coord{
					{-122.4425509770566, 37.80662588061205},
					{-122.4428340530617, 37.8065999493009},
				}),
			),
			want: `<MultiGeometry>` +
				`<LineString>` +
				`<coordinates>` +
				`-122.4425587930444,37.80666418607323 -122.4428379594768,37.80663578323093` +
				`</coordinates>` +
				`</LineString>` +
				`<LineString>` +
				`<coordinates>` +
				`-122.4425509770566,37.80662588061205 -122.4428340530617,37.8065999493009` +
				`</coordinates>` +
				`</LineString>` +
				`</MultiGeometry>`,
		},
	} {
		t.Run(tc.want, func(t *testing.T) {
			sb := &strings.Builder{}
			e := xml.NewEncoder(sb)
			element, err := Encode(tc.g)
			assert.NoError(t, err)
			assert.NoError(t, e.Encode(element))
			assert.Equal(t, tc.want, sb.String())
		})
	}
}
