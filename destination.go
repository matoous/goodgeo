package goodgeo

import "math"

// Destination takes a [Point] and calculates the location of a destination
// [Point] given a distance and bearing in degrees.
// This uses the [Haversine formula](http://en.wikipedia.org/wiki/Haversine_formula) to account for global curvature.
func Destination(origin *Point, distance Meters, bearing Degrees) *Point {
	coordinates1 := origin.Coords()
	longitude1 := DegreesToRadians(Degrees(coordinates1[0]))
	latitude1 := DegreesToRadians(Degrees(coordinates1[1]))
	bearingRad := DegreesToRadians(bearing)
	radians := LengthToRadians(distance)

	latitude2 := Radians(math.Asin(
		math.Sin(float64(latitude1))*math.Cos(float64(radians)) +
			math.Cos(float64(latitude1))*math.Sin(float64(radians))*math.Cos(float64(bearingRad)),
	))
	longitude2 := longitude1 +
		Radians(math.Atan2(
			math.Sin(float64(bearingRad))*math.Sin(float64(radians))*math.Cos(float64(latitude1)),
			math.Cos(float64(radians))-math.Sin(float64(latitude1))*math.Sin(float64(latitude2)),
		))

	lng := RadiansToDegrees(longitude2)
	lat := RadiansToDegrees(latitude2)

	if origin.layout == XYZ {
		return NewPointFlat(origin.layout, []float64{float64(lng), float64(lat), coordinates1[2]})
	}

	return NewPointFlat(origin.layout, []float64{float64(lng), float64(lat)})
}
