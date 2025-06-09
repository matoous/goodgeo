package goodgeo

// Along takes a [LineString] and returns a [Point] at a specified distance along the line.
func Along(line *LineString, distance Meters) *Point {
	coords := line.Coords()
	travelled := Meters(0.)

	for i := range coords {
		if distance >= travelled && i == len(coords)-1 {
			break
		}

		if travelled >= distance {
			overshot := Meters(distance - travelled)
			if overshot == 0 {
				return NewPointFlat(line.layout, coords[i])
			}

			direction := Bearing(coords[i], coords[i-1]) - 180
			interpolated := Destination(
				NewPointFlat(line.layout, coords[i]),
				overshot,
				direction,
			)
			return interpolated
		}

		travelled += Distance(coords[i], coords[i+1])
	}

	return NewPointFlat(line.layout, coords[len(coords)-1])
}
