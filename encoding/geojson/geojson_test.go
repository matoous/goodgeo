package geojson

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/alecthomas/assert/v2"

	"github.com/matoous/goodgeo"
)

func TestGeometryDecode_NilCoordinates(t *testing.T) {
	for _, tc := range []struct {
		geometry Geometry
		want     goodgeo.T
	}{
		{
			geometry: Geometry{Type: "Point"},
			want:     goodgeo.NewPointEmpty(goodgeo.NoLayout),
		},
		{
			geometry: Geometry{Type: "LineString"},
			want:     goodgeo.NewLineString(goodgeo.NoLayout),
		},
		{
			geometry: Geometry{Type: "Polygon"},
			want:     goodgeo.NewPolygon(goodgeo.NoLayout),
		},
		{
			geometry: Geometry{Type: "MultiPoint"},
			want:     goodgeo.NewMultiPoint(goodgeo.NoLayout),
		},
		{
			geometry: Geometry{Type: "MultiLineString"},
			want:     goodgeo.NewMultiLineString(goodgeo.NoLayout),
		},
		{
			geometry: Geometry{Type: "MultiPolygon"},
			want:     goodgeo.NewMultiPolygon(goodgeo.NoLayout),
		},
		{
			geometry: Geometry{Type: "GeometryCollection"},
			want:     goodgeo.NewGeometryCollection(),
		},
	} {
		t.Run(tc.geometry.Type, func(t *testing.T) {
			got, err := tc.geometry.Decode()
			assert.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestGeometry(t *testing.T) {
	for _, tc := range []struct {
		g             goodgeo.T
		opts          []EncodeGeometryOption
		s             string
		skipUnmarshal bool
	}{
		{
			g: nil,
			s: `null`,
		},
		{
			g:    goodgeo.NewPointEmpty(goodgeo.XY),
			opts: []EncodeGeometryOption{EncodeGeometryWithMaxDecimalDigits(15)},
			s:    `{"type":"Point","coordinates":[]}`,
		},
		{
			g: goodgeo.NewPointEmpty(goodgeo.XY),
			s: `{"type":"Point","coordinates":[]}`,
		},
		{
			g: goodgeo.NewLineString(goodgeo.XY),
			s: `{"type":"LineString","coordinates":[]}`,
		},
		{
			g: goodgeo.NewPolygon(goodgeo.XY),
			s: `{"type":"Polygon","coordinates":[]}`,
		},
		{
			g: goodgeo.NewMultiPoint(goodgeo.XY),
			s: `{"type":"MultiPoint","coordinates":[]}`,
		},
		{
			g: goodgeo.NewMultiLineString(goodgeo.XY),
			s: `{"type":"MultiLineString","coordinates":[]}`,
		},
		{
			g: goodgeo.NewMultiPolygon(goodgeo.XY),
			s: `{"type":"MultiPolygon","coordinates":[]}`,
		},
		{
			g: goodgeo.NewGeometryCollection(),
			s: `{"type":"GeometryCollection","geometries":[]}`,
		},
		{
			g: nil,
			opts: []EncodeGeometryOption{
				EncodeGeometryWithBBox(),
				EncodeGeometryWithCRS(&CRS{
					Type: "name",
					Properties: map[string]interface{}{
						"name": "urn:ogc:def:crs:OGC:1.3:CRS84",
					},
				}),
			},
			s: `null`,
		},
		{
			g: goodgeo.NewPoint(DefaultLayout),
			opts: []EncodeGeometryOption{
				EncodeGeometryWithBBox(),
				EncodeGeometryWithCRS(&CRS{
					Type: "name",
					Properties: map[string]interface{}{
						"name": "urn:ogc:def:crs:OGC:1.3:CRS84",
					},
				}),
			},
			s: `{"type":"Point","bbox":[0,0,0,0],"crs":{"type":"name","properties":{"name":"urn:ogc:def:crs:OGC:1.3:CRS84"}},"coordinates":[0,0]}`,
		},
		{
			g: goodgeo.NewPoint(DefaultLayout),
			s: `{"type":"Point","coordinates":[0,0]}`,
		},
		{
			g: goodgeo.NewPoint(DefaultLayout),
			s: `{"type":"Point","coordinates":[0,0]}`,
		},
		{
			g: goodgeo.NewPoint(goodgeo.XY).MustSetCoords(goodgeo.Coord{1, 2}),
			s: `{"type":"Point","coordinates":[1,2]}`,
		},
		{
			g: goodgeo.NewPoint(goodgeo.XYZ).MustSetCoords(goodgeo.Coord{1, 2, 3}),
			s: `{"type":"Point","coordinates":[1,2,3]}`,
		},
		{
			g: goodgeo.NewPoint(goodgeo.XYZM).MustSetCoords(goodgeo.Coord{1, 2, 3, 4}),
			s: `{"type":"Point","coordinates":[1,2,3,4]}`,
		},
		{
			g:             goodgeo.NewPoint(goodgeo.XYZM).MustSetCoords(goodgeo.Coord{1.451, 2.89, 3.14, 4.03}),
			opts:          []EncodeGeometryOption{EncodeGeometryWithMaxDecimalDigits(1)},
			s:             `{"type":"Point","coordinates":[1.5,2.9,3.1,4]}`,
			skipUnmarshal: true,
		},
		{
			g: goodgeo.NewLineString(DefaultLayout),
			s: `{"type":"LineString","coordinates":[]}`,
		},
		{
			g:    goodgeo.NewLineString(DefaultLayout),
			opts: []EncodeGeometryOption{EncodeGeometryWithMaxDecimalDigits(1)},
			s:    `{"type":"LineString","coordinates":[]}`,
		},
		{
			g:             goodgeo.NewLineString(goodgeo.XY).MustSetCoords([]goodgeo.Coord{{1.1234, 2.5678}, {3.1234, 4.01234}}),
			opts:          []EncodeGeometryOption{EncodeGeometryWithMaxDecimalDigits(1)},
			s:             `{"type":"LineString","coordinates":[[1.1,2.6],[3.1,4]]}`,
			skipUnmarshal: true,
		},
		{
			g: goodgeo.NewLineString(goodgeo.XY).MustSetCoords([]goodgeo.Coord{{1, 2}, {3, 4}}),
			s: `{"type":"LineString","coordinates":[[1,2],[3,4]]}`,
		},
		{
			g: goodgeo.NewLineString(goodgeo.XYZ).MustSetCoords([]goodgeo.Coord{{1, 2, 3}, {4, 5, 6}}),
			s: `{"type":"LineString","coordinates":[[1,2,3],[4,5,6]]}`,
		},
		{
			g: goodgeo.NewLineString(goodgeo.XYZM).MustSetCoords([]goodgeo.Coord{{1, 2, 3, 4}, {5, 6, 7, 8}}),
			s: `{"type":"LineString","coordinates":[[1,2,3,4],[5,6,7,8]]}`,
		},
		{
			g: goodgeo.NewPolygon(DefaultLayout),
			s: `{"type":"Polygon","coordinates":[]}`,
		},
		{
			g: goodgeo.NewPolygon(goodgeo.XY).MustSetCoords([][]goodgeo.Coord{{{1, 2}, {3, 4}, {5, 6}, {1, 2}}}),
			s: `{"type":"Polygon","coordinates":[[[1,2],[3,4],[5,6],[1,2]]]}`,
		},
		{
			g: goodgeo.NewPolygon(goodgeo.XYZ).MustSetCoords([][]goodgeo.Coord{{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {1, 2, 3}}}),
			s: `{"type":"Polygon","coordinates":[[[1,2,3],[4,5,6],[7,8,9],[1,2,3]]]}`,
		},
		{
			g:             goodgeo.NewPolygon(goodgeo.XYZ).MustSetCoords([][]goodgeo.Coord{{{1.1, 2.2, 3.3}, {4.4, 5.5, 6.6}, {7.7, 8.8, 9.9}, {1.1, 2.2, 3.3}}}),
			opts:          []EncodeGeometryOption{EncodeGeometryWithMaxDecimalDigits(0)},
			s:             `{"type":"Polygon","coordinates":[[[1,2,3],[4,6,7],[8,9,10],[1,2,3]]]}`,
			skipUnmarshal: true,
		},
		{
			g: goodgeo.NewMultiPoint(DefaultLayout),
			s: `{"type":"MultiPoint","coordinates":[]}`,
		},
		{
			g: goodgeo.NewMultiPoint(goodgeo.XY).MustSetCoords([]goodgeo.Coord{{1, 2}, {3, 4}}),
			s: `{"type":"MultiPoint","coordinates":[[1,2],[3,4]]}`,
		},
		{
			g: goodgeo.NewMultiPoint(goodgeo.XY).MustSetCoords([]goodgeo.Coord{{1, 2}, nil, {3, 4}}),
			// In PostGIS, the empty point is not handled in GeoJSON (it emits invalid JSON).
			s: `{"type":"MultiPoint","coordinates":[[1,2],null,[3,4]]}`,
		},
		{
			g: goodgeo.NewMultiPoint(goodgeo.XYZ).MustSetCoords([]goodgeo.Coord{{1, 2, 3}, {4, 5, 6}}),
			s: `{"type":"MultiPoint","coordinates":[[1,2,3],[4,5,6]]}`,
		},
		{
			g: goodgeo.NewMultiPoint(goodgeo.XYZM).MustSetCoords([]goodgeo.Coord{{1, 2, 3, 4}, {5, 6, 7, 8}}),
			s: `{"type":"MultiPoint","coordinates":[[1,2,3,4],[5,6,7,8]]}`,
		},
		{
			g: goodgeo.NewMultiLineString(DefaultLayout),
			s: `{"type":"MultiLineString","coordinates":[]}`,
		},
		{
			g: goodgeo.NewMultiLineString(goodgeo.XY).MustSetCoords([][]goodgeo.Coord{{{1, 2}, {3, 4}, {5, 6}, {1, 2}}}),
			s: `{"type":"MultiLineString","coordinates":[[[1,2],[3,4],[5,6],[1,2]]]}`,
		},
		{
			g: goodgeo.NewMultiLineString(goodgeo.XYZ).MustSetCoords([][]goodgeo.Coord{{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {1, 2, 3}}}),
			s: `{"type":"MultiLineString","coordinates":[[[1,2,3],[4,5,6],[7,8,9],[1,2,3]]]}`,
		},
		{
			g: goodgeo.NewMultiPolygon(DefaultLayout),
			s: `{"type":"MultiPolygon","coordinates":[]}`,
		},
		{
			g: goodgeo.NewMultiPolygon(goodgeo.XYZ).MustSetCoords([][][]goodgeo.Coord{{{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {1, 2, 3}}, {{-1, -2, -3}, {-4, -5, -6}, {-7, -8, -9}, {-1, -2, -3}}}}),
			s: `{"type":"MultiPolygon","coordinates":[[[[1,2,3],[4,5,6],[7,8,9],[1,2,3]],[[-1,-2,-3],[-4,-5,-6],[-7,-8,-9],[-1,-2,-3]]]]}`,
		},
		{
			g: goodgeo.NewGeometryCollection().MustPush(
				goodgeo.NewPoint(goodgeo.XY).MustSetCoords(goodgeo.Coord{100, 0}),
				goodgeo.NewLineString(goodgeo.XY).MustSetCoords([]goodgeo.Coord{{101, 0}, {102, 1}}),
			),
			s: `{"type":"GeometryCollection","geometries":[{"type":"Point","coordinates":[100,0]},{"type":"LineString","coordinates":[[101,0],[102,1]]}]}`,
		},
		{
			g: goodgeo.NewGeometryCollection().MustPush(
				goodgeo.NewPoint(goodgeo.XY).MustSetCoords(goodgeo.Coord{100.123, 0.456}),
				goodgeo.NewLineString(goodgeo.XY).MustSetCoords([]goodgeo.Coord{{101.569, 0.898}, {102.123, 1.567}}),
			),
			opts:          []EncodeGeometryOption{EncodeGeometryWithMaxDecimalDigits(2)},
			s:             `{"type":"GeometryCollection","geometries":[{"type":"Point","coordinates":[100.12,0.46]},{"type":"LineString","coordinates":[[101.57,0.9],[102.12,1.57]]}]}`,
			skipUnmarshal: true,
		},
		{
			g: goodgeo.NewGeometryCollection().MustPush(
				goodgeo.NewPoint(goodgeo.XY).MustSetCoords(goodgeo.Coord{100.123, 0.456}),
				goodgeo.NewLineString(goodgeo.XY).MustSetCoords([]goodgeo.Coord{{101.569, 0.898}, {102.123, 1.567}}),
			),
			opts:          []EncodeGeometryOption{EncodeGeometryWithMaxDecimalDigits(2), EncodeGeometryWithBBox()},
			s:             `{"type":"GeometryCollection","bbox":[100.12,0.46,102.12,1.57],"geometries":[{"type":"Point","coordinates":[100.12,0.46]},{"type":"LineString","coordinates":[[101.57,0.9],[102.12,1.57]]}]}`,
			skipUnmarshal: true,
		},
	} {
		t.Run(tc.s, func(t *testing.T) {
			got, err := Marshal(tc.g, tc.opts...)
			assert.NoError(t, err)
			assert.Equal(t, tc.s, string(got))

			if !tc.skipUnmarshal {
				var g goodgeo.T
				assert.NoError(t, Unmarshal([]byte(tc.s), &g))
				assert.Equal(t, tc.g, g)
			}
		})
	}
}

func TestFeature(t *testing.T) {
	for _, tc := range []struct {
		skipMarshalTest bool
		useNumber       bool
		f               *Feature
		s               string
	}{
		{
			skipMarshalTest: true,
			f: &Feature{
				ID: "10",
			},
			s: `{"type":"Feature","id":10,"geometry":null,"properties":null}`,
		},
		{
			skipMarshalTest: true,
			useNumber:       true,
			f: &Feature{
				ID: "10",
			},
			s: `{"type":"Feature","id":10.0,"geometry":null,"properties":null}`,
		},
		{
			f: &Feature{},
			s: `{"type":"Feature","geometry":null,"properties":null}`,
		},
		{
			f: &Feature{
				Geometry: goodgeo.NewPoint(goodgeo.XY).MustSetCoords([]float64{125.6, 10.1}),
				Properties: map[string]interface{}{
					"name": "Dinagat Islands",
				},
			},
			s: `{"type":"Feature","geometry":{"type":"Point","coordinates":[125.6,10.1]},"properties":{"name":"Dinagat Islands"}}`,
		},
		{
			f: &Feature{
				Geometry: goodgeo.NewLineString(goodgeo.XY).MustSetCoords([]goodgeo.Coord{{102, 0}, {103, 1}, {104, 0}, {105, 1}}),
				Properties: map[string]interface{}{
					"prop0": "value0",
					"prop1": 0.0,
				},
			},
			s: `{"type":"Feature","geometry":{"type":"LineString","coordinates":[[102,0],[103,1],[104,0],[105,1]]},"properties":{"prop0":"value0","prop1":0}}`,
		},
		{
			f: &Feature{
				BBox:     goodgeo.NewBounds(goodgeo.XY).Set(100, 0, 101, 1),
				Geometry: goodgeo.NewPolygon(goodgeo.XY).MustSetCoords([][]goodgeo.Coord{{{100, 0}, {101, 0}, {101, 1}, {100, 1}, {100, 0}}}),
				Properties: map[string]interface{}{
					"prop0": "value0",
					"prop1": map[string]interface{}{
						"this": "that",
					},
				},
			},
			s: `{"type":"Feature","bbox":[100,0,101,1],"geometry":{"type":"Polygon","coordinates":[[[100,0],[101,0],[101,1],[100,1],[100,0]]]},"properties":{"prop0":"value0","prop1":{"this":"that"}}}`,
		},
		{
			f: &Feature{
				ID:       "0",
				Geometry: goodgeo.NewPoint(goodgeo.XY).MustSetCoords(goodgeo.Coord{1, 2}),
			},
			s: `{"type":"Feature","id":"0","geometry":{"type":"Point","coordinates":[1,2]},"properties":null}`,
		},
		{
			f: &Feature{
				ID:       "f",
				Geometry: goodgeo.NewPoint(goodgeo.XY).MustSetCoords(goodgeo.Coord{1, 2}),
			},
			s: `{"type":"Feature","id":"f","geometry":{"type":"Point","coordinates":[1,2]},"properties":null}`,
		},
	} {
		t.Run(tc.s, func(t *testing.T) {
			if !tc.skipMarshalTest {
				t.Run("marshal", func(t *testing.T) {
					got, err := json.Marshal(tc.f)
					assert.NoError(t, err)
					assert.Equal(t, tc.s, string(got))
				})
			}

			t.Run("unmarshal", func(t *testing.T) {
				f := &Feature{}
				decoder := json.NewDecoder(strings.NewReader(tc.s))
				if tc.useNumber {
					decoder.UseNumber()
				}
				assert.NoError(t, decoder.Decode(f))
				assert.Equal(t, tc.f, f)
			})
		})
	}
}

func TestFeatureCollection(t *testing.T) {
	for _, tc := range []struct {
		fc *FeatureCollection
		s  string
	}{
		{
			fc: &FeatureCollection{
				Features: []*Feature{
					{
						Geometry: goodgeo.NewPoint(goodgeo.XY).MustSetCoords([]float64{125.6, 10.1}),
						Properties: map[string]interface{}{
							"name": "Dinagat Islands",
						},
					},
				},
			},
			s: `{"type":"FeatureCollection","features":[{"type":"Feature","geometry":{"type":"Point","coordinates":[125.6,10.1]},"properties":{"name":"Dinagat Islands"}}]}`,
		},
		{
			fc: &FeatureCollection{
				BBox: goodgeo.NewBounds(goodgeo.XY).Set(100, 0, 125.6, 10.1),
				Features: []*Feature{
					{
						Geometry: goodgeo.NewPoint(goodgeo.XY).MustSetCoords([]float64{125.6, 10.1}),
						Properties: map[string]interface{}{
							"name": "Dinagat Islands",
						},
					},
					{
						Geometry: goodgeo.NewLineString(goodgeo.XY).MustSetCoords([]goodgeo.Coord{{102, 0}, {103, 1}, {104, 0}, {105, 1}}),
						Properties: map[string]interface{}{
							"prop0": "value0",
							"prop1": 0.0,
						},
					},
					{
						Geometry: goodgeo.NewPolygon(goodgeo.XY).MustSetCoords([][]goodgeo.Coord{{{100, 0}, {101, 0}, {101, 1}, {100, 1}, {100, 0}}}),
						Properties: map[string]interface{}{
							"prop0": "value0",
							"prop1": map[string]interface{}{
								"this": "that",
							},
						},
					},
				},
			},
			s: `{"type":"FeatureCollection","bbox":[100,0,125.6,10.1],"features":[{"type":"Feature","geometry":{"type":"Point","coordinates":[125.6,10.1]},"properties":{"name":"Dinagat Islands"}},{"type":"Feature","geometry":{"type":"LineString","coordinates":[[102,0],[103,1],[104,0],[105,1]]},"properties":{"prop0":"value0","prop1":0}},{"type":"Feature","geometry":{"type":"Polygon","coordinates":[[[100,0],[101,0],[101,1],[100,1],[100,0]]]},"properties":{"prop0":"value0","prop1":{"this":"that"}}}]}`,
		},
	} {
		t.Run(tc.s, func(t *testing.T) {
			t.Run("marshal", func(t *testing.T) {
				got, err := json.Marshal(tc.fc)
				assert.NoError(t, err)
				assert.Equal(t, tc.s, string(got))
			})

			t.Run("unmarshal", func(t *testing.T) {
				fc := &FeatureCollection{}
				assert.NoError(t, json.Unmarshal([]byte(tc.s), fc))
				assert.Equal(t, tc.fc, fc)
			})
		})
	}
}
