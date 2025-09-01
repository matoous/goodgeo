package wkb

import (
	"encoding/binary"
	"fmt"
	"testing"

	"github.com/alecthomas/assert/v2"

	"github.com/matoous/goodgeo"
	"github.com/matoous/goodgeo/encoding/wkbcommon"
	"github.com/matoous/goodgeo/internal/goodgeotest"
	"github.com/matoous/goodgeo/internal/testdata"
)

func test(t *testing.T, g goodgeo.T, xdr, ndr []byte, opts ...wkbcommon.WKBOption) {
	t.Helper()
	if xdr != nil {
		t.Run("xdr", func(t *testing.T) {
			t.Run("unmarshal", func(t *testing.T) {
				got, err := Unmarshal(xdr, opts...)
				assert.NoError(t, err)
				assert.Equal(t, g, got)
			})

			t.Run("marshal", func(t *testing.T) {
				got, err := Marshal(g, XDR, opts...)
				assert.NoError(t, err)
				assert.Equal(t, xdr, got)
			})
		})
	}
	if ndr != nil {
		t.Run("ndr", func(t *testing.T) {
			t.Run("unmarshal", func(t *testing.T) {
				got, err := Unmarshal(ndr, opts...)
				assert.NoError(t, err)
				assert.Equal(t, g, got)
			})

			t.Run("marshal", func(t *testing.T) {
				got, err := Marshal(g, NDR, opts...)
				assert.NoError(t, err)
				assert.Equal(t, ndr, got)
			})
		})
	}
	t.Run("scan", func(t *testing.T) {
		switch g := g.(type) {
		case *goodgeo.Point:
			p := Point{opts: opts}
			if xdr != nil {
				t.Run("xdr", func(t *testing.T) {
					assert.NoError(t, p.Scan(xdr))
					assert.Equal(t, Point{g, opts}, p)
				})
			}
			if ndr != nil {
				t.Run("ndr", func(t *testing.T) {
					assert.NoError(t, p.Scan(ndr))
					assert.Equal(t, Point{g, opts}, p)
				})
			}
		case *goodgeo.LineString:
			ls := LineString{opts: opts}
			if xdr != nil {
				t.Run("xdr", func(t *testing.T) {
					assert.NoError(t, ls.Scan(xdr))
					assert.Equal(t, LineString{g, opts}, ls)
				})
			}
			if ndr != nil {
				t.Run("ndr", func(t *testing.T) {
					assert.NoError(t, ls.Scan(ndr))
					assert.Equal(t, LineString{g, opts}, ls)
				})
			}
		case *goodgeo.Polygon:
			p := Polygon{opts: opts}
			if xdr != nil {
				t.Run("xdr", func(t *testing.T) {
					assert.NoError(t, p.Scan(xdr))
					assert.Equal(t, Polygon{g, opts}, p)
				})
			}
			if ndr != nil {
				t.Run("ndr", func(t *testing.T) {
					assert.NoError(t, p.Scan(ndr))
					assert.Equal(t, Polygon{g, opts}, p)
				})
			}
		case *goodgeo.MultiPoint:
			mp := MultiPoint{opts: opts}
			if xdr != nil {
				t.Run("xdr", func(t *testing.T) {
					assert.NoError(t, mp.Scan(xdr))
					assert.Equal(t, MultiPoint{g, opts}, mp)
				})
			}
			if ndr != nil {
				t.Run("ndr", func(t *testing.T) {
					assert.NoError(t, mp.Scan(ndr))
					assert.Equal(t, MultiPoint{g, opts}, mp)
				})
			}
		case *goodgeo.MultiLineString:
			mls := MultiLineString{opts: opts}
			if xdr != nil {
				t.Run("xdr", func(t *testing.T) {
					assert.NoError(t, mls.Scan(xdr))
					assert.Equal(t, MultiLineString{g, opts}, mls)
				})
			}
			if ndr != nil {
				t.Run("ndr", func(t *testing.T) {
					assert.NoError(t, mls.Scan(ndr))
					assert.Equal(t, MultiLineString{g, opts}, mls)
				})
			}
		case *goodgeo.MultiPolygon:
			mp := MultiPolygon{opts: opts}
			if xdr != nil {
				t.Run("xdr", func(t *testing.T) {
					assert.NoError(t, mp.Scan(xdr))
					assert.Equal(t, MultiPolygon{g, opts}, mp)
				})
			}
			if ndr != nil {
				t.Run("ndr", func(t *testing.T) {
					assert.NoError(t, mp.Scan(ndr))
					assert.Equal(t, MultiPolygon{g, opts}, mp)
				})
			}
		case *goodgeo.GeometryCollection:
			gc := GeometryCollection{opts: opts}
			if xdr != nil {
				t.Run("xdr", func(t *testing.T) {
					assert.NoError(t, gc.Scan(xdr))
					assert.Equal(t, GeometryCollection{g, opts}, gc)
				})
			}
			if ndr != nil {
				t.Run("ndr", func(t *testing.T) {
					assert.NoError(t, gc.Scan(ndr))
					assert.Equal(t, GeometryCollection{g, opts}, gc)
				})
			}
		}
	})
}

func Test(t *testing.T) {
	for _, tc := range []struct {
		g    goodgeo.T
		opts []wkbcommon.WKBOption
		xdr  []byte
		ndr  []byte
	}{
		{
			g:    goodgeo.NewPointEmpty(goodgeo.XY),
			opts: []wkbcommon.WKBOption{wkbcommon.WKBOptionEmptyPointHandling(wkbcommon.EmptyPointHandlingNaN)},
			xdr:  goodgeotest.MustHexDecode("00000000017ff80000000000007ff8000000000000"),
			ndr:  goodgeotest.MustHexDecode("0101000000000000000000f87f000000000000f87f"),
		},
		{
			g:    goodgeo.NewPointEmpty(goodgeo.XYM),
			opts: []wkbcommon.WKBOption{wkbcommon.WKBOptionEmptyPointHandling(wkbcommon.EmptyPointHandlingNaN)},
			xdr:  goodgeotest.MustHexDecode("00000007d17ff80000000000007ff80000000000007ff8000000000000"),
			ndr:  goodgeotest.MustHexDecode("01d1070000000000000000f87f000000000000f87f000000000000f87f"),
		},
		{
			g:    goodgeo.NewPointEmpty(goodgeo.XYZ),
			opts: []wkbcommon.WKBOption{wkbcommon.WKBOptionEmptyPointHandling(wkbcommon.EmptyPointHandlingNaN)},
			xdr:  goodgeotest.MustHexDecode("00000003e97ff80000000000007ff80000000000007ff8000000000000"),
			ndr:  goodgeotest.MustHexDecode("01e9030000000000000000f87f000000000000f87f000000000000f87f"),
		},
		{
			g:    goodgeo.NewPointEmpty(goodgeo.XYZM),
			opts: []wkbcommon.WKBOption{wkbcommon.WKBOptionEmptyPointHandling(wkbcommon.EmptyPointHandlingNaN)},
			xdr:  goodgeotest.MustHexDecode("0000000bb97ff80000000000007ff80000000000007ff80000000000007ff8000000000000"),
			ndr:  goodgeotest.MustHexDecode("01b90b0000000000000000f87f000000000000f87f000000000000f87f000000000000f87f"),
		},
		{
			g:    goodgeo.NewGeometryCollection().MustPush(goodgeo.NewPointEmpty(goodgeo.XY)),
			opts: []wkbcommon.WKBOption{wkbcommon.WKBOptionEmptyPointHandling(wkbcommon.EmptyPointHandlingNaN)},
			xdr:  goodgeotest.MustHexDecode("00000000070000000100000000017ff80000000000007ff8000000000000"),
			ndr:  goodgeotest.MustHexDecode("0107000000010000000101000000000000000000f87f000000000000f87f"),
		},
		{
			g:   goodgeo.NewPoint(goodgeo.XY).MustSetCoords(goodgeo.Coord{1, 2}),
			xdr: goodgeotest.MustHexDecode("00000000013ff00000000000004000000000000000"),
			ndr: goodgeotest.MustHexDecode("0101000000000000000000f03f0000000000000040"),
		},
		{
			g:   goodgeo.NewPoint(goodgeo.XYZ).MustSetCoords(goodgeo.Coord{1, 2, 3}),
			xdr: goodgeotest.MustHexDecode("00000003e93ff000000000000040000000000000004008000000000000"),
			ndr: goodgeotest.MustHexDecode("01e9030000000000000000f03f00000000000000400000000000000840"),
		},
		{
			g:   goodgeo.NewPoint(goodgeo.XYM).MustSetCoords(goodgeo.Coord{1, 2, 3}),
			xdr: goodgeotest.MustHexDecode("00000007d13ff000000000000040000000000000004008000000000000"),
			ndr: goodgeotest.MustHexDecode("01d1070000000000000000f03f00000000000000400000000000000840"),
		},
		{
			g:   goodgeo.NewPoint(goodgeo.XYZM).MustSetCoords(goodgeo.Coord{1, 2, 3, 4}),
			xdr: goodgeotest.MustHexDecode("0000000bb93ff0000000000000400000000000000040080000000000004010000000000000"),
			ndr: goodgeotest.MustHexDecode("01b90b0000000000000000f03f000000000000004000000000000008400000000000001040"),
		},
		{
			g:   goodgeo.NewLineString(goodgeo.XY).MustSetCoords([]goodgeo.Coord{{1, 2}, {3, 4}}),
			xdr: goodgeotest.MustHexDecode("0000000002000000023ff0000000000000400000000000000040080000000000004010000000000000"),
			ndr: goodgeotest.MustHexDecode("010200000002000000000000000000f03f000000000000004000000000000008400000000000001040"),
		},
		{
			g:   goodgeo.NewLineString(goodgeo.XYZ).MustSetCoords([]goodgeo.Coord{{1, 2, 3}, {4, 5, 6}}),
			xdr: goodgeotest.MustHexDecode("00000003ea000000023ff000000000000040000000000000004008000000000000401000000000000040140000000000004018000000000000"),
			ndr: goodgeotest.MustHexDecode("01ea03000002000000000000000000f03f00000000000000400000000000000840000000000000104000000000000014400000000000001840"),
		},
		{
			g:   goodgeo.NewLineString(goodgeo.XYM).MustSetCoords([]goodgeo.Coord{{1, 2, 3}, {4, 5, 6}}),
			xdr: goodgeotest.MustHexDecode("00000007d2000000023ff000000000000040000000000000004008000000000000401000000000000040140000000000004018000000000000"),
			ndr: goodgeotest.MustHexDecode("01d207000002000000000000000000f03f00000000000000400000000000000840000000000000104000000000000014400000000000001840"),
		},
		{
			g:   goodgeo.NewLineString(goodgeo.XYZM).MustSetCoords([]goodgeo.Coord{{1, 2, 3, 4}, {5, 6, 7, 8}}),
			xdr: goodgeotest.MustHexDecode("0000000bba000000023ff000000000000040000000000000004008000000000000401000000000000040140000000000004018000000000000401c0000000000004020000000000000"),
			ndr: goodgeotest.MustHexDecode("01ba0b000002000000000000000000f03f000000000000004000000000000008400000000000001040000000000000144000000000000018400000000000001c400000000000002040"),
		},
		{
			g:   goodgeo.NewPolygon(goodgeo.XY).MustSetCoords([][]goodgeo.Coord{{{1, 2}, {3, 4}, {5, 6}, {1, 2}}}),
			xdr: goodgeotest.MustHexDecode("000000000300000001000000043ff0000000000000400000000000000040080000000000004010000000000000401400000000000040180000000000003ff00000000000004000000000000000"),
			ndr: goodgeotest.MustHexDecode("01030000000100000004000000000000000000f03f00000000000000400000000000000840000000000000104000000000000014400000000000001840000000000000f03f0000000000000040"),
		},
		{
			g:   goodgeo.NewPolygon(goodgeo.XYZ).MustSetCoords([][]goodgeo.Coord{{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {1, 2, 3}}}),
			xdr: goodgeotest.MustHexDecode("00000003eb00000001000000043ff000000000000040000000000000004008000000000000401000000000000040140000000000004018000000000000401c000000000000402000000000000040220000000000003ff000000000000040000000000000004008000000000000"),
			ndr: goodgeotest.MustHexDecode("01eb0300000100000004000000000000000000f03f000000000000004000000000000008400000000000001040000000000000144000000000000018400000000000001c4000000000000020400000000000002240000000000000f03f00000000000000400000000000000840"),
		},
		{
			g:   goodgeo.NewPolygon(goodgeo.XYM).MustSetCoords([][]goodgeo.Coord{{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}, {1, 2, 3}}}),
			xdr: goodgeotest.MustHexDecode("00000007d300000001000000043ff000000000000040000000000000004008000000000000401000000000000040140000000000004018000000000000401c000000000000402000000000000040220000000000003ff000000000000040000000000000004008000000000000"),
			ndr: goodgeotest.MustHexDecode("01d30700000100000004000000000000000000f03f000000000000004000000000000008400000000000001040000000000000144000000000000018400000000000001c4000000000000020400000000000002240000000000000f03f00000000000000400000000000000840"),
		},
		{
			g:   goodgeo.NewPolygon(goodgeo.XYZM).MustSetCoords([][]goodgeo.Coord{{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {1, 2, 3, 4}}}),
			xdr: goodgeotest.MustHexDecode("0000000bbb00000001000000043ff000000000000040000000000000004008000000000000401000000000000040140000000000004018000000000000401c000000000000402000000000000040220000000000004024000000000000402600000000000040280000000000003ff0000000000000400000000000000040080000000000004010000000000000"),
			ndr: goodgeotest.MustHexDecode("01bb0b00000100000004000000000000000000f03f000000000000004000000000000008400000000000001040000000000000144000000000000018400000000000001c4000000000000020400000000000002240000000000000244000000000000026400000000000002840000000000000f03f000000000000004000000000000008400000000000001040"),
		},
		{
			g:   goodgeo.NewMultiPoint(goodgeo.XY).MustSetCoords([]goodgeo.Coord{{1, 2}, {3, 4}}),
			xdr: goodgeotest.MustHexDecode("00000000040000000200000000013ff00000000000004000000000000000000000000140080000000000004010000000000000"),
			ndr: goodgeotest.MustHexDecode("0104000000020000000101000000000000000000f03f0000000000000040010100000000000000000008400000000000001040"),
		},
		{
			g:   goodgeo.NewMultiPoint(goodgeo.XYZ).MustSetCoords([]goodgeo.Coord{{1, 2, 3}, {4, 5, 6}}),
			xdr: goodgeotest.MustHexDecode("00000003ec0000000200000003e93ff00000000000004000000000000000400800000000000000000003e9401000000000000040140000000000004018000000000000"),
			ndr: goodgeotest.MustHexDecode("01ec0300000200000001e9030000000000000000f03f0000000000000040000000000000084001e9030000000000000000104000000000000014400000000000001840"),
		},
		{
			g:   goodgeo.NewMultiPoint(goodgeo.XYM).MustSetCoords([]goodgeo.Coord{{1, 2, 3}, {4, 5, 6}}),
			xdr: goodgeotest.MustHexDecode("00000007d40000000200000007d13ff00000000000004000000000000000400800000000000000000007d1401000000000000040140000000000004018000000000000"),
			ndr: goodgeotest.MustHexDecode("01d40700000200000001d1070000000000000000f03f0000000000000040000000000000084001d1070000000000000000104000000000000014400000000000001840"),
		},
		{
			g:   goodgeo.NewMultiPoint(goodgeo.XYZM).MustSetCoords([]goodgeo.Coord{{1, 2, 3, 4}, {5, 6, 7, 8}}),
			xdr: goodgeotest.MustHexDecode("0000000bbc000000020000000bb93ff00000000000004000000000000000400800000000000040100000000000000000000bb940140000000000004018000000000000401c0000000000004020000000000000"),
			ndr: goodgeotest.MustHexDecode("01bc0b00000200000001b90b0000000000000000f03f00000000000000400000000000000840000000000000104001b90b0000000000000000144000000000000018400000000000001c400000000000002040"),
		},
		{
			g:    goodgeo.NewMultiPoint(goodgeo.XY).MustSetCoords([]goodgeo.Coord{nil, {1, 2}, {3, 4}}),
			opts: []wkbcommon.WKBOption{wkbcommon.WKBOptionEmptyPointHandling(wkbcommon.EmptyPointHandlingNaN)},
			xdr:  goodgeotest.MustHexDecode("00000000040000000300000000017ff80000000000007ff800000000000000000000013ff00000000000004000000000000000000000000140080000000000004010000000000000"),
			ndr:  goodgeotest.MustHexDecode("0104000000030000000101000000000000000000f87f000000000000f87f0101000000000000000000f03f0000000000000040010100000000000000000008400000000000001040"),
		},
		{
			g: goodgeo.NewMultiPolygon(goodgeo.XY).MustSetCoords([][][]goodgeo.Coord{
				{{{1, 2}, {2, 3}, {3, 4}, {1, 2}}},
				nil,
				{{{1, 2}, {3, 4}, {5, 6}, {1, 2}}},
			}),
			xdr: goodgeotest.MustHexDecode("000000000600000003000000000300000001000000043ff0000000000000400000000000000040000000000000004008000000000000400800000000000040100000000000003ff00000000000004000000000000000000000000300000000000000000300000001000000043ff0000000000000400000000000000040080000000000004010000000000000401400000000000040180000000000003ff00000000000004000000000000000"),
			ndr: goodgeotest.MustHexDecode("01060000000300000001030000000100000004000000000000000000f03f00000000000000400000000000000040000000000000084000000000000008400000000000001040000000000000f03f000000000000004001030000000000000001030000000100000004000000000000000000f03f00000000000000400000000000000840000000000000104000000000000014400000000000001840000000000000f03f0000000000000040"),
		},
		{
			g:   goodgeo.NewGeometryCollection().MustSetLayout(goodgeo.XY),
			xdr: goodgeotest.MustHexDecode("000000000700000000"),
			ndr: goodgeotest.MustHexDecode("010700000000000000"),
		},
		{
			g:   goodgeo.NewGeometryCollection().MustSetLayout(goodgeo.XYZ),
			ndr: goodgeotest.MustHexDecode("01ef03000000000000"),
			xdr: goodgeotest.MustHexDecode("00000003ef00000000"),
		},
		{
			g:   goodgeo.NewGeometryCollection().MustSetLayout(goodgeo.XYM),
			ndr: goodgeotest.MustHexDecode("01d707000000000000"),
			xdr: goodgeotest.MustHexDecode("00000007d700000000"),
		},
		{
			g:   goodgeo.NewGeometryCollection().MustSetLayout(goodgeo.XYZM),
			ndr: goodgeotest.MustHexDecode("01bf0b000000000000"),
			xdr: goodgeotest.MustHexDecode("0000000bbf00000000"),
		},
		{
			g: goodgeo.NewGeometryCollection().MustPush(
				goodgeo.NewPoint(goodgeo.XY).MustSetCoords(goodgeo.Coord{-79.3698576, 43.6456613}),
				goodgeo.NewLineString(goodgeo.XY).MustSetCoords([]goodgeo.Coord{{-79.3707986, 43.6453697}, {-79.3704747, 43.6454819}, {-79.370186, 43.6455592}, {-79.3699323, 43.6456385}, {-79.3698576, 43.6456613}}),
				goodgeo.NewLineString(goodgeo.XY).MustSetCoords([]goodgeo.Coord{{-79.3698576, 43.6456613}, {-79.3698057, 43.6455265}}),
			),
			xdr: goodgeotest.MustHexDecode("0000000007000000030000000001c053d7abbf360b554045d2a5078be57c000000000200000005c053d7bb2a0d19c44045d29b796daa28c053d7b5db841fb54045d29f26a15479c053d7b1209edbf94045d2a1af11d0e3c053d7acf8868efb4045d2a4484944edc053d7abbf360b554045d2a5078be57c000000000200000002c053d7abbf360b554045d2a5078be57cc053d7aae586d7f64045d2a09cc319c6"),
			ndr: goodgeotest.MustHexDecode("0107000000030000000101000000550B36BFABD753C07CE58B07A5D24540010200000005000000C4190D2ABBD753C028AA6D799BD24540B51F84DBB5D753C07954A1269FD24540F9DB9E20B1D753C0E3D011AFA1D24540FB8E86F8ACD753C0ED444948A4D24540550B36BFABD753C07CE58B07A5D24540010200000002000000550B36BFABD753C07CE58B07A5D24540F6D786E5AAD753C0C619C39CA0D24540"),
		},
	} {
		t.Run(fmt.Sprintf("ndr:%x", tc.ndr), func(t *testing.T) {
			test(t, tc.g, tc.xdr, tc.ndr, tc.opts...)
		})
	}

	t.Run("errors when encoding empty point WKBs by default", func(t *testing.T) {
		_, err := Marshal(goodgeo.NewPointEmpty(goodgeo.XY), binary.LittleEndian)
		matchStr := "cannot encode empty Point in WKB"
		if err == nil || err.Error() != matchStr {
			t.Errorf("expected error matching %s, got %#v", matchStr, err)
		}
	})

	t.Run("errors when encoding empty MultiPoint WKBs by default", func(t *testing.T) {
		_, err := Marshal(goodgeo.NewMultiPoint(goodgeo.XY).MustSetCoords([]goodgeo.Coord{nil, {1, 2}}), binary.LittleEndian)
		matchStr := "cannot encode empty Point in WKB"
		if err == nil || err.Error() != matchStr {
			t.Errorf("expected error matching %s, got %#v", matchStr, err)
		}
	})
}

func TestRandom(t *testing.T) {
	for _, tc := range testdata.Random {
		test(t, tc.G, nil, tc.WKB)
	}
}

func BenchmarkUnmarshal(b *testing.B) {
	for range b.N {
		for _, tc := range testdata.Random {
			if _, err := Unmarshal(tc.WKB); err != nil {
				b.Errorf("unmarshal error %v", err)
			}
		}
	}
}

func BenchmarkMarshal(b *testing.B) {
	for range b.N {
		for _, tc := range testdata.Random {
			if _, err := Marshal(tc.G, NDR); err != nil {
				b.Errorf("marshal error %v", err)
			}
		}
	}
}

func TestCrashes(t *testing.T) {
	// FIXME this test modifies a global variable. It will be racy if tests are
	// run in parallel.
	savedMaxGeometryElements := wkbcommon.MaxGeometryElements
	defer func() {
		wkbcommon.MaxGeometryElements = savedMaxGeometryElements
	}()
	wkbcommon.MaxGeometryElements[1] = 1 << 20
	for _, tc := range []struct {
		s    string
		want error
	}{
		{
			s: "\x01\x03\x00\x00\x00\x04\x00\x00\x00\a\x00\x00tٽ&\xf2\xa6\xd0\x1a" +
				"\xce\xc7\x1a\xfd67\xa3\x98Y.\xa5\xfbH\x1b\xe7|\xbe\xac\xfd%" +
				";\x05\\\x90c\x83\xe9g\x01\xcbk\xa3\xc8\xdb\x0f\xae\x16bYl" +
				"\x1b\x1a\xae\xe0\x95=o\x85/\xec\xd2~\xf3\xce\xe7\xad\x04\x92\xc3\xea" +
				"r\xacE\xe3A\u008cR\x86sb\xd5sҙ\u007f\x82\xec\x88\xff" +
				"\x8aM\xa7\u007f;\x9b\x93\xa2tٽ&\xf2\xa6\xd0\x1a\xce\xc7\x1a\xfd" +
				"67\xa3\x98\x05\x00\x00\x004\xed\x19\x9c/\x8ej\ue643\x018" +
				"?\x01|\x02\xa2\xad\x18Wyʡ\xb4h\xc1j\xf6\xbb\xf0=\xbf" +
				"\x03d%\xe6PsyQ\xce4pѹ\x1dcR\xadr\x14\t" +
				"\x02pm\x86=_\xfb%\x81\"\xde\xdf4\xed\x19\x9c/\x8ej\xee" +
				"\x99\x83\x018?\x01|\x02\x05\x00\x00\x00\xfb#\xbf\xc8\xe2i\xe9'" +
				"<(\xa3\u05ccz\x06a\x8e\x17<\x956\xa4\\K\xccy\u05f7" +
				"\xcc\xdfԴp.\x9b\xce\xef0nx}\xe9\xfc\x10\xf7?\xc9\xcc" +
				"!,\xab\x15}*;\x84K\xeco\u07b6$_\xea\xfb#\xbf\xc8" +
				"\xe2i\xe9'<(\xa3\u05ccz\x06a\x04\x00\x00\x00\x8f\x8a\x9f9" +
				"\x81\x10h!N\xdcf\n\xf0-\xeaL\x02\xba\xe9\x03\xd6/G\xc2" +
				"\x1cj\r\xd8 \xbc\xd6r\x05աTS\xb3\xa5\xdc\xd8\xfb\")" +
				"\xab\x19\xf7̏\x8a\x9f9\x81\x10h!N\xdcf\n\xf0-\xeaL",
			want: wkbcommon.ErrGeometryTooLarge{Level: 1, N: 1946157063, Limit: wkbcommon.MaxGeometryElements[1]},
		},
	} {
		t.Run(tc.s, func(t *testing.T) {
			_, err := Unmarshal([]byte(tc.s))
			assert.Equal(t, tc.want, err)
		})
	}
}
