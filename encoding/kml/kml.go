// Package kml implements KML encoding.
package kml

import (
	"github.com/twpayne/go-kml/v3"

	"github.com/matoous/goodgeo"
)

// Encode encodes an arbitrary geometry.
func Encode(g goodgeo.T) (kml.Element, error) {
	switch g := g.(type) {
	case *goodgeo.Point:
		return EncodePoint(g), nil
	case *goodgeo.LineString:
		return EncodeLineString(g), nil
	case *goodgeo.LinearRing:
		return EncodeLinearRing(g), nil
	case *goodgeo.MultiLineString:
		return EncodeMultiLineString(g), nil
	case *goodgeo.MultiPoint:
		return EncodeMultiPoint(g), nil
	case *goodgeo.MultiPolygon:
		return EncodeMultiPolygon(g), nil
	case *goodgeo.Polygon:
		return EncodePolygon(g), nil
	case *goodgeo.GeometryCollection:
		return EncodeGeometryCollection(g)
	default:
		return nil, goodgeo.UnsupportedTypeError{Value: g}
	}
}

// EncodeLineString encodes a LineString.
func EncodeLineString(ls *goodgeo.LineString) kml.Element {
	flatCoords := ls.FlatCoords()
	return kml.LineString(kml.CoordinatesFlat(flatCoords, 0, len(flatCoords), ls.Stride(), dim(ls.Layout())))
}

// EncodeLinearRing encodes a LinearRing.
func EncodeLinearRing(lr *goodgeo.LinearRing) kml.Element {
	flatCoords := lr.FlatCoords()
	return kml.LinearRing(kml.CoordinatesFlat(flatCoords, 0, len(flatCoords), lr.Stride(), dim(lr.Layout())))
}

// EncodeMultiLineString encodes a MultiLineString.
func EncodeMultiLineString(mls *goodgeo.MultiLineString) kml.Element {
	lineStrings := make([]kml.Element, mls.NumLineStrings())
	flatCoords := mls.FlatCoords()
	ends := mls.Ends()
	stride := mls.Stride()
	d := dim(mls.Layout())
	offset := 0
	for i, end := range ends {
		lineStrings[i] = kml.LineString(kml.CoordinatesFlat(flatCoords, offset, end, stride, d))
		offset = end
	}
	return kml.MultiGeometry(lineStrings...)
}

// EncodeMultiPoint encodes a MultiPoint.
func EncodeMultiPoint(mp *goodgeo.MultiPoint) kml.Element {
	points := make([]kml.Element, mp.NumPoints())
	flatCoords := mp.FlatCoords()
	stride := mp.Stride()
	d := dim(mp.Layout())
	for i, offset, end := 0, 0, len(flatCoords); offset < end; i++ {
		points[i] = kml.Point(kml.CoordinatesFlat(flatCoords, offset, offset+stride, stride, d))
		offset += stride
	}
	return kml.MultiGeometry(points...)
}

// EncodeMultiPolygon encodes a MultiPolygon.
func EncodeMultiPolygon(mp *goodgeo.MultiPolygon) kml.Element {
	polygons := make([]kml.Element, mp.NumPolygons())
	flatCoords := mp.FlatCoords()
	endss := mp.Endss()
	stride := mp.Stride()
	d := dim(mp.Layout())
	offset := 0
	for i, ends := range endss {
		boundaries := make([]kml.Element, len(ends))
		for j, end := range ends {
			linearRing := kml.LinearRing(kml.CoordinatesFlat(flatCoords, offset, end, stride, d))
			if j == 0 {
				boundaries[j] = kml.OuterBoundaryIs(linearRing)
			} else {
				boundaries[j] = kml.InnerBoundaryIs(linearRing)
			}
			offset = end
		}
		polygons[i] = kml.Polygon(boundaries...)
	}
	return kml.MultiGeometry(polygons...)
}

// EncodePoint encodes a Point.
func EncodePoint(p *goodgeo.Point) kml.Element {
	flatCoords := p.FlatCoords()
	return kml.Point(kml.CoordinatesFlat(flatCoords, 0, len(flatCoords), p.Stride(), dim(p.Layout())))
}

// EncodePolygon encodes a Polygon.
func EncodePolygon(p *goodgeo.Polygon) kml.Element {
	boundaries := make([]kml.Element, p.NumLinearRings())
	stride := p.Stride()
	flatCoords := p.FlatCoords()
	d := dim(p.Layout())
	offset := 0
	for i, end := range p.Ends() {
		linearRing := kml.LinearRing(kml.CoordinatesFlat(flatCoords, offset, end, stride, d))
		if i == 0 {
			boundaries[i] = kml.OuterBoundaryIs(linearRing)
		} else {
			boundaries[i] = kml.InnerBoundaryIs(linearRing)
		}
		offset = end
	}
	return kml.Polygon(boundaries...)
}

// EncodeGeometryCollection encodes a GeometryCollection.
func EncodeGeometryCollection(g *goodgeo.GeometryCollection) (kml.Element, error) {
	geometries := make([]kml.Element, g.NumGeoms())
	for i, g := range g.Geoms() {
		var err error
		geometries[i], err = Encode(g)
		if err != nil {
			return nil, err
		}
	}
	return kml.MultiGeometry(geometries...), nil
}

func dim(l goodgeo.Layout) int {
	switch l {
	case goodgeo.XY, goodgeo.XYM:
		return 2
	default:
		return 3
	}
}
