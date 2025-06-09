package ewkb

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert/v2"

	"github.com/matoous/goodgeo"
	"github.com/matoous/goodgeo/internal/goodgeotest"
)

func test(t *testing.T, g goodgeo.T, xdr, ndr []byte) {
	t.Helper()
	if xdr != nil {
		t.Run("xdr", func(t *testing.T) {
			t.Run("unmarshal", func(t *testing.T) {
				got, err := Unmarshal(xdr)
				assert.NoError(t, err)
				assert.Equal(t, g, got)
			})

			t.Run("marshal", func(t *testing.T) {
				got, err := Marshal(g, XDR)
				assert.NoError(t, err)
				assert.Equal(t, xdr, got)
			})
		})
	}
	if ndr != nil {
		t.Run("ndr", func(t *testing.T) {
			t.Run("unmarshal", func(t *testing.T) {
				got, err := Unmarshal(ndr)
				assert.NoError(t, err)
				assert.Equal(t, g, got)
			})

			t.Run("marshal", func(t *testing.T) {
				got, err := Marshal(g, NDR)
				assert.NoError(t, err)
				assert.Equal(t, ndr, got)
			})
		})
	}
	t.Run("scan", func(t *testing.T) {
		switch g := g.(type) {
		case *goodgeo.Point:
			var p Point
			if xdr != nil {
				t.Run("xdr", func(t *testing.T) {
					assert.NoError(t, p.Scan(xdr))
					assert.Equal(t, Point{g}, p)
				})
			}
			if ndr != nil {
				t.Run("ndr", func(t *testing.T) {
					assert.NoError(t, p.Scan(ndr))
					assert.Equal(t, Point{g}, p)
				})
			}
		case *goodgeo.LineString:
			var ls LineString
			if xdr != nil {
				t.Run("xdr", func(t *testing.T) {
					assert.NoError(t, ls.Scan(xdr))
					assert.Equal(t, LineString{g}, ls)
				})
			}
			if ndr != nil {
				t.Run("ndr", func(t *testing.T) {
					assert.NoError(t, ls.Scan(ndr))
					assert.Equal(t, LineString{g}, ls)
				})
			}
		case *goodgeo.Polygon:
			var p Polygon
			if xdr != nil {
				t.Run("xdr", func(t *testing.T) {
					assert.NoError(t, p.Scan(xdr))
					assert.Equal(t, Polygon{g}, p)
				})
			}
			if ndr != nil {
				t.Run("ndr", func(t *testing.T) {
					assert.NoError(t, p.Scan(ndr))
					assert.Equal(t, Polygon{g}, p)
				})
			}
		case *goodgeo.MultiPoint:
			var mp MultiPoint
			if xdr != nil {
				t.Run("xdr", func(t *testing.T) {
					assert.NoError(t, mp.Scan(xdr))
					assert.Equal(t, MultiPoint{g}, mp)
				})
			}
			if ndr != nil {
				t.Run("ndr", func(t *testing.T) {
					assert.NoError(t, mp.Scan(ndr))
					assert.Equal(t, MultiPoint{g}, mp)
				})
			}
		case *goodgeo.MultiLineString:
			var mls MultiLineString
			if xdr != nil {
				t.Run("xdr", func(t *testing.T) {
					assert.NoError(t, mls.Scan(xdr))
					assert.Equal(t, MultiLineString{g}, mls)
				})
			}
			if ndr != nil {
				t.Run("ndr", func(t *testing.T) {
					assert.NoError(t, mls.Scan(ndr))
					assert.Equal(t, MultiLineString{g}, mls)
				})
			}
		case *goodgeo.MultiPolygon:
			var mp MultiPolygon
			if xdr != nil {
				t.Run("xdr", func(t *testing.T) {
					assert.NoError(t, mp.Scan(xdr))
					assert.Equal(t, MultiPolygon{g}, mp)
				})
			}
			if ndr != nil {
				t.Run("ndr", func(t *testing.T) {
					assert.NoError(t, mp.Scan(ndr))
					assert.Equal(t, MultiPolygon{g}, mp)
				})
			}
		case *goodgeo.GeometryCollection:
			var gc GeometryCollection
			if xdr != nil {
				t.Run("xdr", func(t *testing.T) {
					assert.NoError(t, gc.Scan(xdr))
					assert.Equal(t, GeometryCollection{g}, gc)
				})
			}
			if ndr != nil {
				t.Run("ndr", func(t *testing.T) {
					assert.NoError(t, gc.Scan(ndr))
					assert.Equal(t, GeometryCollection{g}, gc)
				})
			}
		}
	})
}

func Test(t *testing.T) {
	for _, tc := range []struct {
		g   goodgeo.T
		xdr []byte
		ndr []byte
	}{
		{
			g:   goodgeo.NewPointEmpty(goodgeo.XY),
			xdr: goodgeotest.MustHexDecode("00000000017ff80000000000007ff8000000000000"),
			ndr: goodgeotest.MustHexDecode("0101000000000000000000f87f000000000000f87f"),
		},
		{
			g:   goodgeo.NewPointEmpty(goodgeo.XYM),
			xdr: goodgeotest.MustHexDecode("00400000017ff80000000000007ff80000000000007ff8000000000000"),
			ndr: goodgeotest.MustHexDecode("0101000040000000000000f87f000000000000f87f000000000000f87f"),
		},
		{
			g:   goodgeo.NewPointEmpty(goodgeo.XYZ),
			xdr: goodgeotest.MustHexDecode("00800000017ff80000000000007ff80000000000007ff8000000000000"),
			ndr: goodgeotest.MustHexDecode("0101000080000000000000f87f000000000000f87f000000000000f87f"),
		},
		{
			g:   goodgeo.NewPointEmpty(goodgeo.XYZM),
			xdr: goodgeotest.MustHexDecode("00c00000017ff80000000000007ff80000000000007ff80000000000007ff8000000000000"),
			ndr: goodgeotest.MustHexDecode("01010000c0000000000000f87f000000000000f87f000000000000f87f000000000000f87f"),
		},
		{
			g:   goodgeo.NewGeometryCollection().MustPush(goodgeo.NewPointEmpty(goodgeo.XY)),
			xdr: goodgeotest.MustHexDecode("00000000070000000100000000017ff80000000000007ff8000000000000"),
			ndr: goodgeotest.MustHexDecode("0107000000010000000101000000000000000000f87f000000000000f87f"),
		},
		{
			g:   goodgeo.NewPointEmpty(goodgeo.XY).SetSRID(4326),
			xdr: goodgeotest.MustHexDecode("0020000001000010e67ff80000000000007ff8000000000000"),
			ndr: goodgeotest.MustHexDecode("0101000020e6100000000000000000f87f000000000000f87f"),
		},
		{
			g:   goodgeo.NewPoint(goodgeo.XY).MustSetCoords(goodgeo.Coord{1, 2}),
			xdr: goodgeotest.MustHexDecode("00000000013ff00000000000004000000000000000"),
			ndr: goodgeotest.MustHexDecode("0101000000000000000000f03f0000000000000040"),
		},
		{
			g:   goodgeo.NewPoint(goodgeo.XYZ).MustSetCoords(goodgeo.Coord{1, 2, 3}),
			xdr: goodgeotest.MustHexDecode("00800000013ff000000000000040000000000000004008000000000000"),
			ndr: goodgeotest.MustHexDecode("0101000080000000000000f03f00000000000000400000000000000840"),
		},
		{
			g:   goodgeo.NewPoint(goodgeo.XYM).MustSetCoords(goodgeo.Coord{1, 2, 3}),
			xdr: goodgeotest.MustHexDecode("00400000013ff000000000000040000000000000004008000000000000"),
			ndr: goodgeotest.MustHexDecode("0101000040000000000000f03f00000000000000400000000000000840"),
		},
		{
			g:   goodgeo.NewPoint(goodgeo.XYZM).MustSetCoords(goodgeo.Coord{1, 2, 3, 4}),
			xdr: goodgeotest.MustHexDecode("00c00000013ff0000000000000400000000000000040080000000000004010000000000000"),
			ndr: goodgeotest.MustHexDecode("01010000c0000000000000f03f000000000000004000000000000008400000000000001040"),
		},
		{
			g:   goodgeo.NewPoint(goodgeo.XY).SetSRID(4326).MustSetCoords(goodgeo.Coord{1, 2}),
			xdr: goodgeotest.MustHexDecode("0020000001000010e63ff00000000000004000000000000000"),
			ndr: goodgeotest.MustHexDecode("0101000020e6100000000000000000f03f0000000000000040"),
		},
		{
			g:   goodgeo.NewPoint(goodgeo.XYZ).SetSRID(4326).MustSetCoords(goodgeo.Coord{1, 2, 3}),
			xdr: goodgeotest.MustHexDecode("00a0000001000010e63ff000000000000040000000000000004008000000000000"),
			ndr: goodgeotest.MustHexDecode("01010000a0e6100000000000000000f03f00000000000000400000000000000840"),
		},
		{
			g:   goodgeo.NewPoint(goodgeo.XYM).SetSRID(4326).MustSetCoords(goodgeo.Coord{1, 2, 3}),
			xdr: goodgeotest.MustHexDecode("0060000001000010e63ff000000000000040000000000000004008000000000000"),
			ndr: goodgeotest.MustHexDecode("0101000060e6100000000000000000f03f00000000000000400000000000000840"),
		},
		{
			g:   goodgeo.NewPoint(goodgeo.XYZM).SetSRID(4326).MustSetCoords(goodgeo.Coord{1, 2, 3, 4}),
			xdr: goodgeotest.MustHexDecode("00e0000001000010e63ff0000000000000400000000000000040080000000000004010000000000000"),
			ndr: goodgeotest.MustHexDecode("01010000e0e6100000000000000000f03f000000000000004000000000000008400000000000001040"),
		},
		{
			g:   goodgeo.NewMultiPoint(goodgeo.XY).SetSRID(4326).MustSetCoords([]goodgeo.Coord{{1, 2}, nil, {3, 4}}),
			xdr: goodgeotest.MustHexDecode("0020000004000010e60000000300000000013ff0000000000000400000000000000000000000017ff80000000000007ff8000000000000000000000140080000000000004010000000000000"),
			ndr: goodgeotest.MustHexDecode("0104000020e6100000030000000101000000000000000000f03f00000000000000400101000000000000000000f87f000000000000f87f010100000000000000000008400000000000001040"),
		},
		{
			g:   goodgeo.NewGeometryCollection().SetSRID(4326).MustSetLayout(goodgeo.XY),
			xdr: goodgeotest.MustHexDecode("0020000007000010e600000000"),
			ndr: goodgeotest.MustHexDecode("0107000020e610000000000000"),
		},
		{
			g:   goodgeo.NewGeometryCollection().SetSRID(4326).MustSetLayout(goodgeo.XYZ),
			ndr: goodgeotest.MustHexDecode("01070000a0e610000000000000"),
			xdr: goodgeotest.MustHexDecode("00a0000007000010e600000000"),
		},
		{
			g:   goodgeo.NewGeometryCollection().SetSRID(4326).MustSetLayout(goodgeo.XYM),
			ndr: goodgeotest.MustHexDecode("0107000060e610000000000000"),
			xdr: goodgeotest.MustHexDecode("0060000007000010e600000000"),
		},
		{
			g:   goodgeo.NewGeometryCollection().SetSRID(4326).MustSetLayout(goodgeo.XYZM),
			ndr: goodgeotest.MustHexDecode("01070000e0e610000000000000"),
			xdr: goodgeotest.MustHexDecode("00e0000007000010e600000000"),
		},
		{
			g: goodgeo.NewGeometryCollection().SetSRID(4326).MustPush(
				goodgeo.NewPoint(goodgeo.XY).MustSetCoords(goodgeo.Coord{1, 2}),
				goodgeo.NewLineString(goodgeo.XY).MustSetCoords([]goodgeo.Coord{{3, 4}, {5, 6}}),
			),
			ndr: goodgeotest.MustHexDecode("0107000020E6100000020000000101000000000000000000F03F00000000000000400102000000020000000000000000000840000000000000104000000000000014400000000000001840"),
			xdr: goodgeotest.MustHexDecode("0020000007000010e60000000200000000013ff000000000000040000000000000000000000002000000024008000000000000401000000000000040140000000000004018000000000000"),
		},
	} {
		t.Run(fmt.Sprintf("ndr:%s", tc.ndr), func(t *testing.T) {
			test(t, tc.g, tc.xdr, tc.ndr)
		})
	}
}
