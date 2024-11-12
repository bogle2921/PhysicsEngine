package object

import (
	"fmt"
	v2 "physicsengine/math/vector2"
)

type Object struct {
	Transform v2.Vector2
	Mass      float64
}

func NewObject(v v2.Vector2, m float64) *Object {
	return &Object{Transform: v, Mass: m}
}

func (o *Object) String() string {
	return fmt.Sprintf("Object is at %v with a mass of %v\n", o.Transform, o.Mass)
}
