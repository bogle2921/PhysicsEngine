package main

import (
	"fmt"
	"physicsengine/math/vector2"
)

func main() {
	v1 := vector2.NewVector2(3, 4)
	v2 := vector2.NewVector2(-5, 1)
	fmt.Printf("v1 + v2 = %v\n", v1.Add(v2))

	fmt.Printf("v1 . v2 = %v\n", v1.Dot(v2))

	v3 := vector2.NewVector2(1, 1)
	fmt.Printf("negate of [1, 1] = %v\n", v3.Negate())
}
