package main

import (
	//	"fmt"
	"fmt"
	"physicsengine/math/vector2"
	t "physicsengine/object/transform"
	//        "physicsengine/math/rotation"
)

func main() {
	v1 := vector2.NewVector2(3, 4)
	v2 := vector2.NewVector2(-5, -7)
	/*v2 := vector2.NewVector2(-5, 1)
	fmt.Printf("v1 + v2 = %v\n", v1.Add(v2))

	fmt.Printf("v1 . v2 = %v\n", v1.Dot(v2))

	v3 := vector2.NewVector2(1, 1)
	fmt.Printf("negate of [1, 1] = %v\n", v3.Negate())
	*/
	t1 := t.NewTransform(*v1)
	fmt.Printf("t1 started at %v\n", v1)
	t1 = t1.Translate(v2, 2)
	fmt.Printf("t1 ended at %v\n", t1.Pos)
}
