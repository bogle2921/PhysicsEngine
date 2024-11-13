package transform

import (
	"fmt"
	v2 "physicsengine/math/vector2"
	"time"
)

type Transform struct {
	Pos v2.Vector2
}

func (t *Transform) Translate(end *v2.Vector2, duration float64) *Transform {
	currTime := 0.0
	startTime := time.Now()

	for currTime < 1.0 {
		elapsed := time.Since(startTime).Seconds()
		currTime = elapsed / duration
		if currTime > 1.0 {
			currTime = 1.0
		}

		currPos := t.Pos.Lerp(end, currTime)
		fmt.Printf("Reached %v\n", currPos)
		time.Sleep(50 * time.Millisecond)
	}
	return &Transform{Pos: *end}
}

func NewTransform(pos v2.Vector2) *Transform {
	return &Transform{Pos: pos}
}

func (t *Transform) String() string {
	return fmt.Sprintf("%v", t.Pos)
}
