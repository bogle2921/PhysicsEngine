package vector2

import (
	"fmt"
	"math"
)

type Vector2 struct {
	X float64
	Y float64
}

var ZERO = Vector2{X: 0, Y: 0}
var UP = Vector2{X: 0, Y: 1}
var Down = Vector2{X: 0, Y: -1}
var LEFT = Vector2{X: -1, Y: 0}
var RIGHT = Vector2{X: 1, Y: 0}

func NewVector2(x float64, y float64) *Vector2 {
	return &Vector2{X: x, Y: y}
}

func (v *Vector2) String() string {
	return fmt.Sprintf("[%f, %f]", v.X, v.Y)
}

func (v *Vector2) Add(other *Vector2) *Vector2 {
	x := v.X + other.X
	y := v.Y + other.Y
	return &Vector2{X: x, Y: y}
}

func (v *Vector2) MagnitudeSqrd() float64 {
	return v.X * v.X + v.Y * v.Y
}

func (v *Vector2) Dot(other *Vector2) float64 {
	return v.X*other.X + v.Y*other.Y
}

func (v *Vector2) Scale(val float64) *Vector2 {
	return &Vector2{X: v.X * val, Y: v.Y * val}
}

func (v *Vector2) Negate() *Vector2 {
	return &Vector2{X: -v.X, Y: -v.Y}
}

func (v *Vector2) Equals(location *Vector2) bool {
	return v.X == location.X && v.Y == location.Y
}

func (v *Vector2) Lerp(end *Vector2, speed float64) *Vector2 {
	resVector := v.Add(end)
        // TODO: Optimize out sqrt()
	d := math.Sqrt(resVector.MagnitudeSqrd())	  
	step := d / speed
	for ! v.Equals(end) {
		v.Add(resVector.Scale(step))
		fmt.Printf("vector at %v\n", v)
	}
	return end
}
