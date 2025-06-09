// Package wkb implements Well Known Binary encoding and decoding.
//
// If you are encoding geometries in WKB to send to PostgreSQL/PostGIS, then
// you must specify binary_parameters=yes in the data source name that you pass
// to sql.Open.
package wkb

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"

	"github.com/matoous/goodgeo"
	"github.com/matoous/goodgeo/encoding/wkbcommon"
)

var (
	// XDR is big endian.
	XDR = wkbcommon.XDR
	// NDR is little endian.
	NDR = wkbcommon.NDR
)

const (
	wkbXYID   = 0
	wkbXYZID  = 1000
	wkbXYMID  = 2000
	wkbXYZMID = 3000
)

// Read reads an arbitrary geometry from r.
func Read(r io.Reader, opts ...wkbcommon.WKBOption) (goodgeo.T, error) {
	params := wkbcommon.InitWKBParams(
		wkbcommon.WKBParams{
			EmptyPointHandling: wkbcommon.EmptyPointHandlingError,
		},
		opts...,
	)

	wkbByteOrder, err := wkbcommon.ReadByte(r)
	if err != nil {
		return nil, err
	}
	var byteOrder binary.ByteOrder
	switch wkbByteOrder {
	case wkbcommon.XDRID:
		byteOrder = XDR
	case wkbcommon.NDRID:
		byteOrder = NDR
	default:
		return nil, wkbcommon.ErrUnknownByteOrder(wkbByteOrder)
	}

	wkbGeometryType, err := wkbcommon.ReadUInt32(r, byteOrder)
	if err != nil {
		return nil, err
	}
	t := wkbcommon.Type(wkbGeometryType)

	var layout goodgeo.Layout
	switch 1000 * (t / 1000) {
	case wkbXYID:
		layout = goodgeo.XY
	case wkbXYZID:
		layout = goodgeo.XYZ
	case wkbXYMID:
		layout = goodgeo.XYM
	case wkbXYZMID:
		layout = goodgeo.XYZM
	default:
		return nil, wkbcommon.ErrUnknownType(t)
	}

	switch t % 1000 {
	case wkbcommon.PointID:
		flatCoords, err := wkbcommon.ReadFlatCoords0(r, byteOrder, layout.Stride())
		if err != nil {
			return nil, err
		}
		if params.EmptyPointHandling == wkbcommon.EmptyPointHandlingNaN {
			return goodgeo.NewPointFlatMaybeEmpty(layout, flatCoords), nil
		}
		return goodgeo.NewPointFlat(layout, flatCoords), nil
	case wkbcommon.LineStringID:
		flatCoords, err := wkbcommon.ReadFlatCoords1(r, byteOrder, layout.Stride())
		if err != nil {
			return nil, err
		}
		return goodgeo.NewLineStringFlat(layout, flatCoords), nil
	case wkbcommon.PolygonID:
		flatCoords, ends, err := wkbcommon.ReadFlatCoords2(r, byteOrder, layout.Stride())
		if err != nil {
			return nil, err
		}
		return goodgeo.NewPolygonFlat(layout, flatCoords, ends), nil
	case wkbcommon.MultiPointID:
		n, err := wkbcommon.ReadUInt32(r, byteOrder)
		if err != nil {
			return nil, err
		}
		if limit := wkbcommon.MaxGeometryElements[1]; limit >= 0 && int(n) > limit {
			return nil, wkbcommon.ErrGeometryTooLarge{Level: 1, N: int(n), Limit: limit}
		}
		mp := goodgeo.NewMultiPoint(layout)
		for range n {
			g, err := Read(r, opts...)
			if err != nil {
				return nil, err
			}
			p, ok := g.(*goodgeo.Point)
			if !ok {
				return nil, wkbcommon.ErrUnexpectedType{Got: g, Want: &goodgeo.Point{}}
			}
			if err = mp.Push(p); err != nil {
				return nil, err
			}
		}
		return mp, nil
	case wkbcommon.MultiLineStringID:
		n, err := wkbcommon.ReadUInt32(r, byteOrder)
		if err != nil {
			return nil, err
		}
		if limit := wkbcommon.MaxGeometryElements[2]; limit >= 0 && int(n) > limit {
			return nil, wkbcommon.ErrGeometryTooLarge{Level: 2, N: int(n), Limit: limit}
		}
		mls := goodgeo.NewMultiLineString(layout)
		for range n {
			g, err := Read(r, opts...)
			if err != nil {
				return nil, err
			}
			p, ok := g.(*goodgeo.LineString)
			if !ok {
				return nil, wkbcommon.ErrUnexpectedType{Got: g, Want: &goodgeo.LineString{}}
			}
			if err = mls.Push(p); err != nil {
				return nil, err
			}
		}
		return mls, nil
	case wkbcommon.MultiPolygonID:
		n, err := wkbcommon.ReadUInt32(r, byteOrder)
		if err != nil {
			return nil, err
		}
		if limit := wkbcommon.MaxGeometryElements[3]; limit >= 0 && int(n) > limit {
			return nil, wkbcommon.ErrGeometryTooLarge{Level: 3, N: int(n), Limit: limit}
		}
		mp := goodgeo.NewMultiPolygon(layout)
		for range n {
			g, err := Read(r, opts...)
			if err != nil {
				return nil, err
			}
			p, ok := g.(*goodgeo.Polygon)
			if !ok {
				return nil, wkbcommon.ErrUnexpectedType{Got: g, Want: &goodgeo.Polygon{}}
			}
			if err = mp.Push(p); err != nil {
				return nil, err
			}
		}
		return mp, nil
	case wkbcommon.GeometryCollectionID:
		n, err := wkbcommon.ReadUInt32(r, byteOrder)
		if err != nil {
			return nil, err
		}
		gc := goodgeo.NewGeometryCollection()
		for range n {
			g, err := Read(r, opts...)
			if err != nil {
				return nil, err
			}
			if err := gc.Push(g); err != nil {
				return nil, err
			}
		}
		// If EMPTY, mark the collection with a fixed layout to differentiate
		// GEOMETRYCOLLECTION EMPTY between 2D/Z/M/ZM.
		if gc.Empty() && gc.NumGeoms() == 0 {
			if err := gc.SetLayout(layout); err != nil {
				return nil, err
			}
		}
		return gc, nil
	default:
		return nil, wkbcommon.ErrUnsupportedType(wkbGeometryType)
	}
}

// Unmarshal unmrshals an arbitrary geometry from a []byte.
func Unmarshal(data []byte, opts ...wkbcommon.WKBOption) (goodgeo.T, error) {
	return Read(bytes.NewBuffer(data), opts...)
}

// Write writes an arbitrary geometry to w.
func Write(w io.Writer, byteOrder binary.ByteOrder, g goodgeo.T, opts ...wkbcommon.WKBOption) error {
	params := wkbcommon.InitWKBParams(
		wkbcommon.WKBParams{
			EmptyPointHandling: wkbcommon.EmptyPointHandlingError,
		},
		opts...,
	)

	var wkbByteOrder byte
	switch byteOrder {
	case XDR:
		wkbByteOrder = wkbcommon.XDRID
	case NDR:
		wkbByteOrder = wkbcommon.NDRID
	default:
		return wkbcommon.ErrUnsupportedByteOrder{}
	}
	if err := wkbcommon.WriteByte(w, wkbByteOrder); err != nil {
		return err
	}

	var wkbGeometryType uint32
	switch g.(type) {
	case *goodgeo.Point:
		wkbGeometryType = wkbcommon.PointID
	case *goodgeo.LineString:
		wkbGeometryType = wkbcommon.LineStringID
	case *goodgeo.Polygon:
		wkbGeometryType = wkbcommon.PolygonID
	case *goodgeo.MultiPoint:
		wkbGeometryType = wkbcommon.MultiPointID
	case *goodgeo.MultiLineString:
		wkbGeometryType = wkbcommon.MultiLineStringID
	case *goodgeo.MultiPolygon:
		wkbGeometryType = wkbcommon.MultiPolygonID
	case *goodgeo.GeometryCollection:
		wkbGeometryType = wkbcommon.GeometryCollectionID
	default:
		return goodgeo.UnsupportedTypeError{Value: g}
	}
	switch g.Layout() {
	case goodgeo.NoLayout:
		// Special case for empty GeometryCollections
		if _, ok := g.(*goodgeo.GeometryCollection); !ok || !g.Empty() {
			return goodgeo.UnsupportedLayoutError(g.Layout())
		}
	case goodgeo.XY:
		wkbGeometryType += wkbXYID
	case goodgeo.XYZ:
		wkbGeometryType += wkbXYZID
	case goodgeo.XYM:
		wkbGeometryType += wkbXYMID
	case goodgeo.XYZM:
		wkbGeometryType += wkbXYZMID
	default:
		return goodgeo.UnsupportedLayoutError(g.Layout())
	}
	if err := wkbcommon.WriteUInt32(w, byteOrder, wkbGeometryType); err != nil {
		return err
	}

	switch g := g.(type) {
	case *goodgeo.Point:
		if g.Empty() {
			switch params.EmptyPointHandling {
			case wkbcommon.EmptyPointHandlingNaN:
				return wkbcommon.WriteEmptyPointAsNaN(w, byteOrder, g.Stride())
			case wkbcommon.EmptyPointHandlingError:
				return errors.New("cannot encode empty Point in WKB")
			default:
				return fmt.Errorf("cannot encode empty Point in WKB (unknown option: %d)", wkbcommon.EmptyPointHandlingNaN)
			}
		}
		return wkbcommon.WriteFlatCoords0(w, byteOrder, g.FlatCoords())
	case *goodgeo.LineString:
		return wkbcommon.WriteFlatCoords1(w, byteOrder, g.FlatCoords(), g.Stride())
	case *goodgeo.Polygon:
		return wkbcommon.WriteFlatCoords2(w, byteOrder, g.FlatCoords(), g.Ends(), g.Stride())
	case *goodgeo.MultiPoint:
		n := g.NumPoints()
		if err := wkbcommon.WriteUInt32(w, byteOrder, uint32(n)); err != nil {
			return err
		}
		for i := range n {
			if err := Write(w, byteOrder, g.Point(i), opts...); err != nil {
				return err
			}
		}
		return nil
	case *goodgeo.MultiLineString:
		n := g.NumLineStrings()
		if err := wkbcommon.WriteUInt32(w, byteOrder, uint32(n)); err != nil {
			return err
		}
		for i := range n {
			if err := Write(w, byteOrder, g.LineString(i), opts...); err != nil {
				return err
			}
		}
		return nil
	case *goodgeo.MultiPolygon:
		n := g.NumPolygons()
		if err := wkbcommon.WriteUInt32(w, byteOrder, uint32(n)); err != nil {
			return err
		}
		for i := range n {
			if err := Write(w, byteOrder, g.Polygon(i), opts...); err != nil {
				return err
			}
		}
		return nil
	case *goodgeo.GeometryCollection:
		n := g.NumGeoms()
		if err := wkbcommon.WriteUInt32(w, byteOrder, uint32(n)); err != nil {
			return err
		}
		for i := range n {
			if err := Write(w, byteOrder, g.Geom(i), opts...); err != nil {
				return err
			}
		}
		return nil
	default:
		return goodgeo.UnsupportedTypeError{Value: g}
	}
}

// Marshal marshals an arbitrary geometry to a []byte.
func Marshal(g goodgeo.T, byteOrder binary.ByteOrder, opts ...wkbcommon.WKBOption) ([]byte, error) {
	w := bytes.NewBuffer(nil)
	if err := Write(w, byteOrder, g, opts...); err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}
