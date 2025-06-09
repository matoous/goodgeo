package goodgeo

import (
	"math"
)

// RamerDouglasPeucker simplifies a line string using the
// [Ramer–Douglas–Peucker](https://en.wikipedia.org/wiki/Ramer%E2%80%93Douglas%E2%80%93Peucker_algorithm) algorithm.
func RamerDouglasPeucker(ls *LineString, epsilon Meters) *LineString {
	nls := NewLineString(ls.layout)

	coords := ls.Coords()
	n := len(coords)
	var simplified []Coord

	if n == 0 {
		return nil
	}

	if n == 1 {
		nls.MustSetCoords([]Coord{coords[0]})
		return nil
	}

	simplified = append(simplified, coords[0])
	ramerDouglasPeuckerRecursive(coords, epsilon, 0, n-1, &simplified)
	simplified = append(simplified, coords[n-1])

	nls.MustSetCoords(simplified)
	return nls
}

func ramerDouglasPeuckerRecursive(
	points []Coord,
	epsilon Meters,
	start, end int,
	simplified *[]Coord,
) {
	var (
		largestIndex    = -1
		largestDistance Meters
	)

	for i := start + 1; i < end; i++ {
		d := Crossarc(points[start], points[end], points[i])
		if d > largestDistance || largestIndex == -1 {
			largestIndex = i
			largestDistance = d
		}
	}

	if largestDistance > epsilon && largestIndex != -1 {
		ramerDouglasPeuckerRecursive(points, epsilon, start, largestIndex, simplified)
		*simplified = append(*simplified, points[largestIndex])
		ramerDouglasPeuckerRecursive(points, epsilon, largestIndex, end, simplified)
	}
}

func Crossarc(a, b, c Coord) Meters {
	bear12 := Bearing(a, b)
	bear13 := Bearing(a, c)
	dis13 := Distance(a, c)

	diff := math.Abs(float64(bear13 - bear12))
	if diff > math.Pi {
		diff = 2*math.Pi - diff
	}

	if diff > math.Pi/2 {
		return dis13
	}

	dxt := math.Asin(math.Sin(float64(dis13)/EarthRadius)*math.Sin(float64(bear13-bear12))) * EarthRadius

	dis12 := Distance(a, b)
	dis14 := math.Acos(math.Cos(float64(dis13)/EarthRadius)/math.Cos(dxt/EarthRadius)) * EarthRadius

	if dis14 > float64(dis12*EarthRadius) {
		return Distance(b, c)
	}

	return Meters(math.Abs(dxt))
}
