package life

// Color represents an RGBA color.
type Color struct {
	R, G, B, A uint8
}

// Interpolate performs linear interpolation between two colors.
func (c Color) Interpolate(target Color, factor float64) Color {
	return Color{
		R: uint8(float64(c.R) + float64(target.R-c.R)*factor),
		G: uint8(float64(c.G) + float64(target.G-c.G)*factor),
		B: uint8(float64(c.B) + float64(target.B-c.B)*factor),
		A: uint8(float64(c.A) + float64(target.A-c.A)*factor),
	}
}
