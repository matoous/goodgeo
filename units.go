package goodgeo

import (
	"fmt"
)

// The Earth radius in meters. Used by geom functions that model the Earth as a sphere.
// The [mean radius](https://en.wikipedia.org/wiki/Earth_radius#Arithmetic_mean_radius)
// was selected because it is [recommended](https://rosettacode.org/wiki/Haversine_formula#:~:text=This%20value%20is%20recommended)
// by the Haversine formula to reduce error.
const EarthRadius = 6371008.8

type Meter struct{}

type Meters float64

func (m Meters) String() string {
	return fmt.Sprintf("%.3f m", m)
}

type Radian struct{}

type Radians float64

type Degree struct{}

type Degrees float64

type Mile struct{}

type Miles float64

type Yard struct{}

type Yards float64

type Inch struct{}

type Inches float64

type NauticalMile struct{}

type NauticalMiles float64
