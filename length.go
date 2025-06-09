package goodgeo

// Length measures a length of a geometry in the specified units, (Multi)Point's distance are ignored.
func Length(ls *LineString) Meters {
	length := Meters(0)
	for i := ls.stride; i < len(ls.flatCoords); i += ls.stride {
		a := ls.flatCoords[i-ls.stride : i]
		b := ls.flatCoords[i : i+ls.stride]
		length += Distance(a, b)
	}
	return length
}
