package core

import (
	raylib "github.com/gen2brain/raylib-go/raylib"
)

type Transform struct {
	Position raylib.Vector2
	Scale    raylib.Vector2
	Rotation float32
}

func NewTransform(x, y float32) *Transform {
	return &Transform{
		Position: raylib.Vector2{X: x, Y: y},
		Scale:    raylib.Vector2{X: 1, Y: 1},
		Rotation: 0,
	}
}
