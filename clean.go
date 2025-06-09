package goodgeo

import (
	"math"
)

// CleanLine removes redundant collinear points from a polyline.
// Polygon handling is intentionally skipped.
func CleanLine(line *LineString) *LineString {
	points := line.Coords()
	n := line.NumCoords()
	if n == 0 || n == 1 {
		return line.Clone()
	}

	if n == 2 && !equals(points[0], points[1]) {
		return line.Clone()
	}

	var out []Coord
	// Indices for the "segment extension" approach: a-b vs a-c.
	a, b, c := 0, 1, 2

	// Always keep the first point.
	out = append(out, points[a])

	for c < n {
		if BooleanPointOnLine(points[b], []Coord{points[a], points[c]}, nil) {
			// b lies on segment a-c, so b is redundant; extend our basis to a-c.
			b = c
		} else {
			// b is not on a-c; commit b.
			out = append(out, points[b])
			a = b
			b++
			c = b
		}
		c++
	}

	// Commit the last point in the current a-b segment.
	out = append(out, points[b])

	nls := NewLineString(line.layout)
	nls.setCoords(points)
	return nls
}

// equals compares two positions with a small tolerance.
// TODO: 3d?
func equals(p, q Coord) bool {
	const eps = 1e-12
	return math.Abs(p[0]-q[0]) <= eps && math.Abs(p[1]-q[1]) <= eps
}

// BooleanPointOnLine returns true if pt lies on any segment of line.
// If opts.IgnoreEndVertices is true, the line's first and last vertices are excluded.
// If opts.Epsilon is nil, exact arithmetic is used for collinearity; otherwise |cross| <= *Epsilon.
func BooleanPointOnLine(pt Coord, line []Coord, opts *struct {
	IgnoreEndVertices bool
	Epsilon           *float64
}) bool {
	n := len(line)
	if n < 2 {
		return false
	}

	for i := 0; i < n-1; i++ {
		exclude := ExcludeNone
		if opts != nil && opts.IgnoreEndVertices {
			switch {
			case i == 0 && i+1 == n-1:
				exclude = ExcludeBoth
			case i == 0:
				exclude = ExcludeStart
			case i == n-2:
				exclude = ExcludeEnd
			}
		}
		if IsPointOnLineSegment(line[i], line[i+1], pt, exclude, getEpsilon(opts)) {
			return true
		}
	}
	return false
}

func getEpsilon(opts *struct {
	IgnoreEndVertices bool
	Epsilon           *float64
}) *float64 {
	if opts == nil {
		return nil
	}
	return opts.Epsilon
}

// BoundaryExclusion controls whether endpoints are considered "on" the segment.
type BoundaryExclusion string

const (
	ExcludeNone  BoundaryExclusion = "none"
	ExcludeStart BoundaryExclusion = "start"
	ExcludeEnd   BoundaryExclusion = "end"
	ExcludeBoth  BoundaryExclusion = "both"
)

// IsPointOnLineSegment checks if pt lies on the finite segment [start,end].
// If epsilon is nil, exact arithmetic is used for collinearity; otherwise |cross| <= *epsilon.
// Boundary handling follows the JS reference: inclusive by default, or exclude start/end/both.
func IsPointOnLineSegment(start, end, pt Coord, exclude BoundaryExclusion, epsilon *float64) bool {
	x, y := pt[0], pt[1]
	x1, y1 := start[0], start[1]
	x2, y2 := end[0], end[1]

	dxc := x - x1
	dyc := y - y1
	dxl := x2 - x1
	dyl := y2 - y1

	// Collinearity via cross product.
	cross := dxc*dyl - dyc*dxl
	if epsilon != nil {
		if math.Abs(cross) > *epsilon {
			return false
		}
	} else if cross != 0 {
		return false
	}

	// Zero-length segment special-case.
	if dxl == 0 && dyl == 0 {
		switch exclude {
		case ExcludeStart, ExcludeEnd, ExcludeBoth:
			return false
		default: // inclusive
			return x == x1 && y == y1
		}
	}

	// Dominant axis bounding checks with endpoint inclusivity per `exclude`.
	absdxl := math.Abs(dxl)
	absdyl := math.Abs(dyl)

	switch exclude {
	case ExcludeStart:
		if absdxl >= absdyl {
			if dxl > 0 {
				return x1 < x && x <= x2
			}
			return x2 <= x && x < x1
		}
		if dyl > 0 {
			return y1 < y && y <= y2
		}
		return y2 <= y && y < y1

	case ExcludeEnd:
		if absdxl >= absdyl {
			if dxl > 0 {
				return x1 <= x && x < x2
			}
			return x2 < x && x <= x1
		}
		if dyl > 0 {
			return y1 <= y && y < y2
		}
		return y2 < y && y <= y1

	case ExcludeBoth:
		if absdxl >= absdyl {
			if dxl > 0 {
				return x1 < x && x < x2
			}
			return x2 < x && x < x1
		}
		if dyl > 0 {
			return y1 < y && y < y2
		}
		return y2 < y && y < y1

	default: // ExcludeNone or unspecified: inclusive
		if absdxl >= absdyl {
			if dxl > 0 {
				return x1 <= x && x <= x2
			}
			return x2 <= x && x <= x1
		}
		if dyl > 0 {
			return y1 <= y && y <= y2
		}
		return y2 <= y && y <= y1
	}
}
