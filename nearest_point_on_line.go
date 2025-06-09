package goodgeo

import (
	"math"
)

type NearestPointResult struct {
	Point    Coord
	Distance Meters
	Index    int
	Location Meters
}

// NearestPointOnLine returns the closest point on a line to a given point.
func NearestPointOnLine(line *LineString, pt Coord) NearestPointResult {
	closest := NearestPointResult{
		Point:    nil,
		Distance: Meters(math.Inf(1)),
		Index:    -1,
		Location: -1,
	}

	ptVec := lngLatToVector(pt)

	coords := line.flatCoords
	stride := line.stride
	var length Meters

	for i := stride; i < len(coords); i += stride {
		start := coords[i-stride : i]
		stop := coords[i : i+stride]

		segLen := Distance(start, stop)

		var intersect Coord
		var useEnd bool

		switch {
		case EqualCoords(start, pt):
			intersect, useEnd = start, false
		case EqualCoords(stop, pt):
			intersect, useEnd = stop, true
		case EqualCoords(start, stop):
			intersect, useEnd = stop, true
		default:
			intersect, useEnd = nearestPointOnSegment(start, stop, ptVec)
		}

		d := Distance(pt, intersect)
		if d < closest.Distance {
			closest = NearestPointResult{
				Point:    intersect,
				Distance: d,
				Index:    i/stride + btoi(useEnd),
				Location: length + Distance(start, intersect),
			}
		}

		length += segLen
	}

	return closest
}

func EqualCoords(a, b Coord) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func nearestPointOnSegment(a, b Coord, cVec Vector) (Coord, bool) {
	// Based heavily on this article on finding cross track distance to an arc:
	// https://gis.stackexchange.com/questions/209540/projecting-cross-track-distance-on-great-circle

	// Convert spherical (lng, lat) to cartesian vector coords (x, y, z)
	// In the below https://tikz.net/spherical_1/ we convert lng (𝜙) and lat (𝜃)
	// into vectors with x, y, and z components with a length (r) of 1.
	A := lngLatToVector(a)
	B := lngLatToVector(b)
	C := cVec

	// Calculate coefficients.
	D, E, F := cross(A, B)
	aCoeff := E*C[2] - F*C[1]
	bCoeff := F*C[0] - D*C[2]
	cCoeff := D*C[1] - E*C[0]

	f := cCoeff*E - bCoeff*F
	g := aCoeff*F - cCoeff*D
	h := bCoeff*D - aCoeff*E

	// Vectors to the two points these great circles intersect.
	t := 1.0 / math.Sqrt(f*f+g*g+h*h)
	I1 := Vector{f * t, g * t, h * t}
	I2 := Vector{-f * t, -g * t, -h * t}

	// Figure out which is the closest intersection to this segment of the great
	// circle.
	angleAB := angle(A, B)
	angleAI1 := angle(A, I1)
	angleBI1 := angle(B, I1)
	angleAI2 := angle(A, I2)
	angleBI2 := angle(B, I2)

	var I Vector
	if (angleAI1 < angleAI2 && angleAI1 < angleBI2) || (angleBI1 < angleAI2 && angleBI1 < angleBI2) {
		I = I1
	} else {
		I = I2
	}

	// I is the closest intersection to the segment, though might not actually be
	// ON the segment.

	// If angle AI or BI is greater than angleAB, I lies on the circle *beyond* A
	// and B so use the closest of A or B as the intersection
	if angle(A, I) > angleAB || angle(B, I) > angleAB {
		if Distance(vectorToLngLat(I), vectorToLngLat(A)) <= Distance(vectorToLngLat(I), vectorToLngLat(B)) {
			return vectorToLngLat(A), false
		}
		return vectorToLngLat(B), true
	}

	return vectorToLngLat(I), false
}

type Vector = [3]float64

func dot(v1, v2 Vector) float64 {
	return v1[0]*v2[0] + v1[1]*v2[1] + v1[2]*v2[2]
}

// https://en.wikipedia.org/wiki/Cross_product
func cross(v1, v2 Vector) (float64, float64, float64) {
	return v1[1]*v2[2] - v1[2]*v2[1],
		v1[2]*v2[0] - v1[0]*v2[2],
		v1[0]*v2[1] - v1[1]*v2[0]
}

func magnitude(v Vector) float64 {
	return math.Sqrt(v[0]*v[0] + v[1]*v[1] + v[2]*v[2])
}

func angle(v1, v2 Vector) float64 {
	theta := dot(v1, v2) / (magnitude(v1) * magnitude(v2))
	return math.Acos(math.Min(math.Max(theta, -1), 1))
}

func lngLatToVector(pos Coord) Vector {
	lat := DegreesToRadians(Degrees(pos[1]))
	lng := DegreesToRadians(Degrees(pos[0]))
	return Vector{
		math.Cos(float64(lat)) * math.Cos(float64(lng)),
		math.Cos(float64(lat)) * math.Sin(float64(lng)),
		math.Sin(float64(lat)),
	}
}

func vectorToLngLat(v Vector) Coord {
	lat := RadiansToDegrees(Radians(math.Asin(v[2])))
	lng := RadiansToDegrees(Radians(math.Atan2(v[1], v[0])))
	return Coord{float64(lng), float64(lat)}
}
