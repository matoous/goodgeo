package goodgeo

// LineSlice returns a subsection of a LineString between startPt and stopPt.
// The coordinates do not need to lie exactly on the line.
func LineSlice(line *LineString, startPt, stopPt Coord) *LineString {
	start := NearestPointOnLine(line, startPt)
	stop := NearestPointOnLine(line, stopPt)

	startIndex := start.Index
	stopIndex := stop.Index

	// Ensure startIndex <= stopIndex
	if startIndex > stopIndex {
		start, stop = stop, start
		startIndex, stopIndex = stopIndex, startIndex
	}

	// Prepare sliced coordinates
	stride := line.stride
	flat := line.flatCoords
	var slicedCoords []float64

	// Start point (interpolated or snapped)
	slicedCoords = append(slicedCoords, start.Point...)

	// Intermediate points
	for i := startIndex; i < stopIndex; i++ {
		slicedCoords = append(slicedCoords, flat[i*stride:i*stride+stride]...)
	}

	// End point (interpolated or snapped)
	slicedCoords = append(slicedCoords, stop.Point...)

	return NewLineStringFlat(line.layout, slicedCoords)
}
