<div align="center">

# GoodGeo

[![Go Reference](https://pkg.go.dev/badge/github.com/matoous/goodgeo.svg)](https://pkg.go.dev/github.com/matoous/goodgeo)
[![CI Status](https://github.com/matoous/goodgeo/workflows/CI/badge.svg)](https://github.com/matoous/goodgeo/actions/workflows/ci.yml)
[![License](https://img.shields.io/github/license/matoous/goodgeo)](./LICENSE)

</div>

GoodGeo provides geospatial data types, processing, and encoding.

This package is based [twpayne/go-geom](https://github.com/twpayne/go-geom), but assumes that all coordinates are in geographic coordinate reference system, using the World Geodetic System 1984 ([WGS84](https://apps.dtic.mil/sti/citations/ADA280358)) datum, with longitude and latitude units of decimal degrees.

GoodGeo also implements geospatial analysis functions inspired by [Turf.js](https://github.com/Turfjs/turf).

## Philosophy

GoodGeo doesn't aims to be convenient, easy to use, and provide wide range of tools for geospatial analysis. Speed and efficiency is important but for GoodGeo is second on the list of priorities.

## Installation

Import the SDK using:

```go
import (
  "github.com/matoous/goodgeo"
)
```

And run any of `go build`/`go install`/`go test` which will resolve the package automatically.

Alternatively, you can install the SDK using:

```bash
go get github.com/matoous/goodgeo
```

## Example

```go
package main

import (
  "fmt"
  "os"

  "github.com/matoous/goodgeo"
)

func main() {
  data, err := os.ReadFile("geo.json")
  if err != nil {
    panic(err)
  }

  var fc geojson.FeatureCollection
  err := json.Unmarshal(data, &fc)
  if err != nil {
    panic(err)
  }

  ls := gg.Features[0].Geometry.(*geom.LineString)
  length := goodgeo.Length(fc)
  fmt.Println(length)
}
```

## Key features

* OpenGeo Consortium-style geometries.
* Support for 2D and 3D geometries, measures (time and/or distance), and
  unlimited extra dimensions.
* Encoding and decoding of common geometry formats (GeoJSON, KML, WKB, and
  others) including [`sql.Scanner`](https://pkg.go.dev/database/sql#Scanner) and
  [`driver.Value`](https://pkg.go.dev/database/sql/driver#Value) interface
  implementations for easy database integration.
* [2D](https://pkg.go.dev/github.com/matoous/goodgeo/xy) and
  [3D](https://pkg.go.dev/github.com/matoous/goodgeo/xyz) topology functions.
* Efficient, cache-friendly [internal representation](INTERNALS.md).
* Optional protection against malicious or malformed inputs.

## Examples

* [PostGIS, EWKB, and GeoJSON](https://github.com/matoous/goodgeo/tree/master/examples/postgis).

## Detailed features

### Geometry types

* [Point](https://pkg.go.dev/github.com/matoous/goodgeo#Point)
* [LineString](https://pkg.go.dev/github.com/matoous/goodgeo#LineString)
* [Polygon](https://pkg.go.dev/github.com/matoous/goodgeo#Polygon)
* [MultiPoint](https://pkg.go.dev/github.com/matoous/goodgeo#MultiPoint)
* [MultiLineString](https://pkg.go.dev/github.com/matoous/goodgeo#MultiLineString)
* [MultiPolygon](https://pkg.go.dev/github.com/matoous/goodgeo#MultiPolygon)
* [GeometryCollection](https://pkg.go.dev/github.com/matoous/goodgeo#GeometryCollection)

### Encoding and decoding

* [GeoJSON](https://pkg.go.dev/github.com/matoous/goodgeo/encoding/geojson)
* [GPX](https://pkg.go.dev/github.com/matoous/goodgeo/encoding/gpx)
* [KML](https://pkg.go.dev/github.com/matoous/goodgeo/encoding/kml)
* [Polyline](https://pkg.go.dev/github.com/matoous/goodgeo/encoding/polyline)
* [IGC](https://pkg.go.dev/github.com/matoous/goodgeo/encoding/igc)
* [KML](https://pkg.go.dev/github.com/matoous/goodgeo/encoding/kml) (encoding only)
* [WKB](https://pkg.go.dev/github.com/matoous/goodgeo/encoding/wkb)
* [EWKB](https://pkg.go.dev/github.com/matoous/goodgeo/encoding/ewkb)
* [WKT](https://pkg.go.dev/github.com/matoous/goodgeo/encoding/wkt) (encoding only)
* [WKB Hex](https://pkg.go.dev/github.com/matoous/goodgeo/encoding/wkbhex)
* [EWKB Hex](https://pkg.go.dev/github.com/matoous/goodgeo/encoding/ewkbhex)

## Protection against malicious or malformed inputs

The WKB and EWKB formats encode geometry sizes, and memory is allocated for
those geometries. If the input is malicious or malformed, the memory allocation
can be very large, leading to a memory starvation denial-of-service attack
against the server. For example, a client might send a `MultiPoint` with header
indicating that it contains 2^32-1 points. This will result in the server
reading that geometry to allocate 2 × `sizeof(float64)` × (2^32-1) = 64GB of
memory to store those points. By default, malicious or malformed input
protection is disabled, but can be enabled by setting positive values for
`wkbcommon.MaxGeometryElements`.

## License

BSD-2-Clause
