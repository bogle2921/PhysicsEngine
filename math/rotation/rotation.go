package rotation

import (
  "math"
  v2 "physicsengine/math/vector2"
)

/*
| cosTheta   -sinTheta | * | x | = | x * cosTheta - y * sinTheta | 
|  sinTheta  cosTheta  |   | y |   | x * sinTheta + y * cosTheta |
*/
type rotationMatrix struct {
  cosTheta  float64
  sinTheta  float64
}

type Rotation struct {
  vector v2.Vector2
  matrix rotationMatrix
}

func RotateVector(v *v2.Vector2, angle float64) *v2.Vector2 {
  var rm rotationMatrix
  radians := (angle *math.Pi) / 180
  rm.cosTheta = math.Cos(radians)
  rm.sinTheta = math.Sin(radians)
  newX := v.X * rm.cosTheta - v.Y * rm.sinTheta
  newY := v.X * rm.sinTheta + v.Y * rm.cosTheta
  return &v2.Vector2{X: newX, Y: newY}
}
