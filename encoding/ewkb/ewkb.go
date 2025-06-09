// Package ewkb implements Extended Well Known Binary encoding and decoding.
// See https://github.com/postgis/postgis/blob/2.1.0/doc/ZMSgeoms.txt.
//
// If you are encoding geometries in EWKB to send to PostgreSQL/PostGIS, then
// you must specify binary_parameters=yes in the data source name that you pass
// to sql.Open.
package ewkb

import (
	"bytes"
	"encoding/binary"
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
	ewkbZ    = 0x80000000
	ewkbM    = 0x40000000
	ewkbSRID = 0x20000000
)

// Read reads an arbitrary geometry from r.
func Read(r io.Reader) (goodgeo.T, error) {
	ewkbByteOrder, err := wkbcommon.ReadByte(r)
	if err != nil {
		return nil, err
	}
	var byteOrder binary.ByteOrder
	switch ewkbByteOrder {
	case wkbcommon.XDRID:
		byteOrder = XDR
	case wkbcommon.NDRID:
		byteOrder = NDR
	default:
		return nil, wkbcommon.ErrUnknownByteOrder(ewkbByteOrder)
	}

	ewkbGeometryType, err := wkbcommon.ReadUInt32(r, byteOrder)
	if err != nil {
		return nil, err
	}
	t := wkbcommon.Type(ewkbGeometryType)

	var layout goodgeo.Layout
	switch t & (ewkbZ | ewkbM) {
	case 0:
		layout = goodgeo.XY
	case ewkbZ:
		layout = goodgeo.XYZ
	case ewkbM:
		layout = goodgeo.XYM
	case ewkbZ | ewkbM:
		layout = goodgeo.XYZM
	default:
		return nil, wkbcommon.ErrUnknownType(t)
	}

	var srid uint32
	if ewkbGeometryType&ewkbSRID != 0 {
		srid, err = wkbcommon.ReadUInt32(r, byteOrder)
		if err != nil {
			return nil, err
		}
	}

	switch t &^ (ewkbZ | ewkbM | ewkbSRID) {
	case wkbcommon.PointID:
		flatCoords, err := wkbcommon.ReadFlatCoords0(r, byteOrder, layout.Stride())
		if err != nil {
			return nil, err
		}
		return goodgeo.NewPointFlatMaybeEmpty(layout, flatCoords).SetSRID(int(srid)), nil
	case wkbcommon.LineStringID:
		flatCoords, err := wkbcommon.ReadFlatCoords1(r, byteOrder, layout.Stride())
		if err != nil {
			return nil, err
		}
		return goodgeo.NewLineStringFlat(layout, flatCoords).SetSRID(int(srid)), nil
	case wkbcommon.PolygonID:
		flatCoords, ends, err := wkbcommon.ReadFlatCoords2(r, byteOrder, layout.Stride())
		if err != nil {
			return nil, err
		}
		return goodgeo.NewPolygonFlat(layout, flatCoords, ends).SetSRID(int(srid)), nil
	case wkbcommon.MultiPointID:
		n, err := wkbcommon.ReadUInt32(r, byteOrder)
		if err != nil {
			return nil, err
		}
		if limit := wkbcommon.MaxGeometryElements[1]; limit >= 0 && int(n) > limit {
			return nil, wkbcommon.ErrGeometryTooLarge{Level: 1, N: int(n), Limit: limit}
		}
		mp := goodgeo.NewMultiPoint(layout).SetSRID(int(srid))
		for range n {
			g, err := Read(r)
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
		mls := goodgeo.NewMultiLineString(layout).SetSRID(int(srid))
		for range n {
			g, err := Read(r)
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
		mp := goodgeo.NewMultiPolygon(layout).SetSRID(int(srid))
		for range n {
			g, err := Read(r)
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
		if limit := wkbcommon.MaxGeometryElements[1]; limit >= 0 && int(n) > limit {
			return nil, wkbcommon.ErrGeometryTooLarge{Level: 1, N: int(n), Limit: limit}
		}
		gc := goodgeo.NewGeometryCollection().SetSRID(int(srid))
		for range n {
			g, err := Read(r)
			if err != nil {
				return nil, err
			}
			if err = gc.Push(g); err != nil {
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
		return nil, wkbcommon.ErrUnsupportedType(ewkbGeometryType)
	}
}

// Unmarshal unmrshals an arbitrary geometry from a []byte.
func Unmarshal(data []byte) (goodgeo.T, error) {
	return Read(bytes.NewBuffer(data))
}

// Write writes an arbitrary geometry to w.
func Write(w io.Writer, byteOrder binary.ByteOrder, g goodgeo.T) error {
	var ewkbByteOrder byte
	switch byteOrder {
	case XDR:
		ewkbByteOrder = wkbcommon.XDRID
	case NDR:
		ewkbByteOrder = wkbcommon.NDRID
	default:
		return wkbcommon.ErrUnsupportedByteOrder{}
	}
	if err := binary.Write(w, byteOrder, ewkbByteOrder); err != nil {
		return err
	}

	var ewkbGeometryType uint32
	switch g.(type) {
	case *goodgeo.Point:
		ewkbGeometryType = wkbcommon.PointID
	case *goodgeo.LineString:
		ewkbGeometryType = wkbcommon.LineStringID
	case *goodgeo.Polygon:
		ewkbGeometryType = wkbcommon.PolygonID
	case *goodgeo.MultiPoint:
		ewkbGeometryType = wkbcommon.MultiPointID
	case *goodgeo.MultiLineString:
		ewkbGeometryType = wkbcommon.MultiLineStringID
	case *goodgeo.MultiPolygon:
		ewkbGeometryType = wkbcommon.MultiPolygonID
	case *goodgeo.GeometryCollection:
		ewkbGeometryType = wkbcommon.GeometryCollectionID
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
	case goodgeo.XYZ:
		ewkbGeometryType |= ewkbZ
	case goodgeo.XYM:
		ewkbGeometryType |= ewkbM
	case goodgeo.XYZM:
		ewkbGeometryType |= ewkbZ | ewkbM
	default:
		return goodgeo.UnsupportedLayoutError(g.Layout())
	}
	srid := g.SRID()
	if srid != 0 {
		ewkbGeometryType |= ewkbSRID
	}
	if err := binary.Write(w, byteOrder, ewkbGeometryType); err != nil {
		return err
	}
	if ewkbGeometryType&ewkbSRID != 0 {
		if err := binary.Write(w, byteOrder, uint32(srid)); err != nil {
			return err
		}
	}

	switch g := g.(type) {
	case *goodgeo.Point:
		if g.Empty() {
			return wkbcommon.WriteEmptyPointAsNaN(w, byteOrder, g.Stride())
		}
		return wkbcommon.WriteFlatCoords0(w, byteOrder, g.FlatCoords())
	case *goodgeo.LineString:
		return wkbcommon.WriteFlatCoords1(w, byteOrder, g.FlatCoords(), g.Stride())
	case *goodgeo.Polygon:
		return wkbcommon.WriteFlatCoords2(w, byteOrder, g.FlatCoords(), g.Ends(), g.Stride())
	case *goodgeo.MultiPoint:
		n := g.NumPoints()
		if err := binary.Write(w, byteOrder, uint32(n)); err != nil {
			return err
		}
		for i := range n {
			if err := Write(w, byteOrder, g.Point(i)); err != nil {
				return err
			}
		}
		return nil
	case *goodgeo.MultiLineString:
		n := g.NumLineStrings()
		if err := binary.Write(w, byteOrder, uint32(n)); err != nil {
			return err
		}
		for i := range n {
			if err := Write(w, byteOrder, g.LineString(i)); err != nil {
				return err
			}
		}
		return nil
	case *goodgeo.MultiPolygon:
		n := g.NumPolygons()
		if err := binary.Write(w, byteOrder, uint32(n)); err != nil {
			return err
		}
		for i := range n {
			if err := Write(w, byteOrder, g.Polygon(i)); err != nil {
				return err
			}
		}
		return nil
	case *goodgeo.GeometryCollection:
		n := g.NumGeoms()
		if err := binary.Write(w, byteOrder, uint32(n)); err != nil {
			return err
		}
		for i := range n {
			if err := Write(w, byteOrder, g.Geom(i)); err != nil {
				return err
			}
		}
		return nil
	default:
		return goodgeo.UnsupportedTypeError{Value: g}
	}
}

// Marshal marshals an arbitrary geometry to a []byte.
func Marshal(g goodgeo.T, byteOrder binary.ByteOrder) ([]byte, error) {
	w := bytes.NewBuffer(nil)
	if err := Write(w, byteOrder, g); err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}
