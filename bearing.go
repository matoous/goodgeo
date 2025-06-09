package goodgeo

import "math"

// Bearing takes two [Coord] and finds the geographic bearing between them,
// i.e. the angle measured in degrees from the north line (0 degrees).
// Returns bearing in decimal degrees, between -180 and 180 degrees (positive clockwise)
func Bearing(start, end Coord, final ...bool) Degrees {
	// Reverse calculation
	if len(final) > 0 && final[0] {
		return CalculateFinalBearing(start, end)
	}

	lon1 := DegreesToRadians(Degrees(start[0]))
	lon2 := DegreesToRadians(Degrees(end[0]))
	lat1 := DegreesToRadians(Degrees(start[1]))
	lat2 := DegreesToRadians(Degrees(end[1]))

	a := math.Sin(float64(lon2-lon1)) * math.Cos(float64(lat2))
	b := math.Cos(float64(lat1))*math.Sin(float64(lat2)) -
		math.Sin(float64(lat1))*math.Cos(float64(lat2))*math.Cos(float64(lon2-lon1))

	return RadiansToDegrees(Radians(math.Atan2(a, b)))
}

// CalculateFinalBearing calculates Final Bearing
func CalculateFinalBearing(start, end Coord) Degrees {
	// Swap start & end
	bear := Bearing(end, start)
	return Degrees(math.Mod(float64(bear)+180, 360))
}
