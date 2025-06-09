package wkt

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert/v2"

	"github.com/matoous/goodgeo"
)

func TestMarshalAndUnmarshal(t *testing.T) {
	for _, tc := range []struct {
		g goodgeo.T
		s string
	}{
		{
			g: goodgeo.NewPointEmpty(goodgeo.XY),
			s: "POINT EMPTY",
		},
		{
			g: goodgeo.NewPoint(goodgeo.XY).MustSetCoords(goodgeo.Coord{1.337, 2.42}),
			s: "POINT (1.337 2.42)",
		},
		{
			g: goodgeo.NewPoint(goodgeo.XYZ).MustSetCoords(goodgeo.Coord{1, 2, 3}),
			s: "POINT Z (1 2 3)",
		},
		{
			g: goodgeo.NewPoint(goodgeo.XYM).MustSetCoords(goodgeo.Coord{1, 2, 3}),
			s: "POINT M (1 2 3)",
		},
		{
			g: goodgeo.NewPoint(goodgeo.XYZM).MustSetCoords(goodgeo.Coord{1, 2, 3, 4}),
			s: "POINT ZM (1 2 3 4)",
		},
		{
			g: goodgeo.NewLineString(goodgeo.XY),
			s: "LINESTRING EMPTY",
		},
		{
			g: goodgeo.NewLineString(goodgeo.XY).MustSetCoords([]goodgeo.Coord{{1, 2}, {3, 4}}),
			s: "LINESTRING (1 2, 3 4)",
		},
		{
			g: goodgeo.NewLineString(goodgeo.XY).MustSetCoords([]goodgeo.Coord{{0, 0}, {10, 0}, {10, 10}, {0, 0}}),
			s: "LINESTRING (0 0, 10 0, 10 10, 0 0)",
		},
		{
			g: goodgeo.NewLineString(goodgeo.XYZ).MustSetCoords([]goodgeo.Coord{{1, 2, 3}, {4, 5, 6}}),
			s: "LINESTRING Z (1 2 3, 4 5 6)",
		},
		{
			g: goodgeo.NewLineString(goodgeo.XYM).MustSetCoords([]goodgeo.Coord{{1, 2, 3}, {4, 5, 6}}),
			s: "LINESTRING M (1 2 3, 4 5 6)",
		},
		{
			g: goodgeo.NewLineString(goodgeo.XYZM).MustSetCoords([]goodgeo.Coord{{1, 2, 3, 4}, {5, 6, 7, 8}}),
			s: "LINESTRING ZM (1 2 3 4, 5 6 7 8)",
		},
		{
			g: goodgeo.NewPolygon(goodgeo.XY),
			s: "POLYGON EMPTY",
		},
		{
			g: goodgeo.NewPolygon(goodgeo.XY).MustSetCoords([][]goodgeo.Coord{{{1, 2}, {3, 4}, {5, 6}, {1, 2}}}),
			s: "POLYGON ((1 2, 3 4, 5 6, 1 2))",
		},
		{
			g: goodgeo.NewPolygon(goodgeo.XY).MustSetCoords([][]goodgeo.Coord{{{1, 2}, {3, 4}, {5, 6}, {1, 2}}, {{7, 8}, {9, 10}, {11, 12}, {7, 8}}}),
			s: "POLYGON ((1 2, 3 4, 5 6, 1 2), (7 8, 9 10, 11 12, 7 8))",
		},
		{
			g: goodgeo.NewPolygon(goodgeo.XYM).MustSetCoords([][]goodgeo.Coord{{{0, 0, 0}, {1, 0, 1}, {1, 1, 2}, {0, 0, 3}}}),
			s: "POLYGON M ((0 0 0, 1 0 1, 1 1 2, 0 0 3))",
		},
		{
			g: goodgeo.NewPolygon(goodgeo.XYZM).MustSetCoords([][]goodgeo.Coord{{{0, 0, 0, 0}, {1, 0, -1, 1}, {1, 1, -2, 2}, {0, 0, 0, 3}}}),
			s: "POLYGON ZM ((0 0 0 0, 1 0 -1 1, 1 1 -2 2, 0 0 0 3))",
		},
		{
			g: goodgeo.NewMultiPoint(goodgeo.XY),
			s: "MULTIPOINT EMPTY",
		},
		{
			g: goodgeo.NewMultiPoint(goodgeo.XY).MustSetCoords([]goodgeo.Coord{nil, nil}),
			s: "MULTIPOINT (EMPTY, EMPTY)",
		},
		{
			g: goodgeo.NewMultiPoint(goodgeo.XY).MustSetCoords([]goodgeo.Coord{{1, 2}}),
			s: "MULTIPOINT (1 2)",
		},
		{
			g: goodgeo.NewMultiPoint(goodgeo.XY).MustSetCoords([]goodgeo.Coord{{1, 2}, nil, {3, 4}}),
			s: "MULTIPOINT (1 2, EMPTY, 3 4)",
		},
		{
			g: goodgeo.NewMultiPoint(goodgeo.XYZM).MustSetCoords([]goodgeo.Coord{{1, 2, 1, 42}, nil, {3, 4, 1, 43}}),
			s: "MULTIPOINT ZM (1 2 1 42, EMPTY, 3 4 1 43)",
		},
		{
			g: goodgeo.NewMultiLineString(goodgeo.XY),
			s: "MULTILINESTRING EMPTY",
		},
		{
			g: goodgeo.NewMultiLineString(goodgeo.XY).MustSetCoords([][]goodgeo.Coord{nil, nil}),
			s: "MULTILINESTRING (EMPTY, EMPTY)",
		},
		{
			g: goodgeo.NewMultiLineString(goodgeo.XY).MustSetCoords([][]goodgeo.Coord{{{1, 2}, {3, 4}}}),
			s: "MULTILINESTRING ((1 2, 3 4))",
		},
		{
			g: goodgeo.NewMultiLineString(goodgeo.XY).MustSetCoords([][]goodgeo.Coord{{{1, 2}, {3, 4}}, nil, {{5, 6}, {7, 8}}}),
			s: "MULTILINESTRING ((1 2, 3 4), EMPTY, (5 6, 7 8))",
		},
		{
			g: goodgeo.NewMultiPolygon(goodgeo.XY),
			s: "MULTIPOLYGON EMPTY",
		},
		{
			g: goodgeo.NewMultiPolygon(goodgeo.XY).MustSetCoords([][][]goodgeo.Coord{nil, nil}),
			s: "MULTIPOLYGON (EMPTY, EMPTY)",
		},
		{
			g: goodgeo.NewMultiPolygon(goodgeo.XY).MustSetCoords([][][]goodgeo.Coord{{{{1, 2}, {3, 4}, {5, 6}, {1, 2}}}}),
			s: "MULTIPOLYGON (((1 2, 3 4, 5 6, 1 2)))",
		},
		{
			g: goodgeo.NewMultiPolygon(goodgeo.XY).MustSetCoords([][][]goodgeo.Coord{{{{1, 2}, {3, 4}, {5, 6}, {1, 2}}}, nil, {{{7, 8}, {9, 10}, {11, 12}, {7, 8}}}}),
			s: "MULTIPOLYGON (((1 2, 3 4, 5 6, 1 2)), EMPTY, ((7 8, 9 10, 11 12, 7 8)))",
		},
		{
			g: goodgeo.NewMultiPolygon(goodgeo.XYZM).MustSetCoords([][][]goodgeo.Coord{
				{
					{{-1, -1, 10, 42}, {1000, -1, 10, 42}, {1000, 1000, 10, 42}, {-1, -1, 10, 42}},
				},
				{
					{{0, 0, 10, 42}, {100, 0, 10, 42}, {100, 100, 10, 42}, {0, 0, 10, 42}},
					{{10, 10, 10, 42}, {90, 10, 10, 42}, {90, 90, 10, 42}, {10, 10, 10, 42}},
				},
			}),
			s: "MULTIPOLYGON ZM (((-1 -1 10 42, 1000 -1 10 42, 1000 1000 10 42, -1 -1 10 42)), ((0 0 10 42, 100 0 10 42, 100 100 10 42, 0 0 10 42), (10 10 10 42, 90 10 10 42, 90 90 10 42, 10 10 10 42)))",
		},
		{
			g: goodgeo.NewGeometryCollection().MustSetLayout(goodgeo.XY),
			s: "GEOMETRYCOLLECTION EMPTY",
		},
		{
			g: goodgeo.NewGeometryCollection().MustPush(goodgeo.NewGeometryCollection().MustSetLayout(goodgeo.XY)).MustSetLayout(goodgeo.XY),
			s: "GEOMETRYCOLLECTION (GEOMETRYCOLLECTION EMPTY)",
		},
		{
			g: goodgeo.NewGeometryCollection().MustPush(
				goodgeo.NewPointEmpty(goodgeo.XY),
				goodgeo.NewLineString(goodgeo.XY),
				goodgeo.NewPolygon(goodgeo.XY),
			).MustSetLayout(goodgeo.XY),
			s: "GEOMETRYCOLLECTION (POINT EMPTY, LINESTRING EMPTY, POLYGON EMPTY)",
		},
		{
			g: goodgeo.NewGeometryCollection().MustPush(
				goodgeo.NewPoint(goodgeo.XY).MustSetCoords(goodgeo.Coord{1, 2}),
				goodgeo.NewLineString(goodgeo.XY).MustSetCoords([]goodgeo.Coord{{3, 4}, {5, 6}}),
			).MustSetLayout(goodgeo.XY),
			s: "GEOMETRYCOLLECTION (POINT (1 2), LINESTRING (3 4, 5 6))",
		},
	} {
		t.Run(tc.s, func(t *testing.T) {
			t.Run("marshal", func(t *testing.T) {
				got, err := Marshal(tc.g)
				assert.NoError(t, err)
				assert.Equal(t, tc.s, got)
			})

			t.Run("unmarshal", func(t *testing.T) {
				got, err := Unmarshal(tc.s)
				assert.NoError(t, err)
				assert.Equal(t, tc.g, got)
			})
		})
	}
}

func TestEncoder(t *testing.T) {
	for _, tc := range []struct {
		encoder *Encoder
		g       goodgeo.T
		s       string
	}{
		{
			encoder: NewEncoder(EncodeOptionWithMaxDecimalDigits(0)),
			g:       goodgeo.NewPointFlat(goodgeo.XY, []float64{1.001, 1.066}),
			s:       "POINT (1 1)",
		},
		{
			encoder: NewEncoder(EncodeOptionWithMaxDecimalDigits(0)),
			g:       goodgeo.NewPointFlat(goodgeo.XY, []float64{10.001, 100.066}),
			s:       "POINT (10 100)",
		},
		{
			encoder: NewEncoder(EncodeOptionWithMaxDecimalDigits(1)),
			g:       goodgeo.NewPointFlat(goodgeo.XY, []float64{10.001, 1.066}),
			s:       "POINT (10 1.1)",
		},
		{
			encoder: NewEncoder(EncodeOptionWithMaxDecimalDigits(1)),
			g:       goodgeo.NewPointFlat(goodgeo.XY, []float64{1.001, 1.066}),
			s:       "POINT (1 1.1)",
		},
		{
			encoder: NewEncoder(EncodeOptionWithMaxDecimalDigits(2)),
			g:       goodgeo.NewPointFlat(goodgeo.XY, []float64{1.001, 1.066}),
			s:       "POINT (1 1.07)",
		},
		{
			encoder: NewEncoder(EncodeOptionWithMaxDecimalDigits(3)),
			g:       goodgeo.NewPointFlat(goodgeo.XY, []float64{1.001, 1.066}),
			s:       "POINT (1.001 1.066)",
		},
		{
			encoder: NewEncoder(EncodeOptionWithMaxDecimalDigits(4)),
			g:       goodgeo.NewPointFlat(goodgeo.XY, []float64{1.001, 1.066}),
			s:       "POINT (1.001 1.066)",
		},
	} {
		t.Run(fmt.Sprintf("%s(encoder=%#v)", tc.s, tc.encoder), func(t *testing.T) {
			got, err := tc.encoder.Encode(tc.g)
			assert.NoError(t, err)
			assert.Equal(t, tc.s, got)
		})
	}
}

func TestUnmarshalEmptyGeomWithArbitrarySpaces(t *testing.T) {
	for _, tc := range []struct {
		g goodgeo.T
		s string
	}{
		{
			g: goodgeo.NewPointEmpty(goodgeo.XY),
			s: "POINT      EMPTY",
		},
		{
			g: goodgeo.NewLineString(goodgeo.XY),
			s: "LINESTRING     EMPTY",
		},
		{
			g: goodgeo.NewPolygon(goodgeo.XY),
			s: "POLYGON      EMPTY",
		},
		{
			g: goodgeo.NewMultiPoint(goodgeo.XY),
			s: "MULTIPOINT      EMPTY",
		},
		{
			g: goodgeo.NewMultiLineString(goodgeo.XY),
			s: "MULTILINESTRING   EMPTY",
		},
		{
			g: goodgeo.NewMultiPolygon(goodgeo.XY),
			s: "MULTIPOLYGON                EMPTY",
		},
		{
			g: goodgeo.NewGeometryCollection().MustSetLayout(goodgeo.XY),
			s: "GEOMETRYCOLLECTION      EMPTY",
		},
	} {
		t.Run(tc.s, func(t *testing.T) {
			got, err := Unmarshal(tc.s)
			assert.NoError(t, err)
			assert.Equal(t, tc.g, got)
		})
	}
}

func TestUnmarshal(t *testing.T) {
	testCases := []struct {
		desc        string
		equivInputs []string
		expected    goodgeo.T
	}{
		// POINT tests
		{
			desc:        "parse 2D point",
			equivInputs: []string{"POINT(0 1)", "POINT (0 1)", "point(0 1)", "point ( 0 1 )"},
			expected:    goodgeo.NewPointFlat(goodgeo.XY, []float64{0, 1}),
		},
		{
			desc:        "parse 2D point with scientific notation",
			equivInputs: []string{"POINT(1e-2 2e3)", "POINT(0.1e-1 2e3)", "POINT(0.01e-0 2e+3)", "POINT(0.01 2000)"},
			expected:    goodgeo.NewPointFlat(goodgeo.XY, []float64{1e-2, 2e3}),
		},
		{
			desc:        "parse 2D+M point",
			equivInputs: []string{"POINT M (-2 0 0.5)", "POINTM(-2 0 0.5)", "POINTM(-2 0 .5)"},
			expected:    goodgeo.NewPointFlat(goodgeo.XYM, []float64{-2, 0, 0.5}),
		},
		{
			desc:        "parse 3D point",
			equivInputs: []string{"POINT Z (2 3 4)", "POINTZ(2 3 4)", "POINT(2 3 4)"},
			expected:    goodgeo.NewPointFlat(goodgeo.XYZ, []float64{2, 3, 4}),
		},
		{
			desc:        "parse 4D point",
			equivInputs: []string{"POINT ZM (0 5 -10 15)", "POINTZM (0 5 -10 15)", "POINT(0 5 -10 15)"},
			expected:    goodgeo.NewPointFlat(goodgeo.XYZM, []float64{0, 5, -10, 15}),
		},
		{
			desc:        "parse empty 2D point",
			equivInputs: []string{"POINT EMPTY"},
			expected:    goodgeo.NewPointEmpty(goodgeo.XY),
		},
		{
			desc:        "parse empty 2D+M point",
			equivInputs: []string{"POINT M EMPTY", "POINTM EMPTY"},
			expected:    goodgeo.NewPointEmpty(goodgeo.XYM),
		},
		{
			desc:        "parse empty 3D point",
			equivInputs: []string{"POINT Z EMPTY", "POINTZ EMPTY"},
			expected:    goodgeo.NewPointEmpty(goodgeo.XYZ),
		},
		{
			desc:        "parse empty 4D point",
			equivInputs: []string{"POINT ZM EMPTY", "POINTZM EMPTY"},
			expected:    goodgeo.NewPointEmpty(goodgeo.XYZM),
		},
		// LINESTRING tests
		{
			desc:        "parse 2D linestring",
			equivInputs: []string{"LINESTRING(0 0, 1 1, 3 4)", "LINESTRING (0 0, 1 1, 3 4)", "linestring ( 0 0, 1 1, 3 4 )"},
			expected:    goodgeo.NewLineStringFlat(goodgeo.XY, []float64{0, 0, 1, 1, 3, 4}),
		},
		{
			desc:        "parse 2D+M linestring",
			equivInputs: []string{"LINESTRING M(0 0 200, 0.1 -1 -20)", "LINESTRINGM(0 0 200, .1 -1 -20)"},
			expected:    goodgeo.NewLineStringFlat(goodgeo.XYM, []float64{0, 0, 200, 0.1, -1, -20}),
		},
		{
			desc:        "parse 3D linestring",
			equivInputs: []string{"LINESTRING(0 -1 1, 7 -1 -9)", "LINESTRING Z(0 -1 1, 7 -1 -9)", "LINESTRINGZ(0 -1 1, 7 -1 -9)"},
			expected:    goodgeo.NewLineStringFlat(goodgeo.XYZ, []float64{0, -1, 1, 7, -1, -9}),
		},
		{
			desc:        "parse 4D linestring",
			equivInputs: []string{"LINESTRING(0 0 0 0, 1 1 1 1)", "LINESTRING ZM (0 0 0 0, 1 1 1 1)", "LINESTRINGZM (0 0 0 0, 1 1 1 1)"},
			expected:    goodgeo.NewLineStringFlat(goodgeo.XYZM, []float64{0, 0, 0, 0, 1, 1, 1, 1}),
		},
		{
			desc:        "parse empty 2D linestring",
			equivInputs: []string{"LINESTRING EMPTY"},
			expected:    goodgeo.NewLineString(goodgeo.XY),
		},
		{
			desc:        "parse empty 2D+M linestring",
			equivInputs: []string{"LINESTRING M EMPTY", "LINESTRINGM EMPTY"},
			expected:    goodgeo.NewLineString(goodgeo.XYM),
		},
		{
			desc:        "parse empty 3D linestring",
			equivInputs: []string{"LINESTRING Z EMPTY", "LINESTRINGZ EMPTY"},
			expected:    goodgeo.NewLineString(goodgeo.XYZ),
		},
		{
			desc:        "parse empty 4D linestring",
			equivInputs: []string{"LINESTRING ZM EMPTY", "LINESTRINGZM EMPTY"},
			expected:    goodgeo.NewLineString(goodgeo.XYZM),
		},
		// POLYGON tests
		{
			desc:        "parse 2D polygon",
			equivInputs: []string{"POLYGON((0 0, 1 -1, 2 0, 0 0))", "POLYGON ((0 0, 1 -1, 2 0, 0 0))"},
			expected:    goodgeo.NewPolygonFlat(goodgeo.XY, []float64{0, 0, 1, -1, 2, 0, 0, 0}, []int{8}),
		},
		{
			desc:        "parse 2D polygon with hole",
			equivInputs: []string{"POLYGON((0 0, 0 100, 100 100, 100 0, 0 0),(10 10, 11 11, 12 10, 10 10))"},
			expected: goodgeo.NewPolygonFlat(goodgeo.XY,
				[]float64{0, 0, 0, 100, 100, 100, 100, 0, 0, 0, 10, 10, 11, 11, 12, 10, 10, 10}, []int{10, 18}),
		},
		{
			desc:        "parse 2D polygon with two holes",
			equivInputs: []string{"POLYGON((0 0, 0 100, 100 100, 100 0, 0 0),(10 10, 11 11, 12 10, 10 10), (2 2, 4 4, 5 1, 2 2))"},
			expected: goodgeo.NewPolygonFlat(goodgeo.XY,
				[]float64{0, 0, 0, 100, 100, 100, 100, 0, 0, 0, 10, 10, 11, 11, 12, 10, 10, 10, 2, 2, 4, 4, 5, 1, 2, 2}, []int{10, 18, 26}),
		},
		{
			desc:        "parse 2D+M polygon",
			equivInputs: []string{"POLYGONM((0 0 7, 1 -1 -50, 2 0 0, 0 0 7))", "POLYGON M ((0 0 7, 1 -1 -50, 2 0 0, 0 0 7))"},
			expected:    goodgeo.NewPolygonFlat(goodgeo.XYM, []float64{0, 0, 7, 1, -1, -50, 2, 0, 0, 0, 0, 7}, []int{12}),
		},
		{
			desc:        "parse 3D polygon",
			equivInputs: []string{"POLYGON((0 0 7, 1 -1 -50, 2 0 0, 0 0 7))", "POLYGON Z ((0 0 7, 1 -1 -50, 2 0 0, 0 0 7))"},
			expected:    goodgeo.NewPolygonFlat(goodgeo.XYZ, []float64{0, 0, 7, 1, -1, -50, 2, 0, 0, 0, 0, 7}, []int{12}),
		},
		{
			desc:        "parse 4D polygon",
			equivInputs: []string{"POLYGON((0 0 12 7, 1 -1 12 -50, 2 0 12 0, 0 0 12 7))", "POLYGON ZM ((0 0 12 7, 1 -1 12 -50, 2 0 12 0, 0 0 12 7))"},
			expected:    goodgeo.NewPolygonFlat(goodgeo.XYZM, []float64{0, 0, 12, 7, 1, -1, 12, -50, 2, 0, 12, 0, 0, 0, 12, 7}, []int{16}),
		},
		{
			desc:        "parse empty 2D polygon",
			equivInputs: []string{"POLYGON EMPTY"},
			expected:    goodgeo.NewPolygon(goodgeo.XY),
		},
		{
			desc:        "parse empty 2D+M polygon",
			equivInputs: []string{"POLYGON M EMPTY", "POLYGONM EMPTY"},
			expected:    goodgeo.NewPolygon(goodgeo.XYM),
		},
		{
			desc:        "parse empty 3D polygon",
			equivInputs: []string{"POLYGON Z EMPTY", "POLYGONZ EMPTY"},
			expected:    goodgeo.NewPolygon(goodgeo.XYZ),
		},
		{
			desc:        "parse empty 4D polygon",
			equivInputs: []string{"POLYGON ZM EMPTY", "POLYGONZM EMPTY"},
			expected:    goodgeo.NewPolygon(goodgeo.XYZM),
		},
		// MULTIPOINT tests
		{
			desc:        "parse 2D multipoint",
			equivInputs: []string{"MULTIPOINT(0 0, 1 1, 2 2)", "MULTIPOINT((0 0), 1 1, (2 2))", "MULTIPOINT (0 0, 1 1, 2 2)"},
			expected:    goodgeo.NewMultiPointFlat(goodgeo.XY, []float64{0, 0, 1, 1, 2, 2}),
		},
		{
			desc:        "parse 2D+M multipoint",
			equivInputs: []string{"MULTIPOINTM((-1 5 -16), .23 7 0)", "MULTIPOINT M (-1 5 -16, 0.23 7.0 0)"},
			expected:    goodgeo.NewMultiPointFlat(goodgeo.XYM, []float64{-1, 5, -16, 0.23, 7, 0}),
		},
		{
			desc:        "parse 3D multipoint",
			equivInputs: []string{"MULTIPOINT(2 1 3)", "MULTIPOINTZ(2 1 3)", "MULTIPOINT Z ((2 1 3))"},
			expected:    goodgeo.NewMultiPointFlat(goodgeo.XYZ, []float64{2, 1, 3}),
		},
		{
			desc:        "parse 4D multipoint",
			equivInputs: []string{"MULTIPOINT(2 -8 17 45, (0 0 0 0))", "MULTIPOINTZM((2 -8 17 45), (0 0 0 0))", "MULTIPOINT ZM (2 -8 17 45, 0 0 0 0)"},
			expected:    goodgeo.NewMultiPointFlat(goodgeo.XYZM, []float64{2, -8, 17, 45, 0, 0, 0, 0}),
		},
		{
			desc:        "parse 2D multipoint with EMPTY points",
			equivInputs: []string{"MULTIPOINT(EMPTY, 2 3, EMPTY)", "MULTIPOINT (EMPTY, (2 3), EMPTY)"},
			expected:    goodgeo.NewMultiPointFlat(goodgeo.XY, []float64{2, 3}, goodgeo.NewMultiPointFlatOptionWithEnds([]int{0, 2, 2})),
		},
		{
			desc:        "parse 2D+M multipoint with EMPTY points",
			equivInputs: []string{"MULTIPOINTM(2 3 1, EMPTY)", "MULTIPOINT M ((2 3 1), EMPTY)"},
			expected:    goodgeo.NewMultiPointFlat(goodgeo.XYM, []float64{2, 3, 1}, goodgeo.NewMultiPointFlatOptionWithEnds([]int{3, 3})),
		},
		{
			desc:        "parse 3D multipoint with EMPTY points",
			equivInputs: []string{"MULTIPOINTZ (EMPTY, EMPTY)", "MULTIPOINT Z (EMPTY, EMPTY)"},
			expected:    goodgeo.NewMultiPointFlat(goodgeo.XYZ, []float64(nil), goodgeo.NewMultiPointFlatOptionWithEnds([]int{0, 0})),
		},
		{
			desc:        "parse 4D multipoint with EMPTY points",
			equivInputs: []string{"MULTIPOINTZM(EMPTY, 1 -1 1 -1)", "MULTIPOINT ZM (EMPTY, (1 -1 1 -1))"},
			expected:    goodgeo.NewMultiPointFlat(goodgeo.XYZM, []float64{1, -1, 1, -1}, goodgeo.NewMultiPointFlatOptionWithEnds([]int{0, 4})),
		},
		{
			desc:        "parse empty 2D multipoint",
			equivInputs: []string{"MULTIPOINT EMPTY"},
			expected:    goodgeo.NewMultiPoint(goodgeo.XY),
		},
		{
			desc:        "parse empty 2D+M multipoint",
			equivInputs: []string{"MULTIPOINT M EMPTY", "MULTIPOINTM EMPTY"},
			expected:    goodgeo.NewMultiPoint(goodgeo.XYM),
		},
		{
			desc:        "parse empty 3D multipoint",
			equivInputs: []string{"MULTIPOINT Z EMPTY", "MULTIPOINTZ EMPTY"},
			expected:    goodgeo.NewMultiPoint(goodgeo.XYZ),
		},
		{
			desc:        "parse empty 4D multipoint",
			equivInputs: []string{"MULTIPOINT ZM EMPTY", "MULTIPOINTZM EMPTY"},
			expected:    goodgeo.NewMultiPoint(goodgeo.XYZM),
		},
		// MULTILINESTRING tests
		{
			desc:        "parse 2D multilinestring",
			equivInputs: []string{"MULTILINESTRING((0 0, 1 1), EMPTY)", "MULTILINESTRING (( 0 0, 1 1 ), EMPTY )"},
			expected:    goodgeo.NewMultiLineStringFlat(goodgeo.XY, []float64{0, 0, 1, 1}, []int{4, 4}),
		},
		{
			desc:        "parse 2D+M multilinestring",
			equivInputs: []string{"MULTILINESTRINGM((0 -1 -2, 2 5 7))", "multilinestring m ((0 -1 -2, 2 5 7))"},
			expected:    goodgeo.NewMultiLineStringFlat(goodgeo.XYM, []float64{0, -1, -2, 2, 5, 7}, []int{6}),
		},
		{
			desc:        "parse 3D multilinestring",
			equivInputs: []string{"MULTILINESTRING((0 -1 -2, 2 5 7))", "MULTILINESTRINGZ((0 -1 -2, 2 5 7))", "MULTILINESTRING Z ((0 -1 -2, 2 5 7))"},
			expected:    goodgeo.NewMultiLineStringFlat(goodgeo.XYZ, []float64{0, -1, -2, 2, 5, 7}, []int{6}),
		},
		{
			desc: "parse 4D multilinestring",
			equivInputs: []string{
				"MULTILINESTRING((0 0 0 0, 1 1 1 1), (-2 -3 -4 -5, 0.5 -0.75 1 -1.25, 0 1 5 7))",
				"MULTILINESTRING ZM ((0 0 0 0, 1 1 1 1), (-2 -3 -4 -5, 0.5 -0.75 1 -1.25, 0 1 5 7))",
				"multilinestringzm((0 0 0 0, 1 1 1 1), (-2 -3 -4 -5, 0.5 -0.75 1 -1.25, 0 1 5 7))",
			},
			expected: goodgeo.NewMultiLineStringFlat(goodgeo.XYZM,
				[]float64{0, 0, 0, 0, 1, 1, 1, 1, -2, -3, -4, -5, 0.5, -0.75, 1, -1.25, 0, 1, 5, 7}, []int{8, 20}),
		},
		{
			desc:        "parse 2D+M multilinestring with EMPTY linestrings",
			equivInputs: []string{"MultiLineString M ((1 -1 2, 3 -0.4 7), EMPTY, (0 0 0, -2 -4 -89))", "MULTILINESTRINGM ((1 -1 2, 3 -0.4 7), EMPTY, (0 0 0, -2 -4 -89))"},
			expected:    goodgeo.NewMultiLineStringFlat(goodgeo.XYM, []float64{1, -1, 2, 3, -0.4, 7, 0, 0, 0, -2, -4, -89}, []int{6, 6, 12}),
		},
		{
			desc:        "parse 3D multilinestring with EMPTY linestrings",
			equivInputs: []string{"MULTILINESTRINGZ(EMPTY, EMPTY, (1 1 1, 2 2 2, 3 3 3))", "multilinestring z (EMPTY, empty, (1 1 1, 2 2 2, 3 3 3))"},
			expected:    goodgeo.NewMultiLineStringFlat(goodgeo.XYZ, []float64{1, 1, 1, 2, 2, 2, 3, 3, 3}, []int{0, 0, 9}),
		},
		{
			desc:        "parse 4D multilinestring with EMPTY linestrings",
			equivInputs: []string{"MULTILINESTRINGZM(EMPTY)", "MuLTIliNeStRiNg zM (EMPTY)"},
			expected:    goodgeo.NewMultiLineStringFlat(goodgeo.XYZM, []float64(nil), []int{0}),
		},
		{
			desc:        "parse empty 2D multilinestring",
			equivInputs: []string{"MULTILINESTRING EMPTY"},
			expected:    goodgeo.NewMultiLineString(goodgeo.XY),
		},
		{
			desc:        "parse empty 2D+M multilinestring",
			equivInputs: []string{"MULTILINESTRING M EMPTY", "MULTILINESTRINGM EMPTY"},
			expected:    goodgeo.NewMultiLineString(goodgeo.XYM),
		},
		{
			desc:        "parse empty 3D multilinestring",
			equivInputs: []string{"MULTILINESTRING Z EMPTY", "MULTILINESTRINGZ EMPTY"},
			expected:    goodgeo.NewMultiLineString(goodgeo.XYZ),
		},
		{
			desc:        "parse empty 4D multilinestring",
			equivInputs: []string{"MULTILINESTRING ZM EMPTY", "MULTILINESTRINGZM EMPTY"},
			expected:    goodgeo.NewMultiLineString(goodgeo.XYZM),
		},
		// MULTIPOLYGON tests
		{
			desc:        "parse 2D multipolygon",
			equivInputs: []string{"MULTIPOLYGON(((1 0, 2 5, -2 5, 1 0)))"},
			expected:    goodgeo.NewMultiPolygonFlat(goodgeo.XY, []float64{1, 0, 2, 5, -2, 5, 1, 0}, [][]int{{8}}),
		},
		{
			desc:        "parse 2D multipolygon with EMPTY at rear",
			equivInputs: []string{"MULTIPOLYGON(((1 0, 2 5, -2 5, 1 0)), EMPTY)"},
			expected:    goodgeo.NewMultiPolygonFlat(goodgeo.XY, []float64{1, 0, 2, 5, -2, 5, 1, 0}, [][]int{{8}, []int(nil)}),
		},
		{
			desc:        "parse 2D multipolygon with EMPTY at front",
			equivInputs: []string{"MULTIPOLYGON(EMPTY, ((1 0, 2 5, -2 5, 1 0)))"},
			expected:    goodgeo.NewMultiPolygonFlat(goodgeo.XY, []float64{1, 0, 2, 5, -2, 5, 1, 0}, [][]int{[]int(nil), {8}}),
		},
		{
			desc:        "parse 2D multipolygon with multiple polygons",
			equivInputs: []string{"MULTIPOLYGON(((1 0, 2 5, -2 5, 1 0)), EMPTY, ((-1 -1, 2 7, 3 0, -1 -1)))"},
			expected:    goodgeo.NewMultiPolygonFlat(goodgeo.XY, []float64{1, 0, 2, 5, -2, 5, 1, 0, -1, -1, 2, 7, 3, 0, -1, -1}, [][]int{{8}, []int(nil), {16}}),
		},
		{
			desc:        "parse 2D+M multipolygon",
			equivInputs: []string{"MULTIPOLYGON M (((0 0 0, 1 1 1, 2 3 1, 0 0 0)))"},
			expected:    goodgeo.NewMultiPolygonFlat(goodgeo.XYM, []float64{0, 0, 0, 1, 1, 1, 2, 3, 1, 0, 0, 0}, [][]int{{12}}),
		},
		{
			desc:        "parse 3D multipolygon",
			equivInputs: []string{"MULTIPOLYGON(((0 0 0, 1 1 1, 2 3 1, 0 0 0)))", "MULTIPOLYGON Z (((0 0 0, 1 1 1, 2 3 1, 0 0 0)))"},
			expected:    goodgeo.NewMultiPolygonFlat(goodgeo.XYZ, []float64{0, 0, 0, 1, 1, 1, 2, 3, 1, 0, 0, 0}, [][]int{{12}}),
		},
		{
			desc:        "parse 4D multipolygon",
			equivInputs: []string{"MULTIPOLYGON(((0 0 0 0, 1 1 1 -1, 2 3 1 -2, 0 0 0 0)))", "MULTIPOLYGON ZM (((0 0 0 0, 1 1 1 -1, 2 3 1 -2, 0 0 0 0)))"},
			expected:    goodgeo.NewMultiPolygonFlat(goodgeo.XYZM, []float64{0, 0, 0, 0, 1, 1, 1, -1, 2, 3, 1, -2, 0, 0, 0, 0}, [][]int{{16}}),
		},
		{
			desc:        "parse empty 2D multipolygon",
			equivInputs: []string{"MULTIPOLYGON EMPTY"},
			expected:    goodgeo.NewMultiPolygon(goodgeo.XY),
		},
		{
			desc:        "parse empty 2D+M multipolygon",
			equivInputs: []string{"MULTIPOLYGON M EMPTY", "MULTIPOLYGONM EMPTY"},
			expected:    goodgeo.NewMultiPolygon(goodgeo.XYM),
		},
		{
			desc:        "parse empty 3D multipolygon",
			equivInputs: []string{"MULTIPOLYGON Z EMPTY", "MULTIPOLYGONZ EMPTY"},
			expected:    goodgeo.NewMultiPolygon(goodgeo.XYZ),
		},
		{
			desc:        "parse empty 4D multipolygon",
			equivInputs: []string{"MULTIPOLYGON ZM EMPTY", "MULTIPOLYGONZM EMPTY"},
			expected:    goodgeo.NewMultiPolygon(goodgeo.XYZM),
		},
		// GEOMETRYCOLLECTION tests
		{
			desc:        "parse 2D geometrycollection with a single point",
			equivInputs: []string{"GEOMETRYCOLLECTION(POINT(0 0))"},
			expected:    goodgeo.NewGeometryCollection().MustSetLayout(goodgeo.XY).MustPush(goodgeo.NewPointFlat(goodgeo.XY, []float64{0, 0})),
		},
		{
			desc:        "parse 2D+M base type geometrycollection",
			equivInputs: []string{"GEOMETRYCOLLECTION M (POINT M (0 0 0))", "GEOMETRYCOLLECTION(POINT M (0 0 0))"},
			expected:    goodgeo.NewGeometryCollection().MustSetLayout(goodgeo.XYM).MustPush(goodgeo.NewPointFlat(goodgeo.XYM, []float64{0, 0, 0})),
		},
		{
			desc: "parse 2D+M geometrycollection with base type empty geometry",
			equivInputs: []string{
				"GEOMETRYCOLLECTION M (LINESTRING EMPTY)",
				"GEOMETRYCOLLECTION(LINESTRING M EMPTY)",
				"GEOMETRYCOLLECTION M (LINESTRING M EMPTY)",
			},
			expected: goodgeo.NewGeometryCollection().MustSetLayout(goodgeo.XYM).MustPush(goodgeo.NewLineString(goodgeo.XYM)),
		},
		{
			desc:        "parse 3D geometrycollection with base type empty geometry",
			equivInputs: []string{"GEOMETRYCOLLECTION Z (LINESTRING EMPTY)", "GEOMETRYCOLLECTION Z (LINESTRING Z EMPTY)"},
			expected:    goodgeo.NewGeometryCollection().MustSetLayout(goodgeo.XYZ).MustPush(goodgeo.NewLineString(goodgeo.XYZ)),
		},
		{
			desc: "parse 2D+M geometrycollection with nested geometrycollection and empty geometry",
			equivInputs: []string{
				"GEOMETRYCOLLECTION(GEOMETRYCOLLECTION(LINESTRING M EMPTY), LINESTRING M EMPTY)",
				"GEOMETRYCOLLECTION(GEOMETRYCOLLECTION M (LINESTRING EMPTY), LINESTRING M EMPTY)",
				"GEOMETRYCOLLECTION(GEOMETRYCOLLECTION M (LINESTRING M EMPTY), LINESTRING M EMPTY)",
				"GEOMETRYCOLLECTION M (GEOMETRYCOLLECTION(LINESTRING EMPTY), LINESTRING EMPTY)",
				"GEOMETRYCOLLECTION M (GEOMETRYCOLLECTION M (LINESTRING EMPTY), LINESTRING EMPTY)",
				"GEOMETRYCOLLECTION M (GEOMETRYCOLLECTION(LINESTRING M EMPTY), LINESTRING EMPTY)",
				"GEOMETRYCOLLECTION M (GEOMETRYCOLLECTION(LINESTRING EMPTY), LINESTRING M EMPTY)",
				"GEOMETRYCOLLECTION M (GEOMETRYCOLLECTION(LINESTRING M EMPTY), LINESTRING M EMPTY)",
				"GEOMETRYCOLLECTION M (GEOMETRYCOLLECTION M (LINESTRING EMPTY), LINESTRING M EMPTY)",
				"GEOMETRYCOLLECTION M (GEOMETRYCOLLECTION M (LINESTRING M EMPTY), LINESTRING M EMPTY)",
			},
			expected: goodgeo.NewGeometryCollection().MustSetLayout(goodgeo.XYM).MustPush(
				goodgeo.NewGeometryCollection().MustSetLayout(goodgeo.XYM).MustPush(goodgeo.NewLineString(goodgeo.XYM)),
				goodgeo.NewLineString(goodgeo.XYM),
			),
		},
		{
			desc: "parse 2D+M geometrycollection with empty geometrycollection",
			equivInputs: []string{
				"GEOMETRYCOLLECTION M (GEOMETRYCOLLECTION EMPTY)",
				"GEOMETRYCOLLECTION M (GEOMETRYCOLLECTION M EMPTY)",
			},
			expected: goodgeo.NewGeometryCollection().MustSetLayout(goodgeo.XYM).MustPush(goodgeo.NewGeometryCollection().MustSetLayout(goodgeo.XYM)),
		},
		{
			desc: "parse 3D geometry collection with nested geometrycollection and empty geometry",
			equivInputs: []string{
				"GEOMETRYCOLLECTION(GEOMETRYCOLLECTION Z (POLYGON Z EMPTY), POINT Z EMPTY)",
				"GEOMETRYCOLLECTION(GEOMETRYCOLLECTION(POLYGON Z EMPTY), POINT Z EMPTY)",
				"GEOMETRYCOLLECTION Z (GEOMETRYCOLLECTION(POLYGON EMPTY), POINT EMPTY)",
				"GEOMETRYCOLLECTION Z (GEOMETRYCOLLECTION Z (POLYGON EMPTY), POINT EMPTY)",
				"GEOMETRYCOLLECTION Z (GEOMETRYCOLLECTION Z (POLYGON Z EMPTY), POINT Z EMPTY)",
			},
			expected: goodgeo.NewGeometryCollection().MustSetLayout(goodgeo.XYZ).MustPush(
				goodgeo.NewGeometryCollection().MustSetLayout(goodgeo.XYZ).MustPush(goodgeo.NewPolygon(goodgeo.XYZ)),
				goodgeo.NewPointEmpty(goodgeo.XYZ),
			),
		},
		{
			desc: "parse 2D geometrycollection",
			equivInputs: []string{`GEOMETRYCOLLECTION(
POINT(0 0),
LINESTRING(1 1, 0 0, 1 4),
POLYGON((0 0, 0 100, 100 100, 100 0, 0 0), (10 10, 11 11, 12 10, 10 10), (2 2, 4 4, 5 1, 2 2)),
MULTIPOINT((23 24), EMPTY),
MULTILINESTRING((1 1, 0 0, 1 4)),
MULTIPOLYGON(((0 0, 0 100, 100 100, 100 0, 0 0))),
GEOMETRYCOLLECTION EMPTY
)`},
			expected: goodgeo.NewGeometryCollection().MustSetLayout(goodgeo.XY).MustPush(
				goodgeo.NewPointFlat(goodgeo.XY, []float64{0, 0}),
				goodgeo.NewLineStringFlat(goodgeo.XY, []float64{1, 1, 0, 0, 1, 4}),
				goodgeo.NewPolygonFlat(goodgeo.XY,
					[]float64{0, 0, 0, 100, 100, 100, 100, 0, 0, 0, 10, 10, 11, 11, 12, 10, 10, 10, 2, 2, 4, 4, 5, 1, 2, 2},
					[]int{10, 18, 26}),
				goodgeo.NewMultiPointFlat(goodgeo.XY, []float64{23, 24}, goodgeo.NewMultiPointFlatOptionWithEnds([]int{2, 2})),
				goodgeo.NewMultiLineStringFlat(goodgeo.XY, []float64{1, 1, 0, 0, 1, 4}, []int{6}),
				goodgeo.NewMultiPolygonFlat(goodgeo.XY, []float64{0, 0, 0, 100, 100, 100, 100, 0, 0, 0}, [][]int{{10}}),
				goodgeo.NewGeometryCollection().MustSetLayout(goodgeo.XY),
			),
		},
		{
			desc:        "parse 2D geometrycollection with nested geometrycollection",
			equivInputs: []string{"GEOMETRYCOLLECTION(POINT(0 0), GEOMETRYCOLLECTION(MULTIPOINT(EMPTY, 2 1)))"},
			expected: goodgeo.NewGeometryCollection().MustSetLayout(goodgeo.XY).MustPush(
				goodgeo.NewPointFlat(goodgeo.XY, []float64{0, 0}),
				goodgeo.NewGeometryCollection().MustSetLayout(goodgeo.XY).MustPush(
					goodgeo.NewMultiPointFlat(goodgeo.XY, []float64{2, 1}, goodgeo.NewMultiPointFlatOptionWithEnds([]int{0, 2})),
				),
			),
		},
		{
			desc: "parse 2D+M geometrycollection",
			equivInputs: []string{
				`GEOMETRYCOLLECTION M (
POINT EMPTY,
POINT M (-2 0 0.5),
LINESTRING M (0 0 200, 0.1 -1 -20),
POLYGON M ((0 0 7, 1 -1 -50, 2 0 0, 0 0 7)),
MULTIPOINT M (-1 5 -16, 0.23 7.0 0),
MULTILINESTRING M ((0 -1 -2, 2 5 7)),
MULTIPOLYGON M (((0 0 0, 1 1 1, 2 3 1, 0 0 0)))
)`,
				`GEOMETRYCOLLECTION M (
POINT M EMPTY,
POINT M (-2 0 0.5),
LINESTRING M (0 0 200, 0.1 -1 -20),
POLYGON M ((0 0 7, 1 -1 -50, 2 0 0, 0 0 7)),
MULTIPOINT M (-1 5 -16, 0.23 7.0 0),
MULTILINESTRING M ((0 -1 -2, 2 5 7)),
MULTIPOLYGON M (((0 0 0, 1 1 1, 2 3 1, 0 0 0)))
)`,
				`GEOMETRYCOLLECTION(
POINT M EMPTY,
POINT M (-2 0 0.5),
LINESTRING M (0 0 200, 0.1 -1 -20),
POLYGON M ((0 0 7, 1 -1 -50, 2 0 0, 0 0 7)),
MULTIPOINT M (-1 5 -16, 0.23 7.0 0),
MULTILINESTRING M ((0 -1 -2, 2 5 7)),
MULTIPOLYGON M (((0 0 0, 1 1 1, 2 3 1, 0 0 0)))
)`,
			},
			expected: goodgeo.NewGeometryCollection().MustSetLayout(goodgeo.XYM).MustPush(
				goodgeo.NewPointEmpty(goodgeo.XYM),
				goodgeo.NewPointFlat(goodgeo.XYM, []float64{-2, 0, 0.5}),
				goodgeo.NewLineStringFlat(goodgeo.XYM, []float64{0, 0, 200, 0.1, -1, -20}),
				goodgeo.NewPolygonFlat(goodgeo.XYM, []float64{0, 0, 7, 1, -1, -50, 2, 0, 0, 0, 0, 7}, []int{12}),
				goodgeo.NewMultiPointFlat(goodgeo.XYM, []float64{-1, 5, -16, 0.23, 7, 0}),
				goodgeo.NewMultiLineStringFlat(goodgeo.XYM, []float64{0, -1, -2, 2, 5, 7}, []int{6}),
				goodgeo.NewMultiPolygonFlat(goodgeo.XYM, []float64{0, 0, 0, 1, 1, 1, 2, 3, 1, 0, 0, 0}, [][]int{{12}}),
			),
		},
		{
			desc: "parse 2D+M geometrycollection with nested geometrycollection",
			equivInputs: []string{
				"GEOMETRYCOLLECTION(GEOMETRYCOLLECTION M (POINT EMPTY, LINESTRING M (0 0 0, 1 1 1)))",
				"GEOMETRYCOLLECTION(GEOMETRYCOLLECTION M (POINT M EMPTY, LINESTRING M (0 0 0, 1 1 1)))",
				"GEOMETRYCOLLECTION(GEOMETRYCOLLECTION(POINT M EMPTY, LINESTRING M (0 0 0, 1 1 1)))",
				"GEOMETRYCOLLECTION M (GEOMETRYCOLLECTION(POINT M EMPTY, LINESTRING M (0 0 0, 1 1 1)))",
				"GEOMETRYCOLLECTION M (GEOMETRYCOLLECTION M (POINT M EMPTY, LINESTRING M (0 0 0, 1 1 1)))",
			},
			expected: goodgeo.NewGeometryCollection().MustSetLayout(goodgeo.XYM).MustPush(
				goodgeo.NewGeometryCollection().MustSetLayout(goodgeo.XYM).MustPush(
					goodgeo.NewPointEmpty(goodgeo.XYM),
					goodgeo.NewLineStringFlat(goodgeo.XYM, []float64{0, 0, 0, 1, 1, 1}),
				),
			),
		},
		{
			desc: "parse 3D geometrycollection",
			equivInputs: []string{
				`GEOMETRYCOLLECTION Z (
POINT Z (2 3 4),
LINESTRING Z (0 -1 1, 7 -1 -9),
POLYGON Z ((0 0 7, 1 -1 -50, 2 0 0, 0 0 7)),
MULTIPOINT Z ((2 3 1), EMPTY),
MULTILINESTRING Z (EMPTY, EMPTY, (1 1 1, 2 2 2, 3 3 3)),
MULTIPOLYGON Z (((0 0 0, 1 1 1, 2 3 1, 0 0 0)))
)`,
				`GEOMETRYCOLLECTION Z (
POINT(2 3 4),
LINESTRING(0 -1 1, 7 -1 -9),
POLYGON((0 0 7, 1 -1 -50, 2 0 0, 0 0 7)),
MULTIPOINT((2 3 1), EMPTY),
MULTILINESTRING(EMPTY, EMPTY, (1 1 1, 2 2 2, 3 3 3)),
MULTIPOLYGON(((0 0 0, 1 1 1, 2 3 1, 0 0 0)))
)`,
				`GEOMETRYCOLLECTION(
POINT Z (2 3 4),
LINESTRING Z (0 -1 1, 7 -1 -9),
POLYGON Z ((0 0 7, 1 -1 -50, 2 0 0, 0 0 7)),
MULTIPOINT Z ((2 3 1), EMPTY),
MULTILINESTRING Z (EMPTY, EMPTY, (1 1 1, 2 2 2, 3 3 3)),
MULTIPOLYGON Z (((0 0 0, 1 1 1, 2 3 1, 0 0 0)))
)`,
				`GEOMETRYCOLLECTION(
POINT(2 3 4),
LINESTRING(0 -1 1, 7 -1 -9),
POLYGON((0 0 7, 1 -1 -50, 2 0 0, 0 0 7)),
MULTIPOINT Z ((2 3 1), EMPTY),
MULTILINESTRING Z (EMPTY, EMPTY, (1 1 1, 2 2 2, 3 3 3)),
MULTIPOLYGON(((0 0 0, 1 1 1, 2 3 1, 0 0 0)))
)`,
			},
			expected: goodgeo.NewGeometryCollection().MustSetLayout(goodgeo.XYZ).MustPush(
				goodgeo.NewPointFlat(goodgeo.XYZ, []float64{2, 3, 4}),
				goodgeo.NewLineStringFlat(goodgeo.XYZ, []float64{0, -1, 1, 7, -1, -9}),
				goodgeo.NewPolygonFlat(goodgeo.XYZ, []float64{0, 0, 7, 1, -1, -50, 2, 0, 0, 0, 0, 7}, []int{12}),
				goodgeo.NewMultiPointFlat(goodgeo.XYZ, []float64{2, 3, 1}, goodgeo.NewMultiPointFlatOptionWithEnds([]int{3, 3})),
				goodgeo.NewMultiLineStringFlat(goodgeo.XYZ, []float64{1, 1, 1, 2, 2, 2, 3, 3, 3}, []int{0, 0, 9}),
				goodgeo.NewMultiPolygonFlat(goodgeo.XYZ, []float64{0, 0, 0, 1, 1, 1, 2, 3, 1, 0, 0, 0}, [][]int{{12}}),
			),
		},
		{
			desc: "parse 4D geometrycollection",
			equivInputs: []string{
				`GEOMETRYCOLLECTION ZM (
POINT ZM (0 5 -10 15),
LINESTRING ZM (0 0 0 0, 1 1 1 1),
POLYGON ZM ((0 0 12 7, 1 -1 12 -50, 2 0 12 0, 0 0 12 7)),
MULTIPOINT ZM ((2 -8 17 45), (0 0 0 0)),
MULTILINESTRING ZM ((0 0 0 0, 1 1 1 1), (-2 -3 -4 -5, 0.5 -0.75 1 -1.25, 0 1 5 7)),
MULTIPOLYGON ZM (((0 0 0 0, 1 1 1 -1, 2 3 1 -2, 0 0 0 0)))
)`,
				`GEOMETRYCOLLECTION ZM (
POINT(0 5 -10 15),
LINESTRING(0 0 0 0, 1 1 1 1),
POLYGON((0 0 12 7, 1 -1 12 -50, 2 0 12 0, 0 0 12 7)),
MULTIPOINT((2 -8 17 45), (0 0 0 0)),
MULTILINESTRING((0 0 0 0, 1 1 1 1), (-2 -3 -4 -5, 0.5 -0.75 1 -1.25, 0 1 5 7)),
MULTIPOLYGON(((0 0 0 0, 1 1 1 -1, 2 3 1 -2, 0 0 0 0)))
)`,
				`GEOMETRYCOLLECTION(
POINT(0 5 -10 15),
LINESTRING(0 0 0 0, 1 1 1 1),
POLYGON((0 0 12 7, 1 -1 12 -50, 2 0 12 0, 0 0 12 7)),
MULTIPOINT((2 -8 17 45), (0 0 0 0)),
MULTILINESTRING((0 0 0 0, 1 1 1 1), (-2 -3 -4 -5, 0.5 -0.75 1 -1.25, 0 1 5 7)),
MULTIPOLYGON(((0 0 0 0, 1 1 1 -1, 2 3 1 -2, 0 0 0 0)))
)`,
				`GEOMETRYCOLLECTION(
POINT ZM (0 5 -10 15),
LINESTRING(0 0 0 0, 1 1 1 1),
POLYGON((0 0 12 7, 1 -1 12 -50, 2 0 12 0, 0 0 12 7)),
MULTIPOINT((2 -8 17 45), (0 0 0 0)),
MULTILINESTRING((0 0 0 0, 1 1 1 1), (-2 -3 -4 -5, 0.5 -0.75 1 -1.25, 0 1 5 7)),
MULTIPOLYGON(((0 0 0 0, 1 1 1 -1, 2 3 1 -2, 0 0 0 0)))
)`,
			},
			expected: goodgeo.NewGeometryCollection().MustSetLayout(goodgeo.XYZM).MustPush(
				goodgeo.NewPointFlat(goodgeo.XYZM, []float64{0, 5, -10, 15}),
				goodgeo.NewLineStringFlat(goodgeo.XYZM, []float64{0, 0, 0, 0, 1, 1, 1, 1}),
				goodgeo.NewPolygonFlat(goodgeo.XYZM, []float64{0, 0, 12, 7, 1, -1, 12, -50, 2, 0, 12, 0, 0, 0, 12, 7}, []int{16}),
				goodgeo.NewMultiPointFlat(goodgeo.XYZM, []float64{2, -8, 17, 45, 0, 0, 0, 0}),
				goodgeo.NewMultiLineStringFlat(goodgeo.XYZM,
					[]float64{0, 0, 0, 0, 1, 1, 1, 1, -2, -3, -4, -5, 0.5, -0.75, 1, -1.25, 0, 1, 5, 7}, []int{8, 20}),
				goodgeo.NewMultiPolygonFlat(goodgeo.XYZM, []float64{0, 0, 0, 0, 1, 1, 1, -1, 2, 3, 1, -2, 0, 0, 0, 0}, [][]int{{16}}),
			),
		},
		{
			desc:        "parse empty 2D geometrycollection",
			equivInputs: []string{"GEOMETRYCOLLECTION EMPTY"},
			expected:    goodgeo.NewGeometryCollection().MustSetLayout(goodgeo.XY),
		},
		{
			desc:        "parse empty 2D+M geometrycollection",
			equivInputs: []string{"GEOMETRYCOLLECTION M EMPTY", "GEOMETRYCOLLECTIONM EMPTY"},
			expected:    goodgeo.NewGeometryCollection().MustSetLayout(goodgeo.XYM),
		},
		{
			desc:        "parse empty 3D geometrycollection",
			equivInputs: []string{"GEOMETRYCOLLECTION Z EMPTY", "GEOMETRYCOLLECTIONZ EMPTY"},
			expected:    goodgeo.NewGeometryCollection().MustSetLayout(goodgeo.XYZ),
		},
		{
			desc:        "parse empty 4D geometrycollection",
			equivInputs: []string{"GEOMETRYCOLLECTION ZM EMPTY", "GEOMETRYCOLLECTIONZM EMPTY"},
			expected:    goodgeo.NewGeometryCollection().MustSetLayout(goodgeo.XYZM),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			want := tc.expected
			for _, input := range tc.equivInputs {
				got, err := Unmarshal(input)
				assert.NoError(t, err)
				assert.Equal(t, want, got)
			}
		})
	}
}

func TestUnmarshalError(t *testing.T) {
	errorTestCases := []struct {
		desc           string
		input          string
		expectedErrStr string
	}{
		// LexError
		{
			desc:  "invalid character",
			input: "POINT{0 0}",
			expectedErrStr: `syntax error: invalid character at line 1, pos 5
LINE 1: POINT{0 0}
             ^`,
		},
		{
			desc:  "invalid keyword",
			input: "DOT(0 0)",
			expectedErrStr: `syntax error: invalid keyword at line 1, pos 0
LINE 1: DOT(0 0)
        ^`,
		},
		{
			desc:  "invalid number",
			input: "POINT(2 2.3.7)",
			expectedErrStr: `syntax error: invalid number at line 1, pos 8
LINE 1: POINT(2 2.3.7)
                ^`,
		},
		{
			desc:  "invalid scientific notation number missing number before the e",
			input: "POINT(e-1 2)",
			expectedErrStr: `syntax error: invalid keyword at line 1, pos 6
LINE 1: POINT(e-1 2)
              ^`,
		},
		{
			desc:  "invalid scientific notation number with non-integer power",
			input: "POINT(5e-1.5 2)",
			expectedErrStr: `syntax error: invalid number at line 1, pos 6
LINE 1: POINT(5e-1.5 2)
              ^`,
		},
		{
			desc:  "invalid number with a + at the start (PostGIS does not allow this)",
			input: "POINT(+1 2)",
			expectedErrStr: `syntax error: invalid character at line 1, pos 6
LINE 1: POINT(+1 2)
              ^`,
		},
		{
			desc:  "invalid keyword when extraneous spaces are present in ZM",
			input: "POINT Z M (1 1 1 1)",
			expectedErrStr: `syntax error: invalid keyword at line 1, pos 8
LINE 1: POINT Z M (1 1 1 1)
                ^`,
		},
		{
			desc: "invalid geometry type split over multiple lines",
			input: `POINT
Z
       M (
          0
          0
)`,
			expectedErrStr: `syntax error: invalid keyword at line 3, pos 7
LINE 3:        M (
               ^`,
		},
		{
			desc:  "invalid keyword towards the front of a very long line",
			input: "POINT(aslfaskfjhaskfjhaksjfhkajshfkjahskfjahskfjhaksjfhkajshfkajhsfkjahskfjhaskfjhaksjhfkajshfkj)",
			expectedErrStr: `syntax error: invalid keyword at line 1, pos 6
LINE 1: POINT(aslfaskfjhaskfjhaksjfhkajshfkj...
              ^`,
		},
		{
			desc:  "invalid character towards the end of a very long line",
			input: "MULTIPOINT(0 0, 0 0, 0 0, 0 0, 0 0, 0 0, 0 0, 0 0, 0 0, 0 0, 0 0}",
			expectedErrStr: `syntax error: invalid character at line 1, pos 64
LINE 1: ..., 0 0, 0 0, 0 0, 0 0, 0 0, 0 0}
                                         ^`,
		},
		// ParseError
		{
			desc:  "invalid point",
			input: "POINT P",
			expectedErrStr: `syntax error: invalid keyword at line 1, pos 6
LINE 1: POINT P
              ^`,
		},
		{
			desc:  "point missing closing bracket",
			input: "POINT(0 0",
			expectedErrStr: `syntax error: unexpected $end, expecting ')' at line 1, pos 9
LINE 1: POINT(0 0
                 ^`,
		},
		{
			desc:  "2D point with extra comma",
			input: "POINT(0, 0)",
			expectedErrStr: `syntax error: not enough coordinates at line 1, pos 7
LINE 1: POINT(0, 0)
               ^
HINT: each point needs at least 2 coords`,
		},
		{
			desc:  "2D linestring with no points",
			input: "LINESTRING()",
			expectedErrStr: `syntax error: unexpected ')', expecting NUM at line 1, pos 11
LINE 1: LINESTRING()
                   ^`,
		},
		{
			desc:  "2D linestring with not enough points",
			input: "LINESTRING(0 0)",
			expectedErrStr: `syntax error: non-empty linestring with only one point at line 1, pos 14
LINE 1: LINESTRING(0 0)
                      ^
HINT: minimum number of points is 2`,
		},
		{
			desc:  "linestring with mixed dimensionality",
			input: "LINESTRING(0 0, 1 1 1)",
			expectedErrStr: `syntax error: mixed dimensionality, parsed layout is XY so expecting 2 coords but got 3 coords at line 1, pos 21
LINE 1: LINESTRING(0 0, 1 1 1)
                             ^`,
		},
		{
			desc:  "2D polygon with not enough points",
			input: "POLYGON((0 0, 1 1, 2 0))",
			expectedErrStr: `syntax error: polygon ring doesn't have enough points at line 1, pos 22
LINE 1: POLYGON((0 0, 1 1, 2 0))
                              ^
HINT: minimum number of points is 4`,
		},
		{
			desc:  "2D polygon with ring that isn't closed",
			input: "POLYGON((0 0, 1 1, 2 0, 1 -1))",
			expectedErrStr: `syntax error: polygon ring not closed at line 1, pos 28
LINE 1: POLYGON((0 0, 1 1, 2 0, 1 -1))
                                    ^
HINT: ensure first and last point are the same`,
		},
		{
			desc:  "2D polygon with empty second ring",
			input: "POLYGON((0 0, 1 -1, 2 0, 0 0), ())",
			expectedErrStr: `syntax error: unexpected ')', expecting NUM at line 1, pos 32
LINE 1: ...LYGON((0 0, 1 -1, 2 0, 0 0), ())
                                         ^`,
		},
		{
			desc:  "2D polygon with EMPTY as second ring",
			input: "POLYGON((0 0, 1 -1, 2 0, 0 0), EMPTY)",
			expectedErrStr: `syntax error: unexpected EMPTY, expecting '(' at line 1, pos 31
LINE 1: ...OLYGON((0 0, 1 -1, 2 0, 0 0), EMPTY)
                                         ^`,
		},
		{
			desc:  "2D polygon with invalid second ring",
			input: "POLYGON((0 0, 1 -1, 2 0, 0 0), (0.5 -0.5))",
			expectedErrStr: `syntax error: polygon ring doesn't have enough points at line 1, pos 40
LINE 1: ... 0, 1 -1, 2 0, 0 0), (0.5 -0.5))
                                         ^
HINT: minimum number of points is 4`,
		},
		{
			desc:  "2D multipoint without any points",
			input: "MULTIPOINT()",
			expectedErrStr: `syntax error: unexpected ')', expecting EMPTY or NUM or '(' at line 1, pos 11
LINE 1: MULTIPOINT()
                   ^`,
		},
		{
			desc:  "3D multipoint without comma separating points",
			input: "MULTIPOINT Z (0 0 0 0 0 0)",
			expectedErrStr: `syntax error: too many coordinates at line 1, pos 25
LINE 1: MULTIPOINT Z (0 0 0 0 0 0)
                                 ^
HINT: each point can have at most 4 coords`,
		},
		{
			desc:  "2D multipoint with EMPTY inside extraneous parentheses",
			input: "MULTIPOINT((EMPTY))",
			expectedErrStr: `syntax error: unexpected EMPTY, expecting NUM at line 1, pos 12
LINE 1: MULTIPOINT((EMPTY))
                    ^`,
		},
		{
			desc:  "3D multipoint using EMPTY as a point without using Z in type",
			input: "MULTIPOINT(0 0 0, EMPTY)",
			expectedErrStr: `syntax error: mixed dimensionality, parsed layout is XYZ but encountered layout of XY at line 1, pos 18
LINE 1: MULTIPOINT(0 0 0, EMPTY)
                          ^
HINT: EMPTY is XY layout in base geometry type`,
		},
		{
			desc:  "multipoint with mixed dimensionality",
			input: "MULTIPOINT(0 0 0, 1 1)",
			expectedErrStr: `syntax error: mixed dimensionality, parsed layout is XYZ so expecting 3 coords but got 2 coords at line 1, pos 21
LINE 1: MULTIPOINT(0 0 0, 1 1)
                             ^`,
		},
		{
			desc:  "2D multilinestring containing linestring with no points",
			input: "MULTILINESTRING(())",
			expectedErrStr: `syntax error: unexpected ')', expecting NUM at line 1, pos 17
LINE 1: MULTILINESTRING(())
                         ^`,
		},
		{
			desc:  "2D multilinestring containing linestring with only one point",
			input: "MULTILINESTRING((0 0))",
			expectedErrStr: `syntax error: non-empty linestring with only one point at line 1, pos 20
LINE 1: MULTILINESTRING((0 0))
                            ^
HINT: minimum number of points is 2`,
		},
		{
			desc:  "4D multilinestring using EMPTY without using ZM in type",
			input: "MULTILINESTRING(EMPTY, (0 0 0 0, 2 3 -2 -3))",
			expectedErrStr: `syntax error: mixed dimensionality, parsed layout is XY so expecting 2 coords but got 4 coords at line 1, pos 31
LINE 1: ...ULTILINESTRING(EMPTY, (0 0 0 0, 2 3 -2 -3))
                                         ^`,
		},
		{
			desc:  "2D multipolygon with no polygons",
			input: "MULTIPOLYGON()",
			expectedErrStr: `syntax error: unexpected ')', expecting EMPTY or '(' at line 1, pos 13
LINE 1: MULTIPOLYGON()
                     ^`,
		},
		{
			desc:  "2D multipolygon with one polygon missing outer parentheses",
			input: "MULTIPOLYGON((1 0, 2 5, -2 5, 1 0))",
			expectedErrStr: `syntax error: unexpected NUM, expecting '(' at line 1, pos 14
LINE 1: MULTIPOLYGON((1 0, 2 5, -2 5, 1 0))
                      ^`,
		},
		{
			desc:  "multipolygon with mixed dimensionality",
			input: "MULTIPOLYGON(((1 0, 2 5, -2 5, 1 0)), ((1 0 2, 2 5 1, -2 5 -1, 1 0 2)))",
			expectedErrStr: `syntax error: mixed dimensionality, parsed layout is XY so expecting 2 coords but got 3 coords at line 1, pos 45
LINE 1: ...1 0, 2 5, -2 5, 1 0)), ((1 0 2, 2 5 1, -2 5 -1, 1 0 2)))
                                         ^`,
		},
		{
			desc:  "2D multipolygon with polygon that doesn't have enough points",
			input: "MULTIPOLYGON(((0 0, 1 1, 2 0)))",
			expectedErrStr: `syntax error: polygon ring doesn't have enough points at line 1, pos 28
LINE 1: MULTIPOLYGON(((0 0, 1 1, 2 0)))
                                    ^
HINT: minimum number of points is 4`,
		},
		{
			desc:  "2D multipolygon with polygon with ring that isn't closed",
			input: "MULTIPOLYGON(((0 0, 1 1, 2 0, 1 -1)))",
			expectedErrStr: `syntax error: polygon ring not closed at line 1, pos 34
LINE 1: ...IPOLYGON(((0 0, 1 1, 2 0, 1 -1)))
                                         ^
HINT: ensure first and last point are the same`,
		},
		{
			desc:  "2D multipolygon with polygon with empty second ring",
			input: "MULTIPOLYGON(((0 0, 1 -1, 2 0, 0 0), ()))",
			expectedErrStr: `syntax error: unexpected ')', expecting NUM at line 1, pos 38
LINE 1: ...YGON(((0 0, 1 -1, 2 0, 0 0), ()))
                                         ^`,
		},
		{
			desc:  "2D multipolygon with polygon with EMPTY as second ring",
			input: "MULTIPOLYGON(((0 0, 1 -1, 2 0, 0 0), EMPTY))",
			expectedErrStr: `syntax error: unexpected EMPTY, expecting '(' at line 1, pos 37
LINE 1: ...LYGON(((0 0, 1 -1, 2 0, 0 0), EMPTY))
                                         ^`,
		},
		{
			desc:  "2D multipolygon with polygon with invalid second ring",
			input: "MULTIPOLYGON(((0 0, 1 -1, 2 0, 0 0), (0.5 -0.5)))",
			expectedErrStr: `syntax error: polygon ring doesn't have enough points at line 1, pos 46
LINE 1: ... 0, 1 -1, 2 0, 0 0), (0.5 -0.5)))
                                         ^
HINT: minimum number of points is 4`,
		},
		{
			desc:  "3D multipolygon using EMPTY without using Z in its type",
			input: "MULTIPOLYGON(EMPTY, ((0 0 0, 1 1 1, 2 3 1, 0 0 0)))",
			expectedErrStr: `syntax error: mixed dimensionality, parsed layout is XY so expecting 2 coords but got 3 coords at line 1, pos 27
LINE 1: MULTIPOLYGON(EMPTY, ((0 0 0, 1 1 1, 2 3 1, 0 0 0)))
                                   ^`,
		},
		{
			desc:  "2D geometrycollection with EMPTY item",
			input: "GEOMETRYCOLLECTION(EMPTY)",
			expectedErrStr: `syntax error: unexpected EMPTY at line 1, pos 19
LINE 1: GEOMETRYCOLLECTION(EMPTY)
                           ^`,
		},
		{
			desc:  "3D geometrycollection with no items",
			input: "GEOMETRYCOLLECTION Z ()",
			expectedErrStr: `syntax error: unexpected ')' at line 1, pos 22
LINE 1: GEOMETRYCOLLECTION Z ()
                              ^`,
		},
		{
			desc:  "base type geometrycollection with mixed dimensionality",
			input: "GEOMETRYCOLLECTION(POINT M (0 0 0), LINESTRING(0 0, 1 1))",
			expectedErrStr: `syntax error: mixed dimensionality, parsed layout is XYM but encountered layout of not XYM at line 1, pos 36
LINE 1: ...RYCOLLECTION(POINT M (0 0 0), LINESTRING(0 0, 1 1))
                                         ^
HINT: the M variant is required for non-empty XYM geometries in GEOMETRYCOLLECTIONs`,
		},
		{
			desc:  "2D+M geometrycollection with 3 coords point missing M type",
			input: "GEOMETRYCOLLECTION M (POINT(0 0 0))",
			expectedErrStr: `syntax error: mixed dimensionality, parsed layout is XYM but encountered layout of not XYM at line 1, pos 27
LINE 1: GEOMETRYCOLLECTION M (POINT(0 0 0))
                                   ^
HINT: the M variant is required for non-empty XYM geometries in GEOMETRYCOLLECTIONs`,
		},
		{
			desc:  "2D+M geometrycollection with 3 coords linestring missing M type",
			input: "GEOMETRYCOLLECTION M (LINESTRING(0 0 0, 1 1 1))",
			expectedErrStr: `syntax error: mixed dimensionality, parsed layout is XYM but encountered layout of not XYM at line 1, pos 32
LINE 1: ...OMETRYCOLLECTION M (LINESTRING(0 0 0, 1 1 1))
                                         ^
HINT: the M variant is required for non-empty XYM geometries in GEOMETRYCOLLECTIONs`,
		},
		{
			desc:  "2D+M geometrycollection with 3 coords polygon missing M type",
			input: "GEOMETRYCOLLECTION M (POLYGON((0 0 0, 1 1 1, 2 3 1, 0 0 0)))",
			expectedErrStr: `syntax error: mixed dimensionality, parsed layout is XYM but encountered layout of not XYM at line 1, pos 29
LINE 1: GEOMETRYCOLLECTION M (POLYGON((0 0 0, 1 1 1, 2 3 1, 0 0 0))...
                                     ^
HINT: the M variant is required for non-empty XYM geometries in GEOMETRYCOLLECTIONs`,
		},
		{
			desc:  "2D+M geometrycollection with 3 coords multipoint missing M type",
			input: "GEOMETRYCOLLECTION M (MULTIPOINT((0 0 0), 1 1 1))",
			expectedErrStr: `syntax error: mixed dimensionality, parsed layout is XYM but encountered layout of not XYM at line 1, pos 32
LINE 1: ...OMETRYCOLLECTION M (MULTIPOINT((0 0 0), 1 1 1))
                                         ^
HINT: the M variant is required for non-empty XYM geometries in GEOMETRYCOLLECTIONs`,
		},
		{
			desc:  "2D+M geometrycollection with 3 coords multilinestring missing M type",
			input: "GEOMETRYCOLLECTION M (MULTILINESTRING((0 0 0, 1 1 1)))",
			expectedErrStr: `syntax error: mixed dimensionality, parsed layout is XYM but encountered layout of not XYM at line 1, pos 37
LINE 1: ...YCOLLECTION M (MULTILINESTRING((0 0 0, 1 1 1)))
                                         ^
HINT: the M variant is required for non-empty XYM geometries in GEOMETRYCOLLECTIONs`,
		},
		{
			desc:  "2D+M geometrycollection with 3 coords multipolygon missing M type",
			input: "GEOMETRYCOLLECTION M (MULTIPOLYGON(((0 0 0, 1 1 1, 2 3 1, 0 0 0))))",
			expectedErrStr: `syntax error: mixed dimensionality, parsed layout is XYM but encountered layout of not XYM at line 1, pos 34
LINE 1: ...ETRYCOLLECTION M (MULTIPOLYGON(((0 0 0, 1 1 1, 2 3 1, 0 0 0)...
                                         ^
HINT: the M variant is required for non-empty XYM geometries in GEOMETRYCOLLECTIONs`,
		},
		{
			desc:  "3D geometrycollection with mixed dimensionality in nested geometry collection",
			input: "GEOMETRYCOLLECTION Z (GEOMETRYCOLLECTION(POINT(0 0)))",
			expectedErrStr: `syntax error: mixed dimensionality, parsed layout is XYZ so expecting 3 coords but got 2 coords at line 1, pos 50
LINE 1: ... (GEOMETRYCOLLECTION(POINT(0 0)))
                                         ^`,
		},
		{
			desc:  "base type geometrycollection with 3D geometry and base type EMPTY geometry",
			input: "GEOMETRYCOLLECTION(POINT(0 0 0), LINESTRING EMPTY)",
			expectedErrStr: `syntax error: mixed dimensionality, parsed layout is XYZ but encountered layout of XY at line 1, pos 44
LINE 1: ...TION(POINT(0 0 0), LINESTRING EMPTY)
                                         ^
HINT: EMPTY is XY layout in base geometry type`,
		},
		{
			desc:  "base type geometrycollection with base type EMPTY geometry and 3D geometry",
			input: "GEOMETRYCOLLECTION(LINESTRING EMPTY, POINT(0 0 0))",
			expectedErrStr: `syntax error: mixed dimensionality, parsed layout is XY so expecting 2 coords but got 3 coords at line 1, pos 48
LINE 1: ...(LINESTRING EMPTY, POINT(0 0 0))
                                         ^`,
		},
		{
			desc:  "2D+M geometrycollection with base type multipoint with mixed dimensionality",
			input: "GEOMETRYCOLLECTIONM(LINESTRING EMPTY, MULTIPOINT(EMPTY, (0 0 0)))",
			expectedErrStr: `syntax error: mixed dimensionality, parsed layout is XYM but encountered layout of not XYM at line 1, pos 48
LINE 1: ...M(LINESTRING EMPTY, MULTIPOINT(EMPTY, (0 0 0)))
                                         ^
HINT: the M variant is required for non-empty XYM geometries in GEOMETRYCOLLECTIONs`,
		},
		{
			desc:  "geometrycollection with mixed dimensionality between nested geometrycollection and EMPTY linestring 1",
			input: "GEOMETRYCOLLECTION(GEOMETRYCOLLECTION M (LINESTRING EMPTY), LINESTRING EMPTY)",
			expectedErrStr: `syntax error: mixed dimensionality, parsed layout is XYM but encountered layout of not XYM at line 1, pos 60
LINE 1: ...LECTION M (LINESTRING EMPTY), LINESTRING EMPTY)
                                         ^
HINT: the M variant is required for non-empty XYM geometries in GEOMETRYCOLLECTIONs`,
		},
		{
			desc:  "geometrycollection with mixed dimensionality between nested geometrycollection and EMPTY linestring 2",
			input: "GEOMETRYCOLLECTION(GEOMETRYCOLLECTION(LINESTRING M EMPTY), LINESTRING EMPTY)",
			expectedErrStr: `syntax error: mixed dimensionality, parsed layout is XYM but encountered layout of not XYM at line 1, pos 59
LINE 1: ...LLECTION(LINESTRING M EMPTY), LINESTRING EMPTY)
                                         ^
HINT: the M variant is required for non-empty XYM geometries in GEOMETRYCOLLECTIONs`,
		},
		{
			desc:  "geometrycollection with mixed dimensionality between nested geometrycollection and EMPTY linestring 3",
			input: "GEOMETRYCOLLECTION(GEOMETRYCOLLECTION(LINESTRING EMPTY), LINESTRING M EMPTY)",
			expectedErrStr: `syntax error: mixed dimensionality, parsed layout is XY but encountered layout of XYM at line 1, pos 57
LINE 1: ...COLLECTION(LINESTRING EMPTY), LINESTRING M EMPTY)
                                         ^`,
		},
		{
			desc: "geometrycollection with mixed dimensionality with multiple lines",
			input: `GEOMETRYCOLLECTION M (
	POINT EMPTY,
	POINT M (-2 0 0.5),
	LINESTRING M (0 0 200, 0.1 -1 -20),
	POLYGON M ((0 0 7, 1 -1 -50, 2 0 0, 0 0 7)),
	MULTIPOINT(-1 5 -16, 0.23 7.0 0),
	MULTILINESTRING M ((0 -1 -2, 2 5 7)),
	MULTIPOLYGON M (((0 0 0, 1 1 1, 2 3 1, 0 0 0)))
)`,
			expectedErrStr: `syntax error: mixed dimensionality, parsed layout is XYM but encountered layout of not XYM at line 6, pos 11
LINE 6:  MULTIPOINT(-1 5 -16, 0.23 7.0 0),
                   ^
HINT: the M variant is required for non-empty XYM geometries in GEOMETRYCOLLECTIONs`,
		},
	}

	for _, tc := range errorTestCases {
		t.Run(tc.desc, func(t *testing.T) {
			_, err := Unmarshal(tc.input)
			assert.EqualError(t, err, tc.expectedErrStr)
		})
	}
}
