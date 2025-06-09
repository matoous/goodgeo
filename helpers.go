package goodgeo

import "math"

// RadiansToLength convert a distance measurement (assuming a spherical Earth) from radians to meters.
func RadiansToLength(radians Radians) Meters {
	return Meters(float64(radians) * EarthRadius)
}

// DegreesToRadians converts an angle in degrees to radians.
func DegreesToRadians(degrees Degrees) Radians {
	// % 360 degrees in case someone passes value > 360
	normalisedDegrees := math.Mod(float64(degrees), 360.)
	return Radians((normalisedDegrees * math.Pi) / 180.)
}

// Convert a distance measurement (assuming a spherical Earth) from a real-world unit into radians.
func LengthToRadians(distance Meters) Radians {
	return Radians(float64(distance) / EarthRadius)
}

// Converts an angle in radians to degrees.
func RadiansToDegrees(radians Radians) Degrees {
	// % (2 * Math.PI) radians in case someone passes value > 2π
	normalisedRadians := math.Mod(float64(radians), 2*math.Pi)
	return Degrees((normalisedRadians * 180) / math.Pi)
}

// BearingToAzimuth converts any bearing angle from the north line direction (positive clockwise)
// and returns an angle between 0-360 degrees (positive clockwise), 0 being the north line.
func BearingToAzimuth(bearing Degrees) Degrees {
	angle := math.Mod(float64(bearing), 360)
	if angle < 0 {
		angle += 360
	}
	return Degrees(angle)
}

// AzimuthToBearing converts any azimuth angle from the north line direction (positive clockwise)
// and returns an angle between -180 and +180 degrees (positive clockwise), 0 being the north line.
func AzimuthToBearing(angle Degrees) Degrees {
	// Ignore full revolutions (multiples of 360)
	angle = Degrees(math.Mod(float64(angle), 360))

	if angle > 180 {
		return angle - 360
	} else if angle < -180 {
		return angle + 360
	}

	return angle
}
