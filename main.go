package main

import (
	"fmt"
	"physicsengine/math/vector2"
	t "physicsengine/object/transform"
	//        "physicsengine/math/rotation"
)

func main() {
	v1 := vector2.NewVector2(3, 4)
	v2 := vector2.NewVector2(-5, -7)
	t1 := t.NewTransform(*v1)
	fmt.Printf("t1 started at %v\n", v1)
	t1 = t1.Translate(v2, 2)
	fmt.Printf("t1 ended at %v\n", t1.Pos)
}
