package goodgeo

// ElevationStats returns elevation gain and loss (respectively).
func ElevationStats(ls *LineString) (Meters, Meters) {
	loss, gain := 0., 0.
	for i := ls.stride + 2; i < len(ls.flatCoords); i += ls.stride {
		ele := ls.flatCoords[i] - ls.flatCoords[i-ls.stride]
		if ele > 0 {
			gain += ele
		} else if ele < 0 {
			loss -= ele
		}
	}

	return Meters(gain), Meters(loss)
}

// ElevationStats returns smoothed elevation gain and loss (respectively).
func ElevationStatsSmoothed(ls *LineString) (Meters, Meters) {
	smoothed := computeSmoothedElevation(ls, 100.)

	loss, gain := 0., 0.
	for i := 1; i < len(smoothed); i++ {
		ele := smoothed[i] - smoothed[i-1]
		if ele > 0 {
			gain += ele
		} else if ele < 0 {
			loss -= ele
		}
	}

	return Meters(gain), Meters(loss)
}

func ElevationLoss(ls *LineString) Meters {
	loss := 0.
	for i := ls.stride + 2; i < len(ls.flatCoords); i += ls.stride {
		ele := ls.flatCoords[i] - ls.flatCoords[i-ls.stride]
		if ele < 0 {
			loss -= ele
		}
	}

	return Meters(loss)
}

func ElevationLossSmoothed(ls *LineString, window ...Meters) Meters {
	w := Meters(100.)
	if len(window) > 0 {
		w = window[0]
	}

	smoothed := computeSmoothedElevation(ls, w)

	loss := 0.
	for i := 1; i < len(smoothed); i++ {
		ele := smoothed[i] - smoothed[i-1]
		if ele < 0 {
			loss -= ele
		}
	}

	return Meters(loss)
}

func ElevationGain(ls *LineString) Meters {
	gain := 0.
	for i := ls.stride + 2; i < len(ls.flatCoords); i += ls.stride {
		ele := ls.flatCoords[i] - ls.flatCoords[i-ls.stride]
		if ele > 0 {
			gain += ele
		}
	}

	return Meters(gain)
}

func ElevationGainSmoothed(ls *LineString, window ...Meters) Meters {
	w := Meters(100.)
	if len(window) > 0 {
		w = window[0]
	}

	smoothed := computeSmoothedElevation(ls, w)

	gain := 0.
	for i := 1; i < len(smoothed); i++ {
		ele := smoothed[i] - smoothed[i-1]
		if ele > 0 {
			gain += ele
		}
	}

	return Meters(gain)
}

func computeSmoothedElevation(ls *LineString, distanceWindow Meters) []float64 {
	accumulate := func(index int) float64 {
		return ls.Coord(index)[2]
	}

	compute := func(accumulated float64, start, end int) float64 {
		return accumulated / float64(end-start+1)
	}

	coords := ls.Coords()
	smoothed := distanceWindowSmoothing(coords, distanceWindow, accumulate, compute)

	if len(coords) > 0 {
		smoothed[0] = coords[0][2]
		smoothed[len(coords)-1] = coords[len(coords)-1][2]
	}

	return smoothed
}

func distanceWindowSmoothing(
	points []Coord,
	distanceWindow Meters,
	accumulate func(index int) float64,
	compute func(accumulated float64, start, end int) float64,
) []float64 {
	result := make([]float64, len(points))

	start := 0
	end := 0
	accumulated := 0.0

	for i := range points {
		for start+1 < i && Distance(points[start], points[i]) > distanceWindow {
			accumulated -= accumulate(start)
			start++
		}

		for end < len(points) && Distance(points[i], points[end]) <= distanceWindow {
			accumulated += accumulate(end)
			end++
		}

		result[i] = compute(accumulated, start, end-1)
	}

	return result
}
