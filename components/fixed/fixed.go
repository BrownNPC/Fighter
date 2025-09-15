// Package fixed provides simple Q16.16 fixed-point math for fighting games.
package fixed

import "math"

// Fixed represents a 16.16 fixed-point number
type Fixed int32

const (
	Shift = 16
	One   = 1 << Shift
)

// Constructors
func FromInt(n int) Fixed     { return Fixed(n << Shift) }
func FromFloat(f float64) Fixed { return Fixed(f * One) }

// Converters
func (f Fixed) Int() int       { return int(f >> Shift) }
func (f Fixed) Float() float64 { return float64(f) / One }

// Arithmetic
func (a Fixed) Add(b Fixed) Fixed { return a + b }
func (a Fixed) Sub(b Fixed) Fixed { return a - b }

func (a Fixed) Mul(b Fixed) Fixed {
	return Fixed((int64(a) * int64(b)) >> Shift)
}

func (a Fixed) Div(b Fixed) Fixed {
	return Fixed((int64(a) << Shift) / int64(b))
}

// Helpers
func (a Fixed) Abs() Fixed {
	if a < 0 {
		return -a
	}
	return a
}

func (a Fixed) Floor() Fixed { return Fixed(int32(a) &^ (One - 1)) }
func (a Fixed) Ceil() Fixed  { return Fixed((int32(a) + One - 1) &^ (One - 1)) }

// Vector2 is a 2D fixed-point vector
type Vector2 struct {
	X, Y Fixed
}

// Constructors
func Vec2(x, y int) Vector2 {
	return Vector2{FromInt(x), FromInt(y)}
}

func Vec2F(x, y float64) Vector2 {
	return Vector2{FromFloat(x), FromFloat(y)}
}

// Arithmetic
func (v Vector2) Add(u Vector2) Vector2 {
	return Vector2{v.X.Add(u.X), v.Y.Add(u.Y)}
}

func (v Vector2) Sub(u Vector2) Vector2 {
	return Vector2{v.X.Sub(u.X), v.Y.Sub(u.Y)}
}

func (v Vector2) Mul(s Fixed) Vector2 {
	return Vector2{v.X.Mul(s), v.Y.Mul(s)}
}

func (v Vector2) Div(s Fixed) Vector2 {
	return Vector2{v.X.Div(s), v.Y.Div(s)}
}

// Dot product
func (v Vector2) Dot(u Vector2) Fixed {
	return v.X.Mul(u.X).Add(v.Y.Mul(u.Y))
}

// Magnitude (approximate, since sqrt is float-based here)
func (v Vector2) Length() Fixed {
	f := math.Sqrt(v.X.Float()*v.X.Float() + v.Y.Float()*v.Y.Float())
	return FromFloat(f)
}

