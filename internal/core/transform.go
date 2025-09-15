package core

import rl "github.com/gen2brain/raylib-go/raylib"

type Transform struct {
	Position rl.Vector2
	Scale    rl.Vector2
	Rotation float32
}

func NewTransform(x, y float32) *Transform {
	return &Transform{
		Position: rl.Vector2{X: x, Y: y},
		Scale:    rl.Vector2{X: 1, Y: 1},
		Rotation: 0,
	}
}
